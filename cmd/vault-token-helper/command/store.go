package command

import (
	"bufio"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "The token to store will be provided on standard in.",
	Long: `The token to store will be provided on standard in. 
- Strip any leading or trailing whitespace from the provided token on standard in. Vault tokens will never have leading or trailing whitespace, but some shells can inadvertently add it.
- Verify that the provided token is not the empty string. If the value is the empty string, call erase instead.
- Persist the token in a durable manner. Because the binary is executed as a command (not a server), most token helpers cannot persist the token in memory; they will need to write the token to disk, a keychain, or some external tool.
- Exit with status code 0 (ok)`,
	PreRunE: validateHandler,
	Run: func(cmd *cobra.Command, args []string) {
		h, _ := resources.Get(conf.handler)
		var token string
		var err error

		r := bufio.NewReader(os.Stdin)

		if token, err = r.ReadString('\n'); err != nil && err != io.EOF {
			os.Exit(1)
		}

		token = strings.TrimSpace(token)
		if token == "" {
			_ = h.Erase(conf.vaultAddr)
			return
		}

		_ = h.Store(token, conf.vaultAddr)
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
