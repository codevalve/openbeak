package cmd

import (
	"fmt"

	"github.com/codevalve/openbeak/internal/tentacles"
	"github.com/spf13/cobra"
)

var tentaclesCmd = &cobra.Command{
	Use:   "tentacles",
	Short: "Manage and list interactive modules (Hunters and Inks)",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available hunters and inks registered in the engine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available Hunters:")
		for _, h := range tentacles.Hunters {
			fmt.Printf(" - %-20s %s\n", h.Name(), h.Description())
		}
		fmt.Println("\nAvailable Inks:")
		for _, i := range tentacles.Inks {
			fmt.Printf(" - %-20s %s\n", i.Name(), i.Description())
		}
	},
}

func init() {
	tentaclesCmd.AddCommand(listCmd)
	rootCmd.AddCommand(tentaclesCmd)
}
