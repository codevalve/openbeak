# OpenBeak Tentacles Architecture 🐙⚙️

OpenBeak is designed to be highly modular. The scanning engine (`Coordinator`) orchestrates worker pools that utilize plug-in modules called **Tentacles**.

There are two primary distinct categories of Tentacles:

## 1. Hunters (Discovery & Fingerprinting)
Hunters are responsible for interacting with the target to discover misconfigurations, specific services, or OpenClaw deployments. 

Every Hunter implements the `models.Tentacle` interface:
```go
type Tentacle interface {
	Name() string
	Role() string
	Description() string
	Probe(ctx context.Context, target string) (Result, error)
}
```

*Example Hunters:*
- `http_discovery`: Probes for exposed OpenClaw web instances by checking common endpoints and version headers.

## 2. Inks (Data Export & Logging)
Inks are responsible for taking the findings outputted by Hunters and formatting, saving, or exporting them to a destination (e.g., a file, a database, or a SIEM).

Every Ink implements the `models.Ink` interface:
```go
type Ink interface {
	Name() string
	Description() string
	Write(ctx context.Context, result Result) error
	Log(ctx context.Context, event ActivityEvent) error
}
```

*Example Inks:*
- `activity_ink`: Generates a human-readable stream of operational engine events and findings.
- `json_ink`: Dumps highly-structured `models.Result` objects into a JSON array for machine parsing.

## 3. Beaks (Future Extension)
*Roadmap Phase 3:* Beaks will be automated remediation modules capable of neutralizing a rogue agent instance or patching a configuration on the fly.
