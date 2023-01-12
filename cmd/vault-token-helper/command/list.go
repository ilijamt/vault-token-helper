package command

import (
	"fmt"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List currently available vault addresses defined.",
	Long:    `Prints a list of all available addresses. Doesn't mean that the tokens stored are valid, they may be expired.`,
	PreRunE: validateHandler,
	RunE: func(cmd *cobra.Command, args []string) error {
		var h, _ = resources.Get(conf.handler)
		items, err := h.List()
		if err == nil {
			for _, item := range items {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", item.Address)
			}
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
