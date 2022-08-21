package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// setupCmd represents the erase command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the token helper with vault.",
	Long: `The token helper will set itself up with vault. 
- Will create a $HOME/.vault file or VAULT_CONFIG_PATH if the env variable is set, if the file does not exist
- Add/Update a token_helper directive with the current helper
`,
	PreRunE: validateHandler,
	RunE: func(cmd *cobra.Command, args []string) error {
		var ex, _ = os.Executable()
		var payload = fmt.Sprintf("token_helper = \"%s\"\n", ex)
		return os.WriteFile(conf.configPath, []byte(payload), 0650)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
