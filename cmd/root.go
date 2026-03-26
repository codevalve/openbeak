package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openbeak",
	Short: "OpenBeak is a stealthy white-hat security tool for hunting malicious agents",
	Long: `A blazing-fast, concurrent predator written in Go that hunts 
exposed or misconfigured AI agent deployments (OpenClaw).`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
