# AGENTS.md (OpenBeak Agent Blueprint)

Welcome, Comrade Agent! You are helping build **OpenBeak**, a blazing-fast, stealthy white-hat predator written in Go that hunts malicious OpenClaw deployments. 🐙⚡

## 🐙 Project DNA
- **Stealth-First**: Ensure network probing and scanning are low-noise.
- **High Concurrency**: Built with Go to support multiple tentacles probing at once.
- **White-Hat Usage**: Tooling is strictly intended for defensive and authorized scanning operations.
- **Aesthetic**: Uses the **Charmbracelet** toolchain (`bubbletea`, `lipgloss`) for a hacker-themed predator TUI.

---

## 🏗️ Architecture Guide

### 📂 Core Module: `internal/`
- **`tentacles/`**: The individual scanner modules ("tentacles") that probe target OpenClaw networks.
- **`engine/`**: The core concurrent scanner execution and coordination logic.

### 🕹️ TUI: `internal/tui/`
- **`tui.go`**: The `bubbletea.Model` state machine framing the OpenBeak visual feedback.
- **`styles.go`**: All UI styling must be defined here using `lipgloss`. We rely on stealthy, cyber-tactical colors and sharp borders.

### 🛠️ Commands: `cmd/`
- Standard Cobra subcommands.
- Keep `root.go` focused on bootstrapping. New features like scanning pipelines should be specific subcommands.

---

## 🚦 Guidelines for AI Agents

1.  **Keep it "Charm"**: Follow the design aesthetic of `internal/tui/styles.go`.
2.  **Safety First**: When adding new tentacles, provide dry-run modes or scoped boundaries to prevent accidental disruptions on target systems.
3.  **CLI vs TUI**: Ensure critical stealth features available in the TUI are also available as a standalone CLI command for headless automation.
4.  **Conventional Commits**: Use the **Conventional Commits** specification for all commit messages (e.g., `feat:`, `fix:`, `chore:`, `docs:`).
5.  **Issue Lifecycle**: When working on tasks, always create or reference a GitHub issue. Close the issue using the GitHub CLI or API once the task is merged and verified.
6.  **Roadmap & Milestone Sync**: Always ensure `ROADMAP.md` and GitHub milestones are kept in sync whenever feature plans change or releases are completed.
7.  **Bug Fixes**: When a bug is reported, don't start by trying to fix it. Instead, start by writing a test that reproduces the bug in the `tentacles` or `engine` package. Then, fix the bug and prove it with a passing test.

---

## 🧪 Developer Workflow

When developing or testing locally, isolate tests against mock OpenClaw deployments or test sandboxes.

Check `.agents/workflows/` for automated SOPs:
- **`scaffold-tentacle.md`**: How to add a new scanning Tentacle.
- **`release-flow.md`**: Steps for bumping version and release.
- **`triage-issues.md`**: The automated GitHub issue triage and engineering workflow.
