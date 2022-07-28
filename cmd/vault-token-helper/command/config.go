package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

// configCmd represents the erase command
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Show current configuration.",
	PreRunE: validateHandler,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Handler: %s\n", conf.handler)
		fmt.Printf("Vault config path: %s\n", conf.configPath)
		fmt.Printf("Vault address: %s (%s)\n", conf.vaultAddr.String(), conf.vaultAddr.Hostname())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
