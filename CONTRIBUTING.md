# Contributing to OpenBeak 🐙⚡

Welcome, hunter! We're thrilled you want to help build **OpenBeak**, the predator of malicious OpenClaw deployments.

By contributing to this project, you agree to abide by our White-Hat philosophy: all tooling is intended strictly for authorized security operations and defensive research.

---

## 🛠️ Getting Started

1.  **Prerequisites**: You'll need **Go 1.21+** installed.
2.  **Fork & Clone**: Fork the repository and clone it locally.
3.  **Install Dependencies**: Run `go mod download`.
4.  **Verify Setup**: Run `go test ./...` to ensure everything is passing.

## 🐙 Adding New Tentacles

One of the best ways to contribute is by adding new **Tentacles** (scanner modules). 

We have a specialized workflow for this:
- See [.agents/workflows/scaffold-tentacle.md](.agents/workflows/scaffold-tentacle.md) for a step-by-step guide and boilerplate code.

## 🕹️ TUI & Aesthetic Guidelines

OpenBeak uses the **Charmbracelet** stack (`bubbletea`, `lipgloss`) for its terminal interface.
- **Stealth First**: Stick to the cyber-tactical theme defined in `internal/tui/styles.go`.
- **Concurrency**: Ensure TUI updates are thread-safe, especially when reporting results from multiple concurrent tentacles.

## 🚦 Development Workflow

### 📋 Issue Lifecycle
- **Search First**: Check if an issue already exists for your idea or bug.
- **Create an Issue**: Before starting major work, create an issue to discuss the plan.
- **Reference Issues**: Link your Pull Request to the relevant issue.

### 📝 Conventional Commits
We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:
- `feat:` for new "tentacles" or features.
- `fix:` for bug fixes.
- `chore:` for maintenance/scaffolding.
- `docs:` for documentation updates.

### 🧪 Testing
- **Bug Fixes**: Start by writing a test that reproduces the bug. Fix the bug, then prove it with the passing test.
- **New Features**: All new engine or tentacle logic must include unit tests.

## 🚀 Pull Request Process

1.  Create a branch from `develop` (e.g., `feat/new-tentacle-name`).
2.  Ensure `go weight` and `go fmt` pass.
3.  Submit your PR against the `develop` branch.
4.  The core team will review your contribution for stealth, safety, and performance.

---

**Thank you for helping us hunt!** 🐙
