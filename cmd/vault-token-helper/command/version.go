package command

import (
	"fmt"
	vth "github.com/ilijamt/vault-token-helper"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of the application",
	Long:  `Shows the version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Version: %s\n", vth.BuildVersion)
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Git Commit Hash: %s\n", vth.BuildHash)
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Build Date: %s\n", vth.BuildDate)
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "OS: %s\n", runtime.GOOS)
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Architecture: %s\n", runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
