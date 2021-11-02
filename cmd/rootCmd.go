package cmd

import "github.com/spf13/cobra"

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use: "cobra",
}

/* commands
update ER dir names with CUID
check that there is a CUID dir for every entry in a Work Order
*/
