package command

import (
	"fmt"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:     "purge",
	Short:   "Purge all available vault tokens.",
	PreRunE: validateHandler,
	RunE: func(cmd *cobra.Command, args []string) error {
		var h, _ = resources.Get(conf.handler)
		items, err := h.Purge(conf.dryRun)
		if err == nil {
			for _, item := range items {
				if conf.dryRun {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Would delete the token for %s\n", item.Address)
				} else {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Deleted token for %s\n", item.Address)
				}
			}
		}
		return err
	},
}

func init() {
	purgeCmd.Flags().BoolVar(&conf.dryRun, "dry-run", false, "Should we on show which ones we delete?")
	rootCmd.AddCommand(purgeCmd)
}
