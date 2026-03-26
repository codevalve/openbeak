package cmd

import (
	"fmt"
	"os"

	"github.com/codevalve/openbeak/internal/tentacles"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe [name]",
	Short: "Provide documentation for a specific tentacle or reporter",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Available Hunters:")
			for _, h := range tentacles.Hunters {
				fmt.Printf(" - %-20s %s\n", h.Name(), h.Description())
			}
			fmt.Println("\nAvailable Reporters:")
			for _, r := range tentacles.Reporters {
				fmt.Printf(" - %-20s %s\n", r.Name(), r.Description())
			}
			return
		}

		name := args[0]
		if h := tentacles.GetTentacle(name); h != nil {
			fmt.Printf("Name:        %s\n", h.Name())
			fmt.Printf("Role:        %s\n", h.Role())
			fmt.Printf("Description: %s\n", h.Description())
			return
		}

		if r := tentacles.GetReporter(name); r != nil {
			fmt.Printf("Name:        %s\n", r.Name())
			fmt.Printf("Role:        Reporter\n")
			fmt.Printf("Description: %s\n", r.Description())
			return
		}

		fmt.Printf("Error: Module '%s' not found.\n", name)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
