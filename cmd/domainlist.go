package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	domainList = &cobra.Command{
		Use:   "domains",
		Short: "Lists ClickOS domains found via XenStore",
		Long:  `Lists all ClickOS xen domains that can be found by checking the values from the XenStore.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("DOMAIN LIST")
		},
	}
)
