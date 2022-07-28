package command

import (
	"fmt"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
)

// handlersCmd represents the handler command
var handlersCmd = &cobra.Command{
	Use:   "handlers",
	Short: "Lists available handlers",
	Run: func(cmd *cobra.Command, args []string) {
		handlers := resources.Available()
		if len(handlers) == 0 {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "No available handlers\n")
			return
		}
		for _, h := range handlers {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%-15s - %s\n", h.Name(), h.Description())
		}
	},
}

func init() {
	rootCmd.AddCommand(handlersCmd)
}
