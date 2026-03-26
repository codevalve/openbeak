package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "openbeak",
	Short: "OpenBeak: A stealthy white-hat predator for hunting malicious OpenClaw deployments.",
	Long: `OpenBeak (Macroctopus Agentaculum) is a blazing-fast, concurrent scanner 
designed to probe and neutralize malicious or misconfigured OpenClaw instances. 

Built with Go and the Charmbracelet TUI stack for a premium, stealthy terminal experience.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Root flags and global configurations can be defined here if needed.
}
