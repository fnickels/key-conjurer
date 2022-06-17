package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	accountsCmd.Flags().StringVar(&identityProvider, "identity-provider", defaultIdentityProvider, "The identity provider to use. Refer to `"+appname+" identity-providers` for more info.")
}

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Prints the list of accounts you have access to.",
	Long:  "Prints the list of accounts you have access to.",
	// Example: appname + " accounts",
	Args: cobra.ExactArgs(0),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list := []string{}
		return list, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := newClient()
		if err != nil {
			return err
		}

		creds, err := config.GetCredentials()
		if err != nil {
			return err
		}

		accounts, err := client.ListAccounts(ctx, &ListAccountsOptions{
			Credentials:            creds,
			AuthenticationProvider: identityProvider,
		})

		if err != nil {
			return err
		}

		var entries []Account
		for _, acc := range accounts {
			entries = append(entries, Account{ID: acc.ID, Name: acc.Name, Alias: generateDefaultAlias(acc.Name)})
		}

		config.UpdateAccounts(entries)
		config.DumpAccounts(os.Stdout)
		return nil
	},
}
