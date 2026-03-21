#!/usr/bin/env bash
#
# generate.sh — Fetch + fix the Omada OpenAPI spec and generate the Go SDK.
#
# Prerequisites:
#   - Go 1.21+
#   - Docker (for openapi-generator-cli)
#
# Usage:
#   ./generator/generate.sh
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
cd "${ROOT_DIR}"

GENERATOR_IMAGE="openapitools/openapi-generator-cli:v7.12.0"

echo "=== Step 1: Fetch and fix upstream spec ==="
go run tools/fix_spec.go

echo ""
echo "=== Step 2: Generate SDK from fixed spec ==="
docker run --rm \
  -v "${ROOT_DIR}:/local" \
  "${GENERATOR_IMAGE}" generate \
  -i /local/tools/all-fixed.json \
  -g go \
  -o /local/omada \
  -c /local/generator/config.json \
  --git-user-id Tohaker \
  --git-repo-id omada-go-sdk

echo ""
echo "=== Step 3: Tidy Go modules ==="
go mod tidy

echo ""
echo "=== Step 4: Verify build ==="
go build ./...

echo ""
echo "=== Done ==="
