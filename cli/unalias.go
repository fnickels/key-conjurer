package main

import (
	"github.com/spf13/cobra"
)

var unaliasCmd = cobra.Command{
	Use:     "unalias <accountName/alias>",
	Short:   "Remove alias from account.",
	Example: "  " + appname + " unalias bar",
	Args:    cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list := []string{}
		switch len(args) {
		case 0:
			list = config.HintAccounts()
		}
		return list, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.Unalias(args[0])
	}}
