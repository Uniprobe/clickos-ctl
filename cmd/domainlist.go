package cmd

import (
	"fmt"
	"os"

	"github.com/uniprobe/clickos-ctl/xen"

	"github.com/spf13/cobra"
)

var (
	domainList = &cobra.Command{
		Use:   "domains",
		Short: "Lists ClickOS domains found via XenStore",
		Long:  `Lists all ClickOS xen domains that can be found by checking the values from the XenStore.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			domains, err := xen.GetClickOSDomains()
			if err != nil {
				fmt.Printf("🆘 Failed to get domains list!\n")
				if debug {
					fmt.Printf("🆘 Error: %v\n\n", err)
				}
				os.Exit(1)
			}
			if !jsonOutput {
				for _, dom := range domains {
					fmt.Printf("  ID: %d\n", dom.ID)
					fmt.Printf("  Name: %s\n\n", dom.Name)
				}
			} else {

			}
			os.Exit(0)
		},
	}
)
