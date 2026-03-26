package cmd

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codevalve/openbeak/internal/engine"
	"github.com/codevalve/openbeak/internal/tentacles"
	"github.com/codevalve/openbeak/internal/tui"
	"github.com/spf13/cobra"
)

var (
	targetFile string
	workers    int
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the horizon for malicious agent deployments",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Initialize TUI
		m := tui.NewModel()
		p := tea.NewProgram(m)

		// 2. Initialize Engine
		coord := engine.NewCoordinator(workers)
		coord.RegisterTentacle(&tentacles.HTTPDiscoveryTentacle{
			Timeout: 2, // 2 seconds
		})

		// Register JSON Reporter
		coord.RegisterReporter(&tentacles.JSONReporter{
			FilePath: "openbeak_results.json",
		})

		// 3. Load Targets (Mock for now or read from file)
		targets := []string{"localhost:8080", "127.0.0.1:3000"}
		if targetFile != "" {
			// In a real implementation, we'd read the file here.
			fmt.Printf("Reading targets from %s...\n", targetFile)
		}

		// 4. Run Scan in Background
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			coord.Scan(ctx, targets)
			p.Send(tui.ScanCompleteMsg{})
		}()

		// 5. Pipe results to TUI
		go func() {
			for res := range coord.Results {
				p.Send(tui.ResultMsg(res))
			}
		}()

		// 6. Start TUI
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&targetFile, "targets", "t", "", "File containing target hosts")
	scanCmd.Flags().IntVarP(&workers, "workers", "w", 10, "Number of concurrent workers")
	rootCmd.AddCommand(scanCmd)
}
