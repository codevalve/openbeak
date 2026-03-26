package tentacles

import (
	"github.com/codevalve/openbeak/internal/models"
)

// Global registry of all available tentacles and reporters.
var (
	Hunters = []models.Tentacle{
		&HTTPDiscoveryTentacle{Timeout: 2},
	}

	Reporters = []models.Reporter{
		&JSONReporter{FilePath: "openbeak_results.json"},
		&ActivityLogger{FilePath: "openbeak_activity.log"},
	}
)

// GetTentacle returns a hunter by name.
func GetTentacle(name string) models.Tentacle {
	for _, h := range Hunters {
		if h.Name() == name {
			return h
		}
	}
	return nil
}

// GetReporter returns a reporter by name.
func GetReporter(name string) models.Reporter {
	for _, r := range Reporters {
		if r.Name() == name {
			return r
		}
	}
	return nil
}
