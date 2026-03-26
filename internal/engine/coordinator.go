package engine

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/codevalve/openbeak/internal/models"
)

// Coordinator manages the lifecycle of a scan operation.
type Coordinator struct {
	Workers   int
	Inking    models.InkingLevel
	Tentacles []models.Tentacle
	Inks      []models.Ink
	Results   chan models.Result
	OnProgress func(float64) // Callback for progress updates
	WaitGroup  sync.WaitGroup
}

// NewCoordinator initializes a new scan engine.
func NewCoordinator(workers int) *Coordinator {
	return &Coordinator{
		Workers:   workers,
		Results:   make(chan models.Result, 100),
		Tentacles: []models.Tentacle{},
		Inks:      []models.Ink{},
	}
}

// RegisterTentacle adds a scanning module to the engine.
func (c *Coordinator) RegisterTentacle(t models.Tentacle) {
	c.Tentacles = append(c.Tentacles, t)
}

// RegisterInk adds a reporting module to the engine.
func (c *Coordinator) RegisterInk(i models.Ink) {
	c.Inks = append(c.Inks, i)
}

// Scan targets using registered tentacles.
func (c *Coordinator) Scan(ctx context.Context, targets []string) {
	c.log(ctx, models.Info, "Starting scan operation", "")
	total := len(targets)
	var processed int32

	targetChan := make(chan string, len(targets))

	// Spawn workers
	for i := 0; i < c.Workers; i++ {
		c.WaitGroup.Add(1)
		go func() {
			defer c.WaitGroup.Done()
			for target := range targetChan {
				// Update progress
				curr := atomic.AddInt32(&processed, 1)
				if c.OnProgress != nil {
					c.OnProgress(float64(curr) / float64(total))
				}

				c.log(ctx, models.Debug, "Probing target", target)
				for _, t := range c.Tentacles {
					select {
					case <-ctx.Done():
						return
					default:
						res, err := t.Probe(ctx, target)
						if err == nil {
							// Filter findings based on InkingLevel
							shouldFile := false
							switch c.Inking {
							case models.InkingStealth:
								shouldFile = (res.Severity == "High")
							case models.InkingTactical:
								shouldFile = (res.Severity == "High" || res.Severity == "Medium")
							case models.InkingVerbose:
								shouldFile = true
							}

							if shouldFile {
								c.Results <- res
								for _, ink := range c.Inks {
									_ = ink.Write(ctx, res)
								}
							}
						} else if c.Inking == models.InkingVerbose {
							c.log(ctx, models.Debug, fmt.Sprintf("%s: %v", t.Name(), err), target)
						}
					}
				}
			}
		}()
	}

	// Feed targets
	for _, target := range targets {
		targetChan <- target
	}
	close(targetChan)

	// Wait for workers to finish
	c.WaitGroup.Wait()
	c.log(ctx, models.Info, "Scan operation complete", "")
	close(c.Results)
}

func (c *Coordinator) log(ctx context.Context, level models.ActivityLevel, msg string, target string) {
	event := models.ActivityEvent{
		Timestamp: time.Now(),
		Level:     level,
		Component: "engine",
		Message:   msg,
		Target:    target,
	}
	for _, i := range c.Inks {
		_ = i.Log(ctx, event)
	}
}
