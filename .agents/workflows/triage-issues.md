---
description: Pull latest GitHub issues and evaluate them against project philosophy.
---

# Issue Triage & Engineering Workflow

## 1. Automated Triage (Autopilot)
The `.github/workflows/autopilot.yml` action runs daily and automatically:
- [x] Checks for new issues.
- [x] Labeled new issues with `triage/needed`.
- [ ] For manual triage, run `gh issue list --label triage/needed --limit 10`.

## 2. Retrieve Latest Issues
// turbo
- [ ] Run `gh issue list --limit 10` to see the most recent activity.
- [ ] For each new ID, run `gh issue view <id>` to understand the full context.

## 2. Philosophical Evaluation
Evaluate every issue against **OpenBeak's Core Tenets**:
- [ ] **White-Hat First:** Does this feature respect legal boundaries and is it intended for authorized security operations?
- [ ] **Stealth & Concurrency:** Will this change degrade the blazing-fast Go performance or increase our noise signature on a network?
- [ ] **TUI Aesthetic:** If this impacts the UI, does it maintain the stealthy, hacker-themed predator aesthetic built on the Charm tooling stack?

Against **ROADMAP.md**:
- [ ] Should this feature be slated for upcoming releases, or does it belong in v2.0 (e.g., MCP Support)?
- [ ] Does it align with the core vision of "reaching where claws hide"?

## 3. Engineering Plan
For each triage action:
- [ ] Draft a `plan.md` artifact with:
    - **Context**: Summary of the issue.
    - **Technical Approach**: Specific tentacles or CLI packages to modify.
    - **Test Cases**: New test requirements ensuring concurrent safety.
- [ ] Respond to the issue on GitHub using `gh issue comment <id> -b "..."` with the proposed plan.

## 4. Work Dispatch
- [ ] Update `ROADMAP.md` if the issue is high-priority for the next release.
- [ ] Ask the USER to approve the plan before execution.
- [ ] **Closing**: After merging the fix to `main`, close the issue using `gh issue close <id> --comment "Fixed in vX.Y.Z"`.

## 5. Environment Check
// turbo
- [ ] Ensure local git config is correct:
      `git config user.email "john.lovell@codevalve.com" && git config user.name "codevalve"`
