package main

import (
	"github.com/spf13/cobra"
)

var aliasCmd = cobra.Command{
	Use:     "alias <accountName> <alias>",
	Short:   "Give an account a nickname.",
	Long:    "Alias an account to a nickname so you can refer to the account by the nickname.",
	Example: "  " + appname + " alias FooAccount Bar",
	Args:    cobra.ExactArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list := []string{}
		switch len(args) {
		case 0:
			list = config.HintAccounts()
		}

		// add valid flags and subcommands
		list = append(list, flagHints(cmd)...)

		return list, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.Alias(args[0], args[1])
	}}
