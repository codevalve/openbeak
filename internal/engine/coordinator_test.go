package engine

import (
	"context"
	"testing"

	"github.com/codevalve/openbeak/internal/models"
)

type MockTentacle struct {
	severity string
}

func (m *MockTentacle) Name() string        { return "mock" }
func (m *MockTentacle) Role() string        { return "Hunter" }
func (m *MockTentacle) Description() string { return "mock" }
func (m *MockTentacle) Probe(ctx context.Context, target string) (models.Result, error) {
	return models.Result{Severity: m.severity}, nil
}

func TestCoordinator_Inking(t *testing.T) {
	ctx := context.Background()

	// 1. Stealth mode: only High should pass
	coord := NewCoordinator(1)
	coord.Inking = models.InkingStealth
	coord.RegisterTentacle(&MockTentacle{severity: "Medium"})

	results := make(chan models.Result, 10)
	coord.Results = results

	go func() {
		coord.Scan(ctx, []string{"local"})
	}()

	resCount := 0
	for range results {
		resCount++
	}

	if resCount != 0 {
		t.Errorf("Expected 0 results in stealth mode for medium finding, got %d", resCount)
	}

	// 2. Tactical mode: Medium should pass
	coord2 := NewCoordinator(1)
	coord2.Inking = models.InkingTactical
	coord2.RegisterTentacle(&MockTentacle{severity: "Medium"})
	results2 := make(chan models.Result, 10)
	coord2.Results = results2

	go func() {
		coord2.Scan(ctx, []string{"local"})
	}()

	resCount2 := 0
	for range results2 {
		resCount2++
	}

	if resCount2 != 1 {
		t.Errorf("Expected 1 result in tactical mode for medium finding, got %d", resCount2)
	}

	// 3. Verbose mode: Low should pass
	coord3 := NewCoordinator(1)
	coord3.Inking = models.InkingVerbose
	coord3.RegisterTentacle(&MockTentacle{severity: "Low"})
	results3 := make(chan models.Result, 10)
	coord3.Results = results3

	go func() {
		coord3.Scan(ctx, []string{"local"})
	}()

	resCount3 := 0
	for range results3 {
		resCount3++
	}

	if resCount3 != 1 {
		t.Errorf("Expected 1 result in verbose mode for low finding, got %d", resCount3)
	}
}
