package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/codevalve/openbeak/internal/models"
)

// Coordinator manages the lifecycle of a scan operation.
type Coordinator struct {
	Workers   int
	Tentacles []models.Tentacle
	Reporters []models.Reporter
	Results   chan models.Result
	WaitGroup sync.WaitGroup
}

// NewCoordinator initializes a new scan engine.
func NewCoordinator(workers int) *Coordinator {
	return &Coordinator{
		Workers:   workers,
		Results:   make(chan models.Result, 100),
		Tentacles: []models.Tentacle{},
		Reporters: []models.Reporter{},
	}
}

// RegisterTentacle adds a scanning module to the engine.
func (c *Coordinator) RegisterTentacle(t models.Tentacle) {
	c.Tentacles = append(c.Tentacles, t)
}

// RegisterReporter adds a reporting module to the engine.
func (c *Coordinator) RegisterReporter(r models.Reporter) {
	c.Reporters = append(c.Reporters, r)
}

// Scan targets using registered tentacles.
func (c *Coordinator) Scan(ctx context.Context, targets []string) {
	c.log(ctx, models.Info, "Starting scan operation", "")

	targetChan := make(chan string, len(targets))

	// Spawn workers
	for i := 0; i < c.Workers; i++ {
		c.WaitGroup.Add(1)
		workerID := i
		go func() {
			defer c.WaitGroup.Done()
			for target := range targetChan {
				c.log(ctx, models.Debug, fmt.Sprintf("Worker %d probing target", workerID), target)
				for _, t := range c.Tentacles {
					select {
					case <-ctx.Done():
						return
					default:
						res, err := t.Probe(ctx, target)
						if err == nil {
							c.Results <- res
							// Dispatch to reporters
							for _, r := range c.Reporters {
								_ = r.Write(ctx, res)
							}
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
	for _, r := range c.Reporters {
		_ = r.Log(ctx, event)
	}
}
