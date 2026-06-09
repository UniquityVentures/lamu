# Releasing lamu

Lamu is a multi-module Go repo. Every release uses **one version number** for the core module and all plugins.

## Prerequisites

- [Nushell](https://www.nushell.sh/) (`nu`)
- [fd](https://github.com/sharkdp/fd)
- Go 1.26+
- Push access to `github.com/UniquityVentures/lamu`

For private modules, set:

```bash
export GOPRIVATE=github.com/UniquityVentures/*
export GONOSUMDB=github.com/UniquityVentures/*
```

## Versioning model

One release version applies to every module in the repo:

| Module | Example import | Git tag |
|---|---|---|
| Core | `github.com/UniquityVentures/lamu` | `v0.4.15` |
| Plugin | `github.com/UniquityVentures/lamu/plugins/p_otp` | `plugins/p_otp/v0.4.15` |

`release.nu` creates all of these tags on the same commit. Old per-plugin-only tags can stay on GitHub; do not delete them.

When adding a new plugin, put a `go.mod` under `plugins/<name>/`. `release.nu` and `tidy.nu` pick it up automatically via `fd go.mod`.

## Release procedure

Run from the lamu repo root.

### 1. Commit your changes

```bash
git status
git add ...
git commit -m "Describe the change."
```

### 2. Tidy modules (optional before release)

`release.nu` runs this automatically, but you can run it while developing to fix linter errors across all modules:

```bash
nu tidy.nu
```

Commit any `go.mod` / `go.sum` updates before releasing.

### 3. Push master

```bash
git push origin master
```

### 4. Tag and publish

Pick the next patch version (e.g. `0.4.16` after `v0.4.15`):

```bash
nu release.nu v0.4.16
```

This will:

1. Run `nu tidy.nu` on every module
2. Create `v0.4.16` and `plugins/p_*/v0.4.16` tags on `HEAD`
3. Run `git push --tags`

If `tidy.nu` changes files, stop, commit those changes, push master, and run `release.nu` again.

### 5. Verify

```bash
git tag --sort=-creatordate | head -10
```

Confirm the new `v0.4.x` tag and matching `plugins/p_*/v0.4.x` tags exist locally and on GitHub.

## Updating a deployment

Deployments (totschool, uniquity, nirmancampus, etc.) pin lamu modules in their root `go.mod`. Bump every lamu dependency you use to the same version:

```go
require (
    github.com/UniquityVentures/lamu v0.4.16
    github.com/UniquityVentures/lamu/plugins/p_dashboard v0.4.16
    // ...
)
```

Then in the deployment repo:

```bash
go mod tidy
go build -o /dev/null .
git add go.mod go.sum
git commit -m "Bump lamu to v0.4.16."
git tag -a v0.6.4 -m "Release v0.6.4"
git push origin master
git push origin v0.6.4
```

Deployments have their own version tags (e.g. `v0.6.4` for totschool), separate from lamu.

## Notes

- **Do not add `go.work` to lamu.** Plugins use `replace` directives for local development. A stray `go.work.sum` without `go.work` can be deleted.
- **`git push --tags` pushes all local tags**, not only the release you just created. Avoid leaving experimental tags locally before releasing.
- **CI / fresh clones** need `GOPRIVATE` for `github.com/UniquityVentures/*` modules.
