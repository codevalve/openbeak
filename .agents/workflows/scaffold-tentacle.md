---
description: How to add a new scanning Tentacle to OpenBeak.
---

# Scaffold Tentacle Workflow

This workflow describes how to add a new "Tentacle" module to OpenBeak's scanning and probing engine. Tentacles are individual scanner modules designed to probe specific misconfigurations or vulnerabilities in target OpenClaw deployments. 

All new tentacles must be reviewed by the core OpenBeak team and submitted as Pull Requests.

## 1. Create Tentacle File
New tentacles live in the `internal/tentacles/` directory.
- [ ] Create a new file `internal/tentacles/<tentacle_name>.go`.
- [ ] Set the package to `tentacles`.

## 2. Define Tentacle Structure
// turbo
- [ ] Add the following boilerplate implementing the core interface:
```go
package tentacles

import (
	"context"
)

// Example struct for the new tentacle
type <Name>Tentacle struct {
	// configuration fields
}

func New<Name>Tentacle() *<Name>Tentacle {
	return &<Name>Tentacle{}
}

func (t *<Name>Tentacle) Name() string {
	return "<name>"
}

func (t *<Name>Tentacle) Description() string {
	return "Probes for specific OpenClaw misconfigurations..."
}

func (t *<Name>Tentacle) Role() string {
	return "Hunter" // Must be Hunter, Reporter, or Beak
}

func (t *<Name>Tentacle) Probe(ctx context.Context, target string) (Result, error) {
	// Tentacle logic here: fast, highly concurrent, stealthy network calls
	return Result{}, nil
}
```

## 3. Implement Logic
- [ ] Ensure the network and API calls are context-aware for fast cancellation.
- [ ] Adhere to stealth guidelines: use standard user-agents and throttle requests if necessary.
- [ ] (Optional) Add specific TUI reporting structs if the tentacle yields unique output that the Charm-based interface should render.

## 4. Register Tentacle
- [ ] Register your newly created Tentacle in `internal/tentacles/registry.go` so that the main engine can load it dynamically or selectively during a scan.

## 5. Verification
// turbo
- [ ] Run `go test ./internal/tentacles/...`
- [ ] Create a Pull Request against the `develop` branch for the core team to review.
