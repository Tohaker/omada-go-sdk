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
  -w /local \
  -v "${ROOT_DIR}:/local" \
  "${GENERATOR_IMAGE}" batch generator/batch.yaml

echo ""
echo "=== Step 3: Tidy Go modules ==="
go mod tidy

echo ""
echo "=== Step 4: Verify build ==="
go build ./...

echo ""
echo "=== Done ==="
