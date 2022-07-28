package command

import (
	"fmt"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves the vault token from configured storage",
	Long: `The token helper will return the stored token, if any, on standard out.
- Retrieve the token from the durable storage.
- Print the token on standard out.
- Exit with status code 0 (ok).
`,
	PreRunE: validateHandler,
	Run: func(cmd *cobra.Command, args []string) {
		h, _ := resources.Get(conf.handler)
		token, _ := h.Get(conf.vaultAddr)
		_, _ = fmt.Fprint(os.Stdout, token)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
