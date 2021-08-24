package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "localhost", "rbox server")
}

var (
	server  string
	field   string
	value   string
	rootCmd = &cobra.Command{
		Use:     "rboxcli",
		Short:   "Configure remote OpenWrt boxes.",
		Version: "0.0.1",
	}
)

type result struct {
	boxname string
	err     error
	info    string
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
