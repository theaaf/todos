package cmd

import (
	"sort"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/theaaf/todos/app"
	"github.com/theaaf/todos/migrations"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		number, _ := cmd.Flags().GetInt("number")
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		if dryRun {
			logrus.Info("=== DRY RUN ===")
		}

		sort.Slice(migrations.Migrations, func(i, j int) bool {
			return migrations.Migrations[i].Number < migrations.Migrations[j].Number
		})

		a, err := app.New()
		if err != nil {
			return err
		}
		defer a.Close()

		// Make sure Migration table is there
		logrus.Debug("ensuring migrations table is present")
		if err := a.Database.AutoMigrate(&migrations.Migration{}).Error; err != nil {
			return errors.Wrap(err, "unable to automatically migrate migrations table")
		}

		var latest migrations.Migration
		if err := a.Database.Order("number desc").First(&latest).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
			return errors.Wrap(err, "unable to find latest migration")
		}

		noMigrationsApplied := latest.Number == 0

		if noMigrationsApplied && len(migrations.Migrations) == 0 {
			logrus.Info("no migrations to apply")
			return nil
		}

		if latest.Number >= migrations.Migrations[len(migrations.Migrations)-1].Number {
			logrus.Info("no migrations to apply")
			return nil
		}

		if number == -1 {
			number = int(migrations.Migrations[len(migrations.Migrations)-1].Number)
		}

		if uint(number) <= latest.Number && latest.Number > 0 {
			logrus.Info("no migrations to apply, specified number is less than or equal to latest migration; backwards migrations are not supported")
			return nil
		}

		for _, migration := range migrations.Migrations {
			if migration.Number > uint(number) {
				break
			}

			logger := logrus.WithField("migration_number", migration.Number)
			logger.Infof("applying migration %q", migration.Name)

			if dryRun {
				continue
			}

			tx := a.Database.Begin()

			if err := migration.Forwards(tx); err != nil {
				logger.WithError(err).Error("unable to apply migration, rolling back")
				if err := tx.Rollback().Error; err != nil {
					logger.WithError(err).Error("unable to rollback...")
				}
				break
			}

			if err := tx.Commit().Error; err != nil {
				logger.WithError(err).Error("unable to commit transaction...")
				break
			}

			if err := a.Database.Create(migration).Error; err != nil {
				logger.WithError(err).Error("unable to create migration record")
				break
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().Int("number", -1, "the migration to run forwards until; if not set, will run all migrations")
	migrateCmd.Flags().Bool("dry-run", false, "print out migrations to be applied without running them")
}
