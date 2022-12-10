// Package cmd contains commands
/*
Copyright © 2022 clarkezone

*/
package cmd

import (
	"os"
	"strings"

	"github.com/clarkezone/pocketshorten/internal"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	outputMode string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pocketshorten",
	Short: "pocketshorten is a URL shortener.",
	Long: `pocketshorten is a URL shortener based on the pocketbase backend.

To install the completion script on zsh,
ensure that .zshrc contains:
  autoload -U compinit
  compinit -i

Then run the following:
  sudo -- sh -c 'pocketshorten completion zsh > /usr/local/share/zsh/site-functions/_pocketshorten'

to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		clarkezoneLog.SetLevel(internal.LogLevel)
		clarkezoneLog.SetOutputFormat(outputMode)
		clarkezoneLog.Infof("started %s", strings.Join(os.Args, " "))
		return nil
	},
	PersistentPostRun: func(ccmd *cobra.Command, args []string) {
		clarkezoneLog.Infof("finished %s", strings.Join(os.Args, " "))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pocketshorten.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntVar(&internal.Port, internal.PortVar, viper.GetInt(internal.PortVar), "server port")
	rootCmd.PersistentFlags().IntVar(&internal.MetricsPort, internal.MetricsPortVar, viper.GetInt(internal.MetricsPortVar), "metrics port")
	rootCmd.PersistentFlags().StringVarP(&internal.LogLevel, internal.LogLevelVar, "l",
		viper.GetString(internal.LogLevelVar), "amount of information outputted (debug, info, warn, error)")
	rootCmd.PersistentFlags().StringVar(&outputMode, "logoutput", clarkezoneLog.TTYFormat,
		"output format for logs (tty, plain, json)")
	err := viper.BindPFlag(internal.PortVar, rootCmd.PersistentFlags().Lookup(internal.PortVar))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag(internal.MetricsPortVar, rootCmd.PersistentFlags().Lookup(internal.MetricsPortVar))
	if err != nil {
		panic(err)
	}

	err = viper.BindPFlag(internal.LogLevelVar, rootCmd.PersistentFlags().Lookup(internal.LogLevelVar))
	if err != nil {
		panic(err)
	}

	_, err = newTestServerGrpcCmd(rootCmd)
	if err != nil {
		panic(err)
	}

	_, err = newTestClientGrpcCmd(rootCmd)
	if err != nil {
		panic(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".previewd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".previewd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		clarkezoneLog.Infof("Using config file:", viper.ConfigFileUsed())
	}
}
