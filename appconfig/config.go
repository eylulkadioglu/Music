package appconfig

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Salt           string `mapstructure:"SALT"`
	ListenAddress  string `mapstructure:"LISTEN_ADDR"`
	DbDSN          string `mapstructure:"DB_DSN"`
	MailerAddress  string `mapstructure:"MAILER_ADDR"`
	MailerPort     int    `mapstructure:"MAILER_PORT"`
	MailerFrom     string `mapstructure:"MAILER_FROM"`
	MailerPassword string `mapstructure:"MAILER_PASSWD"`
}

func ReadConfig() *AppConfig {
	var appConfig *AppConfig

	viper.SetConfigName("musicdbgo") // name of config file (without extension)
	viper.SetConfigType("yaml")      // Required if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/musicdbgo")
	viper.AutomaticEnv() 

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("Cannot find the configuration file, can't work without one. Quitting\n")
		os.Exit(-1)
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		fmt.Printf("Failed to process config file\n")
		os.Exit(-2)
	}

	return appConfig
}
