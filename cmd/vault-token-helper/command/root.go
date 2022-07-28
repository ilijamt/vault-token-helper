package command

import (
	"fmt"
	vth "github.com/ilijamt/vault-token-helper"
	"github.com/ilijamt/vault-token-helper/internal/handler/localpath"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
	"net/url"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   vth.Name,
	Short: "Vault token helper",
	Long:  vth.Description,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

type config struct {
	handler    string
	configPath string
	vaultAddr  *url.URL
}

var conf config

func validateHandler(cmd *cobra.Command, args []string) (err error) {
	if conf.vaultAddr, err = url.Parse(os.Getenv("VAULT_ADDR")); err != nil {
		return err
	}

	conf.configPath = fmt.Sprintf("%s/.vault", os.Getenv("HOME"))
	if val := os.Getenv("VAULT_CONFIG_PATH"); val != "" {
		conf.configPath = val
	}

	_, err = resources.Get(conf.handler)
	return err
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&conf.handler, "handler", "n", localpath.NameLocalPath, "Which handler to use, if you want to see the available handlers run the handlers command")
}
