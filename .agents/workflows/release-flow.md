---
description: Steps for bumping version and release.
---

# Release Flow Workflow

This workflow outlines the standard process for releasing a new version of OpenBeak.

## 1. Preparation
Ensure you are on the `main` branch and it is up-to-date.
// turbo
- [ ] Run `git checkout main && git pull origin main`
- [ ] Run `go test ./...` to ensure all tests pass locally before proceeding.

## 2. Version Bump
Update the version string in the source code.
- [ ] Locate `cmd/root.go` and update `var Version = "X.Y.Z"` to the new version.
- [ ] Update `README.md` or `ROADMAP.md` if necessary to reflect the release.

## 3. Commit and Push
// turbo
- [ ] Run `git add cmd/root.go README.md` (and any other modified doc files)
- [ ] Run `git commit -m "chore: bump version to vX.Y.Z"`
- [ ] Run `git push origin main`

## 4. Tagging
Create a signed git tag for the release. This tag will trigger CI/CD pipelines to build the static binaries.
// turbo
- [ ] Run `git tag -a vX.Y.Z -m "Release vX.Y.Z"`
- [ ] Run `git push origin vX.Y.Z`

## 5. Branch Synchronization
Ensure `develop` branch is updated with the release commit.
// turbo
- [ ] Run `git checkout develop && git merge main && git push origin develop && git checkout main`

## 6. Verification
- [ ] Check GitHub Actions to ensure the `release` workflow completes successfully.
- [ ] Verify the new release appears on GitHub and that the statically compiled binaries for all target platforms (Linux, macOS, Windows) are attached to the release.
