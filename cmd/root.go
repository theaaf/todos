package cmd

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todo Web Application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !terminal.IsTerminal(unix.Stdout) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		} else {
			logrus.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339Nano,
			})
		}

		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "make output more verbose")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/todos")
		viper.AddConfigPath("$HOME/.todos")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Warnf("unable to read config from file")
	}
}
