package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var addUserMigration_0001 = &Migration{
	Number: 1,
	Name:   "Add user",
	Forwards: func(db *gorm.DB) error {
		const addUserSQL = `
			CREATE TABLE users(
 				id serial PRIMARY KEY,
 				email text UNIQUE NOT NULL,
				hashed_password bytea NOT NULL,
 				created_at TIMESTAMP NOT NULL,
 				updated_at TIMESTAMP NOT NULL,
 				deleted_at TIMESTAMP
			);
		`

		err := db.Exec(addUserSQL).Error
		return errors.Wrap(err, "unable to create users table")
	},
}

func init() {
	Migrations = append(Migrations, addUserMigration_0001)
}
