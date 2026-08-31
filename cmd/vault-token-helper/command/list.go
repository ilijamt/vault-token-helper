package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"text/tabwriter"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/ilijamt/vault-token-helper/internal/handler"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/spf13/cobra"
)

var listExtended bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List currently available vault addresses defined.",
	Long:    `Prints a list of all available addresses. Doesn't mean that the tokens stored are valid, they may be expired.`,
	PreRunE: validateHandler,
	RunE: func(cmd *cobra.Command, args []string) error {
		var h, _ = resources.Get(conf.handler)
		items, err := h.List()
		if err != nil {
			return err
		}

		if !listExtended {
			for _, item := range items {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", item.Address)
			}
			return nil
		}

		addrs := make([]string, 0, len(items))
		for _, item := range items {
			addrs = append(addrs, item.Address)
		}
		extendedList(cmd, h, addrs)
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listExtended, "extended", "e", false, "Lookup and print extended details about each token by querying each Vault server")
	rootCmd.AddCommand(listCmd)
}

type extendedRow struct {
	line string
	err  error
}

// extendedList queries each address's own Vault server (in parallel) with
// its stored token and renders a table of display_name/ttl/renewable/policies.
// A failure on one address (expired token, unreachable server, ...) is
// rendered as an error row rather than aborting the whole listing.
func extendedList(cmd *cobra.Command, h handler.Handler, addrs []string) {
	wg := &sync.WaitGroup{}
	rows := make([]extendedRow, len(addrs))

	for i, addr := range addrs {
		wg.Add(1)
		go func(i int, addr string) {
			defer wg.Done()
			rows[i] = lookupExtendedRow(h, addr)
		}(i, addr)
	}
	wg.Wait()

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 5, 0, 2, ' ', tabwriter.DiscardEmptyColumns)
	_, _ = fmt.Fprintln(w, "VAULT_ADDR\tdisplay_name\tttl\trenewable\tpolicies\t")
	_, _ = fmt.Fprintln(w, "----------\t------------\t---\t---------\t--------\t")
	for i, addr := range addrs {
		if rows[i].err != nil {
			_, _ = fmt.Fprintf(w, "%s\t** ERROR **\t%s\t\t\t\n", addr, rows[i].err)
			continue
		}
		_, _ = fmt.Fprintln(w, rows[i].line)
	}
	_ = w.Flush()
}

func lookupExtendedRow(h handler.Handler, addr string) (row extendedRow) {
	vaultAddr, err := url.Parse(addr)
	if err != nil {
		row.err = err
		return
	}

	token, err := h.Get(vaultAddr)
	if err != nil {
		row.err = err
		return
	}

	vcfg := vault.DefaultConfig() // picks up VAULT_ env vars, overridden below
	vcfg.Timeout = 5 * time.Second
	vcfg.MaxRetries = 1
	vcfg.Address = addr

	client, err := vault.NewClient(vcfg)
	if err != nil {
		row.err = err
		return
	}
	client.SetToken(token)

	s, err := client.Auth().Token().LookupSelf()
	if err != nil {
		row.err = err
		return
	}

	ttl, err := s.Data["ttl"].(json.Number).Int64()
	if err != nil {
		row.err = err
		return
	}

	row.line = fmt.Sprintf("%s\t%s\t%s\t%t\t%s",
		addr,
		s.Data["display_name"],
		time.Duration(ttl)*time.Second,
		s.Data["renewable"],
		s.Data["policies"],
	)
	return
}
