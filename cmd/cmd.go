package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	xsSocketPathDefault string = "/var/run/xenstored/socket"
)

var (
	// Used for flags.
	xsSocketPath string
	userLicense  string

	rootCmd = &cobra.Command{
		Use:   "clickos",
		Short: "Control ClickOS Routers",
		Long: `ClickOS-Ctl allows for the control of ClickOS routers running 
		on the Xen Hypervisor. This version also allows for Unimon elements to
		be used.`,
		Version: "0.1a",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&xsSocketPath, "xs-socket", "", xsSocketPathDefault, "Path to the XenStore UNIX Socket")

	rootCmd.AddCommand(domainList)
	rootCmd.AddCommand(routerList)
	rootCmd.AddCommand(unimonList)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
