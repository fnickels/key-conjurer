package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func flagHints(cmd *cobra.Command) []string {
	list := []string{}

	flagset := cmd.Flags()
	if flagset.HasAvailableFlags() {
		flagset.VisitAll(func(f *pflag.Flag) {
			if !f.Hidden && !f.Changed {
				if f.Name != "" {
					list = append(list, "--"+f.Name)
				}
				if f.Shorthand != "" {
					list = append(list, "-"+f.Shorthand)
				}
			}
		})
	}
	return list
}

func nullLookup(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func dirLookup(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveFilterDirs
}

func fileLookup(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveDefault
}

func outtypeLookup(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	list := []string{
		"awscli",
		"env",
	}
	return list, cobra.ShellCompDirectiveNoFileComp
}

func identityProviderLookup(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	list := []string{}

	providers, err := listproviders()
	if err == nil {
		for _, provider := range providers {
			list = append(list, provider.ID)
		}
	}

	return list, cobra.ShellCompDirectiveNoFileComp
}
