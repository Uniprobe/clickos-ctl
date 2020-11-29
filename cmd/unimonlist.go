package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	unimonList = &cobra.Command{
		Use:       "unimon-elements",
		Short:     "Lists Unimon elements within a given ClickOS router",
		Long:      `Lists all Unimon element found within a ClickOS router.`,
		ValidArgs: []string{"domid", "routerid"},
		Args:      cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("UNIMON LIST")
		},
	}
)
