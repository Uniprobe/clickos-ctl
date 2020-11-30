package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uniprobe/clickos-ctl/xenstore"
)

var (
	xsSocketPath string
	debug        bool
	jsonOutput   bool

	rootCmd = &cobra.Command{
		Use:   "clickos",
		Short: "Control ClickOS Routers",
		Long: `ClickOS-Ctl allows for the control of ClickOS routers running 
		on the Xen Hypervisor. This version also allows for Unimon elements to
		be used.`,
		Version: "0.1a",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Sudo required for XenStore
			if os.Geteuid() != 0 {
				fmt.Printf("🔐 You need to be root (sudo)\n")
				os.Exit(1)
			}
			// Create XenStore Connection
			// almost all actions will require this!
			if err := xenstore.CreateClient(xsSocketPath); err != nil {
				fmt.Printf("🆘 Xenstore Connection Failed!\n")
				if debug {
					panic(err)
				}
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&xsSocketPath, "xs-socket", "", xenstore.XSSocketPathDefault, "Path to the XenStore UNIX Socket")
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON to stdout")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Show debugging output in stdout")

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
