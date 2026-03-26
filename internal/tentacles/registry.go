package tentacles

import (
	"github.com/codevalve/openbeak/internal/models"
)

// Global registry of all available tentacles and inks.
var (
	Hunters = []models.Tentacle{
		&HTTPDiscoveryTentacle{Timeout: 2},
	}

	Inks = []models.Ink{
		&JSONInk{FilePath: "openbeak_results.json"},
		&ActivityInk{FilePath: "openbeak_activity.log"},
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

// GetInk returns an ink by name.
func GetInk(name string) models.Ink {
	for _, i := range Inks {
		if i.Name() == name {
			return i
		}
	}
	return nil
}
