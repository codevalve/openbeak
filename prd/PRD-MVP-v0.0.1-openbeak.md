# PRD — OpenBeak MVP
**Version:** v0.0.1  
**Status:** Draft  
**Owner:** CodeValve  
**Project:** OpenBeak  

---

## 1. Overview

OpenBeak is a defensive security tool designed to detect exposed, misconfigured, or potentially malicious OpenClaw deployments.

OpenClaw agents operate with broad system and network permissions, enabling automation across services, filesystems, and messaging platforms. This creates a growing attack surface when instances are improperly secured or exposed.

**OpenBeak MVP focuses on detection only.**

---

## 2. Problem Statement

The rapid adoption of autonomous AI agents introduces new security risks:

- Exposed agent gateways
- Unauthenticated API access
- Malicious or unverified “skills”
- Hidden automation infrastructure

Security tooling for agent ecosystems is currently:

- Fragmented
- Reactive
- Not purpose-built for agent runtimes

---

## 3. MVP Goal

Deliver a high-performance, concurrent scanning tool that:

- Identifies exposed or misconfigured OpenClaw instances
- Produces structured, actionable findings
- Runs as a single static binary

---

- Remediation / “neutralization” (Reserved for v1.x/v2.x)
- Plugin architecture (Dynamic loading is a v2.x goal)
- Multi-protocol scanning (beyond HTTP/Websockets)
- SIEM/SOAR integrations (Third-party connectors are out of scope for MVP)
- Agent handshake exploitation (MVP is discovery-only)

> [!NOTE]
> While "neutralization" is a non-goal for the MVP, the **Engine** architecture must be designed to support "Beak" (neutralization) modules in future phases.

---

## 5. MVP Design Principles

### 5.1 Stealth-First DNA
OpenBeak is a "White-Hat Predator." Unlike traditional network scanners that are noisy (e.g., `nmap`), OpenBeak must:
- Minimize the footprint on target system logs.
- Use staggered, concurrent request patterns.
- Implement randomized user-agents and low-signature probing.

### 5.2 Aesthetic & UX
- **Charm TUI**: The primary interface is a high-performance terminal UI built on `bubbletea` and `lipgloss`.
- **Hacker-Themed**: Deep blacks, tactical icons (Octicons), and high-contrast styling to reinforce the "stealth hunter" aesthetic.
- **Micro-Animations**: Real-time feedback on "Tentacle" activity and findings.

---

## 6. MVP Scope

### 6.1 Core Components

**Engine (`internal/engine`)**
- Orchestrates scanning
- Manages concurrency
- Aggregates findings

**Tentacle (`internal/tentacles`)**
- HTTP Discovery Tentacle
- Scans for OpenClaw signatures and unauthenticated API endpoints.

**TUI (`internal/tui`)**
- The "Bubble Tea" state machine and rendering logic.
- Visual feedback and progress tracking.

**Output**
- JSON (standard for automation)
- Human-readable (integrated into the TUI flow)

---

### 5.2 Supported Inputs

- Target file (host:port list)
- CIDR range (optional)

---

### 5.3 Supported Detection Types

| Detection        | Description                                      | Severity |
|-----------------|--------------------------------------------------|----------|
| exposed_api     | OpenClaw API accessible without auth             | High     |
| open_endpoint   | Known endpoints respond unexpectedly              | Medium   |
| signature_match | Response matches OpenClaw fingerprint             | Medium   |
| misconfiguration| Indicators of unsafe defaults                     | High     |

---

## 6. Architecture

### 6.1 High-Level Flow

```
Targets → Engine → Worker Pool → HTTP Tentacle → Findings → Aggregator → Output
```

---

### 6.2 Concurrency Model

- Worker pool (configurable)
- Default: 50 concurrent workers
- Channel-based communication

---

### 6.3 Data Models

```go
type Target struct {
    Host string
    Port int
}

type Finding struct {
    Target    string
    Type      string
    Severity  string
    Details   string
    Source    string
    Timestamp time.Time
}
```

---

## 7. Tentacle: HTTP Probe

### Behavior

For each target:

- Attempt requests to:
  - /
  - /health
  - /api
  - /skills

### Detection Signals

- Known response patterns
- Missing authentication
- Open endpoints returning structured data
- Headers or payloads indicating OpenClaw

---

## 8. CLI Interface

### Command

```
openbeak scan --targets targets.txt
```

### Optional

```bash
openbeak scan --cidr 10.0.0.0/24
openbeak scan --workers 100
openbeak scan --output results.json
openbeak scan --stealth-level [1-5]  # Controls scan jitter/noise
```

---

## 9. TUI Experience (Charmbracelet)

The TUI is the heart of the "Macroctopus Agentaculum" experience.

### 9.1 View States
- **Discovery Mode**: Real-time progress bars and live "crevice probing" status.
- **Findings Table**: A list of targets grouped by severity (High/Medium/Low).
- **Detail View**: Detailed information about an exposed endpoint or misconfiguration.

### 9.2 Aesthetic Specs
- **Colors**: Cyber-tactical (Deep Charcoal, Neon Teal, High-Alert Red).
- **Typography**: Optimized for fixed-width terminal fonts.
- **Layout**: Adaptive grid/flex layouts using Lipgloss.

---

## 10. Output Specification

### JSON (Primary)

```json
[
  {
    "target": "10.0.0.12:3000",
    "type": "exposed_api",
    "severity": "high",
    "details": "Unauthenticated /skills endpoint",
    "source": "tentacle:http",
    "timestamp": "2026-03-26T12:00:00Z"
  }
]
```

---

### Text (Secondary)

```
[HIGH] 10.0.0.12:3000 → Unauthenticated OpenClaw endpoint detected
```

---

## 10. Project Structure

```
/openbeak
  /cmd
    scan.go
    root.go
  /internal
    /engine
      coordinator.go
      worker.go
    /tentacles
      http_discovery.go
    /tui
      model.go
      styles.go
      view.go
    /models
      finding.go
  main.go
```

---

## 11. Security & Ethics

- Tool is detection-only
- Must be used only on:
  - Owned systems
  - Authorized environments
- No exploitation or persistence behavior

---

## 12. Success Criteria

MVP is successful when:

- Scans a /24 network in < 60 seconds (baseline)
- Detects known exposed OpenClaw instances
- Produces valid JSON output
- Runs as a single static binary
- Handles failures without crashing

---

## 13. Future Phases

### v0.1.x
- Additional tentacles (ports, websocket, agent handshake)
- Signature library

### v0.2.x
- “Beak” remediation (opt-in only)
- Policy-based enforcement

### v0.3.x
- Plugin system
- SIEM integration

---

## 14. Key Design Principle

Start as a precise detection tool, not a platform.

OpenBeak’s strength comes from:
- speed
- simplicity
- signal quality

Extensibility comes later.
