package command

import (
	"github.com/ilijamt/vault-token-helper/internal/resources"

	"github.com/spf13/cobra"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "The token helper should scrub the stored token, if any.",
	Long: `The token helper should scrub the stored token, If no value is stored, the token helper should simply return. A token helper should never return an error if a value is not stored.
- Purge or scrub any trace of the token.
- Exit with status code 0 (ok).
`,
	PreRunE: validateHandler,
	Run: func(cmd *cobra.Command, args []string) {
		h, _ := resources.Get(conf.handler)
		_ = h.Erase(conf.vaultAddr)
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}
