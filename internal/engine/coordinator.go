package engine

import (
	"context"
	"sync"

	"github.com/codevalve/openbeak/internal/models"
)

// Coordinator manages the lifecycle of a scan operation.
type Coordinator struct {
	Workers   int
	Tentacles []models.Tentacle
	Results   chan models.Result
	WaitGroup sync.WaitGroup
}

// NewCoordinator initializes a new scan engine.
func NewCoordinator(workers int) *Coordinator {
	return &Coordinator{
		Workers: workers,
		Results: make(chan models.Result, 100),
	}
}

// RegisterTentacle adds a scanning module to the engine.
func (c *Coordinator) RegisterTentacle(t models.Tentacle) {
	c.Tentacles = append(c.Tentacles, t)
}

// Scan targets using registered tentacles.
func (c *Coordinator) Scan(ctx context.Context, targets []string) {
	targetChan := make(chan string, len(targets))

	// Spawn workers
	for i := 0; i < c.Workers; i++ {
		c.WaitGroup.Add(1)
		go func() {
			defer c.WaitGroup.Done()
			for target := range targetChan {
				for _, t := range c.Tentacles {
					select {
					case <-ctx.Done():
						return
					default:
						res, err := t.Probe(ctx, target)
						if err == nil {
							c.Results <- res
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
	close(c.Results)
}
