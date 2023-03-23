// Package internal has utilities
package internal

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig(cfg *string) {
	clarkezoneLog.Debugf("initConfig called with %v", *cfg)
	if *cfg != "" {
		// Use config file from the flag.
		viper.SetConfigFile(*cfg)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".pocketshorten" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".pocketshorten")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// To debug this with logging, change the initial log value in main.go
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		clarkezoneLog.Debugf("Using config file: %v", viper.ConfigFileUsed())
	} else {
		clarkezoneLog.Errorf("Error reading config %v", err)
	}
}
