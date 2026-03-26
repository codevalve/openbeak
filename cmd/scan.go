package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codevalve/openbeak/internal/engine"
	"github.com/codevalve/openbeak/internal/models"
	"github.com/codevalve/openbeak/internal/tentacles"
	"github.com/codevalve/openbeak/internal/tui"
	"github.com/spf13/cobra"
)

var (
	targetFile  string
	cidrFlag    string
	workers     int
	inkingLevel string
	outJson     string
	outLog      string
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
		coord.Inking = models.ParseInkingLevel(inkingLevel)
		coord.OnProgress = func(pct float64) {
			p.Send(tui.ProgressMsg(pct))
		}
		for _, h := range tentacles.Hunters {
			coord.RegisterTentacle(h)
		}

		for _, i := range tentacles.Inks {
			if jsonInk, ok := i.(*tentacles.JSONInk); ok && outJson != "" {
				jsonInk.FilePath = outJson
			}
			if activityInk, ok := i.(*tentacles.ActivityInk); ok && outLog != "" {
				activityInk.FilePath = outLog
			}
			coord.RegisterInk(i)
		}

		// 3. Load Targets
		var targets []string
		if targetFile != "" {
			data, err := os.ReadFile(targetFile)
			if err != nil {
				fmt.Printf("Error reading targets file: %v\n", err)
				os.Exit(1)
			}
			for _, line := range strings.Split(string(data), "\n") {
				trimmed := strings.TrimSpace(line)
				if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
					targets = append(targets, trimmed)
				}
			}
		}

		if cidrFlag != "" {
			expanded, err := models.ExpandCIDR(cidrFlag)
			if err != nil {
				fmt.Printf("Error expanding CIDR: %v\n", err)
				os.Exit(1)
			}
			targets = append(targets, expanded...)
		}

		// Fallback to local mocks if no targets provided at all
		if len(targets) == 0 {
			targets = []string{"localhost:8080", "127.0.0.1:3000"}
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
	scanCmd.Flags().StringVarP(&cidrFlag, "cidr", "c", "", "CIDR range to scan (e.g., 192.168.1.0/24)")
	scanCmd.Flags().IntVarP(&workers, "workers", "w", 10, "Number of concurrent workers")
	scanCmd.Flags().StringVarP(&inkingLevel, "inking", "i", "stealth", "Inking detail level (stealth, tactical, verbose)")
	scanCmd.Flags().StringVar(&outJson, "out-json", "openbeak_results.json", "File path to save the JSON results payload")
	scanCmd.Flags().StringVar(&outLog, "out-log", "openbeak_activity.log", "File path to save the operational activity log")
	rootCmd.AddCommand(scanCmd)
}
