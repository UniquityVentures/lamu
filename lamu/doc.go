// Package lamu is the application kernel: configuration loading, HTTP server wiring,
// and aggregated plugin registries (pages, views, routes, models, migrations, etc.).
//
// # Plugins and registries
//
// Each plugin exposes a [registry.Pair] of plugin key and [Plugin]. Feature bundles such as
// [Plugin.Pages] return [PluginFeatures] with Entries (named contributions) and optional Patches
// (transform functions keyed like Entries).
//
// [BuildAllRegistries] merges every plugin’s feature callbacks into global immutable registries
// ([RegistryPage], [RegistryView], …). Call it once during startup with the same slice passed to
// [LoadConfigFromFile] / [Start].
//
// # PluginFeatures.Build and patch functions
//
// [PluginFeatures.Build] turns Entries plus Patches into the slice consumed by
// [registry.NewImmutableRegistry]. For each Entry, every Patch with the same Key is applied in
// registration order: entries[i].Value = patchN(… patch2(patch1(entry)) …).
//
// [FillRegistry] merges plugin features incrementally and runs Build after each merge while
// assembling one registry. Every function assigned to a [Plugin] feature field ([Plugin.Pages],
// [Plugin.Views], [Plugin.Routes], [Plugin.DBInitHooks], etc.) must therefore be deterministic and
// repeat-safe: calling it more than once must describe the same logical contributions and must not
// append to package-level state, mutate previously returned entries, or otherwise depend on call
// count.
//
// Patch functions MUST be pure and idempotent:
//
//   - Pure: do not mutate the input value T in place. [PluginFeatures.Build] shallow-clones the
//     Entries slice only; each Pair’s Value is typically a pointer shared across calls. Mutating it
//     leaks state between Build passes and across merges.
//
//   - Idempotent: applying the same patch again to its own output must yield an equivalent result
//     (no duplicated sidebar rows, no double-append). Prefer returning a copy with updated fields,
//     and use stable [components.PageInterface] keys or explicit guards when inserting children.
//
// Entries stored by pointer (common for pages and views) share identity across Build passes until a
// patch replaces them with a new value; helpers like [components.InsertChildAfter] mutate their
// receiver and are unsafe to call during patching unless you operate on a fresh clone.
//
// When T is a small struct passed by value (for example [Route]), each patch receives a copy of the
// registry entry; assigning fields on that copy does not corrupt other registry keys, but patches
// should still return the updated struct explicitly for clarity.
//
// Feature callbacks themselves ([Plugin.Pages]’s closure, etc.) should construct or return a stable
// description of the plugin’s entries and patches. If they use package-level slices for legacy
// registration-style code, those slices must be fully populated before registry assembly and the
// callback must only read them.
//
// Patches that add child components to pages or layers to views are a common source of accidental
// non-idempotency. Give inserted components stable keys, check whether the key already exists before
// inserting, and return a fresh shallow copy of pointer-backed values instead of appending directly
// to the original value.

package lamu
