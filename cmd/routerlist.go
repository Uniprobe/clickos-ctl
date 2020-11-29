package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	routerList = &cobra.Command{
		Use:       "routers",
		Short:     "Lists ClickOS routers within a given domain",
		Long:      `Lists all ClickOS routers (halted or not) found within a ClickOS domain.`,
		ValidArgs: []string{"domid"},
		Args:      cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ROUTER LIST")
		},
	}
)
