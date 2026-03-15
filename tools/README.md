# Tools

This directory exists to keep code-generation tooling separate from the main SDK module.

Why this is separate:

- The root module stays focused on the SDK runtime and generated client code.
- Tooling dependencies (for example, oapi-codegen) are isolated in `tools/go.mod`.
- Regeneration can be done deterministically from one place.

## What this folder does

`tools/tools.go` defines the generation workflow with `go:generate`:

1. Fetch and normalize the upstream OpenAPI spec (`cmd/fixspec`), writing:

- `tools/all.json`
- `tools/all-fixed.json`

2. Run `oapi-codegen` using `tools/config.client.yaml`.
3. Generate the SDK client into:

- `client/client.go`

`cmd/fixspec` is a temporary compatibility layer. It applies deterministic fixes to the upstream Omada OpenAPI output so code generation can succeed consistently. This should be removed once the upstream OpenAPI spec is corrected.

Reference discussion on upstream spec issues:

- https://community.tp-link.com/en/business/forum/topic/859880?page=1

## Regenerate code

From repository root:

```bash
cd tools
go generate ./...
```

Optional: override the upstream spec URL for `fixspec`:

```bash
OMADA_OPENAPI_URL="https://your-spec-url" go generate ./...
```

## PR and CI expectations

- Generated files are source-of-truth in this repo and should be committed with the change.
- CI runs formatting, vet, tests, and module tidy checks for both root and tools modules.
- Before opening a PR, run regeneration and verify there is no unintended drift.
