package main

import (
	"fmt"

	"github.com/riotgames/key-conjurer/api/keyconjurer"
	"github.com/spf13/cobra"
)

func init() {
	rolesCmd.Flags().StringVar(&identityProvider, "identity-provider", defaultIdentityProvider, "The identity provider to retrieve roles from")

	rolesCmd.RegisterFlagCompletionFunc("identity-provider", identityProviderLookup)
}

var rolesCmd = cobra.Command{
	Use:   "roles",
	Short: "List all the roles that you can assume when using `" + appname + " get`.",
	Args:  cobra.ExactArgs(0),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list := []string{}
		// add valid flags and subcommands
		list = append(list, flagHints(cmd)...)
		return list, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(*cobra.Command, []string) error {
		switch identityProvider {
		case keyconjurer.AuthenticationProviderOneLogin:
			return fmt.Errorf("roles are not used in the OneLogin authentication provider")
		case keyconjurer.AuthenticationProviderOkta:
			return fmt.Errorf(`You cannot retrieve roles for %q from the command line at this time. Instead, please check the instructions you have received from the team that manages KeyConjurer within your organization`, identityProvider)
		default:
			return fmt.Errorf("unsupported identity provider %q", identityProvider)
		}
	},
}
