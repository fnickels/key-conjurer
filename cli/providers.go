package main

import (
	"context"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/riotgames/key-conjurer/api/keyconjurer"
	"github.com/spf13/cobra"
)

var providersCmd = cobra.Command{
	Use:   "identity-providers",
	Short: "List identity providers you may use.",
	Long: fmt.Sprintf(`List all identity providers that KeyConjurer supports through which the user may authenticate.

If KeyConjurer supports multiple providers, you may specify one you wish to use with the --identity-provider flag.

If you do not specify an --identity-provider flag for the commands that support it (get, login, accounts) a default identity provider will be chosen for you (default: %q).
`, defaultIdentityProvider),
	// Example: appname + " identity-providers",
	Args: cobra.ExactArgs(0),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list := []string{}
		// add valid flags and subcommands
		list = append(list, flagHints(cmd)...)
		return list, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		providers, err := listproviders()
		if err != nil {
			return err
		}

		tw := tablewriter.NewWriter(os.Stdout)
		tw.SetHeader([]string{"ID"})
		for _, provider := range providers {
			tw.Append([]string{provider.ID})
		}

		tw.Render()
		return nil
	},
}

func listproviders() ([]keyconjurer.Provider, error) {

	var empty []keyconjurer.Provider

	ctx := context.Background()
	client, err := newClient()
	if err != nil {
		return empty, err
	}

	providers, err := client.ListProviders(ctx, &ListProvidersOptions{})
	if err != nil {
		return empty, err
	}

	return providers, err
}
