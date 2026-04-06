// Command fix_spec fetches the upstream Omada OpenAPI spec, fixes known
// code-generation-blocking issues, and writes tools/all-fixed.json.
//
// Usage:
//
//	go run tools/fix_spec.go
//
// Set OMADA_OPENAPI_URL to override the upstream spec URL.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

const defaultSpecURL = "https://use1-omada-northbound.tplinkcloud.com/v3/api-docs/00%20All"

func main() {
	specURL := os.Getenv("OMADA_OPENAPI_URL")
	if specURL == "" {
		specURL = defaultSpecURL
	}

	log.Printf("fetching spec from %s", specURL)
	raw, err := fetchSpec(specURL)
	if err != nil {
		log.Fatalf("fetch spec: %v", err)
	}

	if err := os.WriteFile("tools/all.json", raw, 0644); err != nil {
		log.Fatalf("write all.json: %v", err)
	}
	log.Printf("wrote tools/all.json (%d bytes)", len(raw))

	// Parse with ordered-key-safe approach: work on raw text to handle
	// duplicate JSON keys (which Go's encoding/json silently deduplicates).
	fixed := fixSpec(raw)

	if err := os.WriteFile("tools/all-fixed.json", fixed, 0644); err != nil {
		log.Fatalf("write all-fixed.json: %v", err)
	}
	log.Printf("wrote tools/all-fixed.json (%d bytes)", len(fixed))
}

func fetchSpec(specURL string) ([]byte, error) {
	parsed, err := url.Parse(specURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}
	if parsed.Scheme != "https" && parsed.Scheme != "http" {
		return nil, fmt.Errorf("unsupported URL scheme: %s", parsed.Scheme)
	}

	resp, err := http.Get(specURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 100*1024*1024)) // 100MB limit
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}
	return body, nil
}

// fixSpec applies all transformations to the raw spec JSON.
// It uses a text-based approach to handle duplicate keys that Go's JSON
// decoder would silently merge.
func fixSpec(raw []byte) []byte {
	text := string(raw)

	// Build a mapping of old schema names to new (sanitized) names.
	nameMap := buildSchemaNameMap(text)

	// Rename schema definitions and all $ref pointers.
	text = renameSchemas(text, nameMap)

	// Fix server URL: http → https (handle any whitespace variations).
	text = strings.ReplaceAll(text,
		`"http://use1-omada-northbound.tplinkcloud.com"`,
		`"https://use1-omada-northbound.tplinkcloud.com"`)

	// Add securitySchemes and global security requirement.
	text = addSecurity(text)

	// Pretty-print the result.
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(text), &parsed); err != nil {
		// If we can't parse (e.g., due to remaining duplicate keys), return as-is.
		log.Printf("warning: could not re-parse JSON for pretty-printing: %v", err)
		return []byte(text)
	}

	// Fix missing path parameters (must operate on parsed structure).
	fixMissingPathParams(parsed)

	// Add placeholder schemas for undefined $ref targets.
	fixUndefinedRefs(parsed)

	// Remove "pattern" fields from schema properties — the Go generator
	// converts these to validate:"regexp=..." struct tags, but regexes
	// containing spaces break Go's struct tag parser.
	removed := stripPatterns(parsed)
	if removed > 0 {
		log.Printf("removed %d pattern fields from schemas", removed)
	}

	// Merge all patch specs into the main spec.
	mergePatchSpecs(parsed)

	out, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		return []byte(text)
	}
	return out
}

// buildSchemaNameMap scans the spec text for schema references and builds a
// mapping from problematic names to sanitized PascalCase names.
func buildSchemaNameMap(text string) map[string]string {
	re := regexp.MustCompile(`#/components/schemas/([^"]+)`)
	matches := re.FindAllStringSubmatch(text, -1)

	seen := make(map[string]bool)
	nameMap := make(map[string]string)

	for _, m := range matches {
		old := m[1]
		if seen[old] {
			continue
		}
		seen[old] = true

		sanitized := sanitizeSchemaName(old)
		if sanitized != old {
			nameMap[old] = sanitized
		}
	}

	// Also find schema definition keys (4-space indented, which are top-level
	// keys under components.schemas in the pretty-printed JSON). This catches
	// schemas that might not be referenced via $ref.
	defRe := regexp.MustCompile(`(?m)^    "([^"]+)":\s*\{`)
	defMatches := defRe.FindAllStringSubmatch(text, -1)
	for _, m := range defMatches {
		old := m[1]
		if seen[old] {
			continue
		}
		seen[old] = true

		sanitized := sanitizeSchemaName(old)
		if sanitized != old {
			nameMap[old] = sanitized
		}
	}

	// Handle collisions: if two different old names map to the same new name,
	// append a suffix to disambiguate.
	reverse := make(map[string][]string)
	for old, nw := range nameMap {
		reverse[nw] = append(reverse[nw], old)
	}
	for nw, olds := range reverse {
		if len(olds) <= 1 {
			continue
		}
		for i, old := range olds {
			if i == 0 {
				continue
			}
			nameMap[old] = fmt.Sprintf("%s%d", nw, i+1)
		}
	}

	// Also check that new names don't collide with existing clean names.
	for old := range seen {
		if _, remapped := nameMap[old]; remapped {
			continue
		}
		// This is a "clean" name. If any remapped name matches it, suffix the remap.
		for mapOld, mapNew := range nameMap {
			if mapNew == old {
				nameMap[mapOld] = mapNew + "Entity"
			}
		}
	}

	return nameMap
}

// sanitizeSchemaName converts a schema name with spaces/special chars to PascalCase.
func sanitizeSchemaName(name string) string {
	// Trim leading/trailing whitespace.
	name = strings.TrimSpace(name)

	// Split on non-alphanumeric characters.
	parts := regexp.MustCompile(`[^a-zA-Z0-9]+`).Split(name, -1)

	var b strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Capitalize first letter of each part.
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		b.WriteString(string(runes))
	}

	result := b.String()
	if result == "" {
		return "Unknown"
	}
	return result
}

// renameSchemas replaces schema names in both definition keys and $ref pointers.
func renameSchemas(text string, nameMap map[string]string) string {
	// Sort by length descending to avoid partial replacements.
	type entry struct {
		old, new string
	}
	var entries []entry
	for old, nw := range nameMap {
		entries = append(entries, entry{old, nw})
	}
	// Sort longest first.
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			if len(entries[j].old) > len(entries[i].old) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}

	for _, e := range entries {
		// Replace $ref pointers.
		text = strings.ReplaceAll(text,
			`#/components/schemas/`+e.old+`"`,
			`#/components/schemas/`+e.new+`"`)

		// Replace schema definition keys.
		// Match the key in quotes at the schema definition level.
		text = strings.ReplaceAll(text,
			`"`+e.old+`":`,
			`"`+e.new+`":`)
	}

	return text
}

// addSecurity injects securitySchemes into components and a global security
// requirement into the spec.
func addSecurity(text string) string {
	// Add securitySchemes to components section.
	securitySchemes := `"securitySchemes": {
      "AccessToken": {
        "type": "apiKey",
        "name": "Authorization",
        "in": "header",
        "description": "Access token obtained from /openapi/authorize/token. Format: AccessToken=<token>"
      }
    }`

	// Find the closing of the "components" object and insert before it.
	// Look for "schemas": { ... } within components. We'll insert after schemas.
	schemasEndIdx := strings.LastIndex(text, `"schemas"`)
	if schemasEndIdx < 0 {
		log.Printf("warning: could not find 'schemas' key to insert securitySchemes")
		return text
	}

	// Find the components closing brace. The components object is typically
	// near the end: "components": { "schemas": { ... } }
	// We'll insert the securitySchemes right before the final components closing brace.
	componentsIdx := strings.Index(text, `"components"`)
	if componentsIdx < 0 {
		log.Printf("warning: could not find 'components' key")
		return text
	}

	// Strategy: find the last `}` before the end of the file (or before a trailing newline).
	// The structure is: { "openapi":..., "paths":..., "components": { "schemas": {...} } }
	// We need to add securitySchemes as a sibling of "schemas" inside "components".

	// Find the position right before the closing `}` of "components".
	// Count braces from "components" to find its closing brace.
	braceCount := 0
	componentEnd := -1
	inString := false
	escaped := false
	for i := componentsIdx; i < len(text); i++ {
		ch := text[i]
		if escaped {
			escaped = false
			continue
		}
		if ch == '\\' && inString {
			escaped = true
			continue
		}
		if ch == '"' {
			inString = !inString
			continue
		}
		if inString {
			continue
		}
		if ch == '{' {
			braceCount++
		} else if ch == '}' {
			braceCount--
			if braceCount == 0 {
				componentEnd = i
				break
			}
		}
	}

	if componentEnd < 0 {
		log.Printf("warning: could not find end of 'components' object")
		return text
	}

	// Insert securitySchemes before the closing brace of components.
	text = text[:componentEnd] + ",\n    " + securitySchemes + "\n  " + text[componentEnd:]

	// Add global security requirement after the components block.
	// Find the closing `}` of the root object and insert before it.
	rootEnd := strings.LastIndex(text, "}")
	if rootEnd < 0 {
		return text
	}

	globalSecurity := `,
  "security": [
    {
      "AccessToken": []
    }
  ]`

	text = text[:rootEnd] + globalSecurity + "\n" + text[rootEnd:]

	return text
}

// fixMissingPathParams finds path parameters used in URL templates that are
// not declared in the operation's parameters list and adds them.
func fixMissingPathParams(spec map[string]interface{}) {
	pathParamRe := regexp.MustCompile(`\{([^}]+)\}`)

	paths, ok := spec["paths"].(map[string]interface{})
	if !ok {
		return
	}

	httpMethods := map[string]bool{
		"get": true, "post": true, "put": true, "delete": true,
		"patch": true, "options": true, "head": true, "trace": true,
	}

	added := 0
	for pathStr, pathItem := range paths {
		// Extract parameter names from the URL template.
		matches := pathParamRe.FindAllStringSubmatch(pathStr, -1)
		if len(matches) == 0 {
			continue
		}

		pathItemMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		// Collect path-level parameters.
		pathLevelParams := getParamNames(pathItemMap["parameters"])

		for method, opValue := range pathItemMap {
			if !httpMethods[method] {
				continue
			}
			op, ok := opValue.(map[string]interface{})
			if !ok {
				continue
			}

			// Collect operation-level parameters.
			opParams := getParamNames(op["parameters"])

			// Merge path-level + operation-level for the full set.
			allDeclared := make(map[string]bool)
			for k, v := range pathLevelParams {
				allDeclared[k] = v
			}
			for k, v := range opParams {
				allDeclared[k] = v
			}

			// Add missing path parameters.
			for _, m := range matches {
				paramName := m[1]
				if allDeclared[paramName] {
					continue
				}

				// Create the missing parameter definition.
				newParam := map[string]interface{}{
					"name":        paramName,
					"in":          "path",
					"required":    true,
					"description": paramName,
					"schema":      map[string]interface{}{"type": "string"},
				}

				params, _ := op["parameters"].([]interface{})
				op["parameters"] = append(params, newParam)
				added++
			}
		}
	}

	if added > 0 {
		log.Printf("added %d missing path parameter definitions", added)
	}
}

// getParamNames extracts path parameter names from a parameters array.
func getParamNames(params interface{}) map[string]bool {
	result := make(map[string]bool)
	arr, ok := params.([]interface{})
	if !ok {
		return result
	}
	for _, p := range arr {
		param, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		if param["in"] == "path" {
			if name, ok := param["name"].(string); ok {
				result[name] = true
			}
		}
	}
	return result
}

// fixUndefinedRefs finds $ref targets in the spec that point to schemas which
// don't exist, and creates empty placeholder schemas for them.
func fixUndefinedRefs(spec map[string]interface{}) {
	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		return
	}
	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		return
	}

	// Collect all $ref targets by walking the entire spec.
	refs := collectRefs(spec)

	added := 0
	for _, ref := range refs {
		const prefix = "#/components/schemas/"
		if !strings.HasPrefix(ref, prefix) {
			continue
		}
		name := strings.TrimPrefix(ref, prefix)
		if _, exists := schemas[name]; !exists {
			schemas[name] = map[string]interface{}{
				"type":        "object",
				"description": "Placeholder for undefined upstream schema",
			}
			added++
			log.Printf("added placeholder schema: %s", name)
		}
	}

	if added > 0 {
		log.Printf("added %d placeholder schemas for undefined refs", added)
	}
}

// stripPatterns recursively removes "pattern" keys from schema property
// objects. The Go generator turns these into validate:"regexp=..." struct tags,
// but regex values containing spaces produce invalid Go struct tags.
func stripPatterns(v interface{}) int {
	removed := 0
	switch val := v.(type) {
	case map[string]interface{}:
		if _, has := val["pattern"]; has {
			delete(val, "pattern")
			removed++
		}
		for _, child := range val {
			removed += stripPatterns(child)
		}
	case []interface{}:
		for _, item := range val {
			removed += stripPatterns(item)
		}
	}
	return removed
}

// mergePatchSpecs reads all specs in generator/patches and merges them into the
// main spec so a single generation pass produces the complete SDK.
//
// Behavior:
// - patch specs deep-merge path operations and overwrite schemas by name.
func mergePatchSpecs(spec map[string]interface{}) {
	patchFiles, err := listPatchSpecFiles("generator/patches")
	if err != nil {
		log.Fatalf("list patch specs: %v", err)
	}

	for _, patchFile := range patchFiles {
		patchSpec, err := readSpecFile(patchFile)
		if err != nil {
			log.Fatalf("read patch spec %s: %v", patchFile, err)
		}
		mergeOneSpec(spec, patchSpec, patchFile, true, true)
	}
}

func listPatchSpecFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".json" && ext != ".yaml" && ext != ".yml" {
			continue
		}
		files = append(files, filepath.Join(dir, entry.Name()))
	}
	sort.Strings(files)
	return files, nil
}

func readSpecFile(path string) (map[string]interface{}, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var spec map[string]interface{}
	switch strings.ToLower(filepath.Ext(path)) {
	case ".json":
		if err := json.Unmarshal(raw, &spec); err != nil {
			return nil, err
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(raw, &spec); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported spec extension: %s", filepath.Ext(path))
	}

	return spec, nil
}

func mergeOneSpec(mainSpec, overlay map[string]interface{}, source string, overwriteSchemas bool, overwritePaths bool) {
	mainPaths, _ := mainSpec["paths"].(map[string]interface{})
	if mainPaths == nil {
		mainPaths = make(map[string]interface{})
		mainSpec["paths"] = mainPaths
	}

	overlayPaths, _ := overlay["paths"].(map[string]interface{})
	pathsAdded := 0
	pathsUpdated := 0
	pathsSkipped := 0
	for path, overlayPathItem := range overlayPaths {
		existingPathItem, exists := mainPaths[path]
		if !exists {
			mainPaths[path] = overlayPathItem
			pathsAdded++
			continue
		}

		if !overwritePaths {
			pathsSkipped++
			continue
		}

		existingMap, okExisting := existingPathItem.(map[string]interface{})
		overlayMap, okOverlay := overlayPathItem.(map[string]interface{})
		if okExisting && okOverlay {
			deepMergeMaps(existingMap, overlayMap)
			pathsUpdated++
			continue
		}

		mainPaths[path] = overlayPathItem
		pathsUpdated++
	}

	mainComponents, _ := mainSpec["components"].(map[string]interface{})
	if mainComponents == nil {
		mainComponents = make(map[string]interface{})
		mainSpec["components"] = mainComponents
	}
	mainSchemas, _ := mainComponents["schemas"].(map[string]interface{})
	if mainSchemas == nil {
		mainSchemas = make(map[string]interface{})
		mainComponents["schemas"] = mainSchemas
	}

	overlayComponents, _ := overlay["components"].(map[string]interface{})
	overlaySchemas, _ := overlayComponents["schemas"].(map[string]interface{})
	schemasAdded := 0
	schemasUpdated := 0
	schemasSkipped := 0
	for name, schema := range overlaySchemas {
		_, exists := mainSchemas[name]
		if exists && !overwriteSchemas {
			schemasSkipped++
			continue
		}
		mainSchemas[name] = schema
		if exists {
			schemasUpdated++
		} else {
			schemasAdded++
		}
	}

	log.Printf("merged spec %s: paths +%d updated %d skipped %d; schemas +%d updated %d skipped %d",
		source, pathsAdded, pathsUpdated, pathsSkipped, schemasAdded, schemasUpdated, schemasSkipped)
}

func deepMergeMaps(dst, src map[string]interface{}) {
	for key, srcVal := range src {
		dstVal, exists := dst[key]
		if !exists {
			dst[key] = srcVal
			continue
		}

		dstMap, okDst := dstVal.(map[string]interface{})
		srcMap, okSrc := srcVal.(map[string]interface{})
		if okDst && okSrc {
			deepMergeMaps(dstMap, srcMap)
			continue
		}

		// Non-map values are overwritten by the overlay value.
		dst[key] = srcVal
	}
}

// collectRefs recursively walks the spec and collects all $ref values.
func collectRefs(v interface{}) []string {
	var refs []string
	switch val := v.(type) {
	case map[string]interface{}:
		if ref, ok := val["$ref"].(string); ok {
			refs = append(refs, ref)
		}
		for _, child := range val {
			refs = append(refs, collectRefs(child)...)
		}
	case []interface{}:
		for _, item := range val {
			refs = append(refs, collectRefs(item)...)
		}
	}
	return refs
}
