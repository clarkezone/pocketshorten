package cmd

import (
	"fmt"

	"github.com/clarkezone/pocketshorten/pkg/config"
	"github.com/spf13/cobra"
)

// Show current version
var versionCommand = getVersionCommand()

func init() {
	rootCmd.AddCommand(versionCommand)
}

func getVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show pocketshorten version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "pocketshorten version:%s hash:%s\n", config.VersionString, config.VersionHash)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
