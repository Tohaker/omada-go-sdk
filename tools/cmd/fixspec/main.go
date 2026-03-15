package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const defaultSpecURL = "https://use1-omada-northbound.tplinkcloud.com/v3/api-docs/00%20All"

type addParamFix struct {
	opID        string
	paramName   string
	description string
}

type removeParamFix struct {
	opID      string
	paramName string
}

var addFixes = []addParamFix{
	{"retryAddDeviceByMsp", "mspId", "MSP ID"},
	{"retryAddDeviceByMsp", "siteId", "Site ID"},
	{"adoptOneForMsp", "mspId", "MSP ID"},
	{"modifyOltConfigForMsp", "deviceMac", "Device MAC address, like AA-BB-CC-DD-EE-FF"},
	{"getPonPortListMsp", "deviceMac", "Device MAC address, like AA-BB-CC-DD-EE-FF"},
	{"getPortsConfigListForMsp", "deviceMac", "Device MAC address, like AA-BB-CC-DD-EE-FF"},
	{"getConfigForMsp", "deviceMac", "Device MAC address, like AA-BB-CC-DD-EE-FF"},
	{"editConfigForMsp", "deviceMac", "Device MAC address, like AA-BB-CC-DD-EE-FF"},
	{"getUser_1", "userID", "User ID"},
	{"modifyMspUser", "userID", "User ID"},
	{"deleteMspUser", "userID", "User ID"},
	{"getUser", "userID", "User ID"},
	{"modifyUser", "userID", "User ID"},
	{"deleteUser", "userID", "User ID"},
	{"getLanStatus", "omadacId", "Omada ID"},
	{"recoveryPoePort", "omadacId", "Omada ID"},
	{"exportAuthedClientListGlobalByCloudAccess", "siteId", "Site ID"},
	{"disconnectHotspotAuthedClient", "id", "Auth Client ID"},
	{"modifyLocalUser", "id", "Local user ID"},
	{"clearLocalUserDynamicMac", "id", "Local user ID"},
	{"editVoucherGroupPattern", "groupId", "Voucher Group ID"},
	{"modifyOtoNat", "otonatId", "One-to-One NAT ID"},
	{"deleteOtoNat", "otonatId", "One-to-One NAT ID"},
	{"modifyPortForwarding", "portForwardingId", "Port Forwarding ID"},
	{"deletePortForwarding", "portForwardingId", "Port Forwarding ID"},
	{"deleteGroupProfile", "groupType", "Type of group profile"},
	{"deleteGroupProfileTemplate", "groupType", "Type of group profile"},
	{"modifyPolicyRouting", "policyRoutingId", "Policy Routing ID"},
	{"deletePolicyRouting", "policyRoutingId", "Policy Routing ID"},
	{"modifyStaticRouting", "staticRoutingId", "Static Routing ID"},
	{"deleteStaticRouting", "staticRoutingId", "Static Routing ID"},
	{"exportDhcpReservation", "mac", "Client MAC address"},
	{"modifyMailServerTemplate", "mailId", "Mail ID"},
	{"downloadVoiceMail", "voiceId", "Voice Mail ID"},
	{"deleteSwitchLag", "lagId", "LAG ID"},
	{"modifySwitchLag", "lagId", "LAG ID"},
	{"deleteSwitchLagTemplate", "lagId", "LAG ID"},
	{"modifySwitchLagTemplate", "lagId", "LAG ID"},
	{"setPortModeForGivenPort", "port", "Port ID"},
	{"setProfileForGivenPort", "port", "Port ID"},
	{"setProfileOverrideForGivenPort", "port", "Port ID"},
	{"setPoeModeForGivenPort", "port", "Port ID"},
	{"setNameForGivenPort", "port", "Port ID"},
	{"setPoeModeForGivenPort_1", "port", "Port ID"},
	{"deleteLockedSslVpnTunnel", "lockTunnelId", "lock Tunnel Id"},
	{"countWlans", "siteId", "Site ID"},
	{"countWlansTemplate", "siteTemplateId", "Site Template ID"},
	{"updateSsidHotspotV2SettingTemplate", "siteTemplateId", "Site Template ID"},
	{"updateSsidHotspotV2SettingTemplate", "wlanId", "WLAN ID"},
	{"updateSsidHotspotV2SettingTemplate", "ssidId", "SSID ID"},
	{"updateSsidHotspotV2Setting", "siteId", "Site ID"},
	{"updateSsidHotspotV2Setting", "wlanId", "WLAN ID"},
	{"updateSsidHotspotV2Setting", "ssidId", "SSID ID"},
	{"modifyAnomalyEventSettingForSiteTemplate", "siteTemplateId", "Site Template ID"},
	{"createIotServerTemplate", "siteTemplateId", "Site Template ID"},
	{"querySimCardQuotaSettingTemplate", "siteTemplateId", "Site Template ID"},
	{"querySmsPolicySettingTemplate", "siteTemplateId", "Site Template ID"},
	{"modifySmsPolicySettingTemplate", "siteTemplateId", "Site Template ID"},
	{"querySmsRouterCommandTemplate", "siteTemplateId", "Site Template ID"},
	{"modifySmsRouterCommandTemplate", "siteTemplateId", "Site Template ID"},
	{"getGatewayInfo", "omadacId", "Omada ID"},
	{"modifyConfigAdvancedTemplate", "omadacId", "Omada ID"},
	{"modifyConfigWirelessAdvancedTemplate", "omadacId", "Omada ID"},
	{"modifyConfigCommonAdvancedTemplate", "omadacId", "Omada ID"},
	{"modifyConfigGeneralTemplate", "omadacId", "Omada ID"},
	{"getPortsTemplate", "omadacId", "Omada ID"},
}

var removeFixes = []removeParamFix{
	{"mspMoveToCustomer_1", "deviceMac"},
	{"getUser_1", "omadacId"},
	{"modifyMspUser", "omadacId"},
	{"deleteMspUser", "omadacId"},
}

func main() {
	rawOut := "all.json"
	out := "all-fixed.json"
	specURL := os.Getenv("OMADA_OPENAPI_URL")
	if specURL == "" {
		specURL = defaultSpecURL
	}

	raw, err := fetchSpec(specURL)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(rawOut, raw, 0644); err != nil {
		panic(err)
	}

	c := string(raw)

	c = strings.ReplaceAll(c, "#/components/schemas/Wan/Lan infos", "#/components/schemas/OsgWanStatusVO")
	wildcardRefMediaType := regexp.MustCompile(`"content"\s*:\s*\{\s*"\*/\*"\s*:\s*\{\s*"schema"\s*:\s*\{\s*"\$ref"\s*:`)
	wildcardCount := len(wildcardRefMediaType.FindAllStringIndex(c, -1))
	c = wildcardRefMediaType.ReplaceAllLiteralString(c, `"content":{"application/json":{"schema":{"$ref":`)
	removeGatewayXGoName := regexp.MustCompile(`(?s)"Gateway Info"\s*:\s*\{\s*"x-go-name"\s*:\s*"GatewayInfoTopology"\s*,`)
	c = removeGatewayXGoName.ReplaceAllString(c, `"Gateway Info":{`)
	addGatewayXGoName := regexp.MustCompile(`"Gateway Info"\s*:\s*\{`)
	c = addGatewayXGoName.ReplaceAllString(c, `"Gateway Info":{"x-go-name":"GatewayInfoTopology",`)

	added := 0
	removed := 0
	for _, f := range addFixes {
		var changed bool
		c, changed = addPathParam(c, f.opID, f.paramName, f.description)
		if changed {
			added++
		}
	}
	for _, f := range removeFixes {
		var changed bool
		c, changed = removePathParam(c, f.opID, f.paramName)
		if changed {
			removed++
		}
	}

	var autoCount int
	c, autoCount = autoAddOmadacID(c)

	var renameCount int
	c, renameCount = renameOperationIds(c)

	if err := os.WriteFile(out, []byte(c), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("FETCHED_URL=%s\n", specURL)
	fmt.Printf("WROTE_RAW=%s\n", rawOut)
	fmt.Printf("WROTE=%s\n", out)
	fmt.Printf("EXPLICIT_ADDS_APPLIED=%d\n", added)
	fmt.Printf("EXPLICIT_REMOVES_APPLIED=%d\n", removed)
	fmt.Printf("AUTO_ADDED_OMADACID=%d\n", autoCount)
	fmt.Printf("WILDCARD_CONTENT_REWRITES=%d\n", wildcardCount)
	fmt.Printf("OPERATION_IDS_RENAMED=%d\n", renameCount)
}

func fetchSpec(specURL string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, specURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch spec: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("fetch spec: http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read spec body: %w", err)
	}
	if len(b) == 0 {
		return nil, fmt.Errorf("read spec body: empty response")
	}

	return b, nil
}

func addPathParam(content, opID, paramName, description string) (string, bool) {
	pattern := `(?s)("operationId"\s*:\s*"` + regexp.QuoteMeta(opID) + `"\s*,\s*"parameters"\s*:\s*\[)(.*?)(\]\s*,)`
	re := regexp.MustCompile(pattern)
	changed := false

	out := re.ReplaceAllStringFunc(content, func(m string) string {
		if changed {
			return m
		}
		parts := re.FindStringSubmatch(m)
		if len(parts) != 4 {
			return m
		}
		prefix := parts[1]
		body := parts[2]
		suffix := parts[3]

		hasParamRe := regexp.MustCompile(`"name"\s*:\s*"` + regexp.QuoteMeta(paramName) + `"\s*,\s*"in"\s*:\s*"path"`)
		if hasParamRe.MatchString(body) {
			return m
		}

		trimmed := strings.TrimRight(body, " \t\r\n")
		if len(trimmed) > 0 && !strings.HasSuffix(trimmed, ",") {
			trimmed += ","
		}

		insertion := fmt.Sprintf(`
          {
            "name": %q,
            "in": "path",
            "description": %q,
            "required": true,
            "schema": { "type": "string" }
          }
`, paramName, description)

		changed = true
		return prefix + trimmed + insertion + "        " + suffix
	})

	return out, changed
}

func removePathParam(content, opID, paramName string) (string, bool) {
	pattern := `(?s)("operationId"\s*:\s*"` + regexp.QuoteMeta(opID) + `"\s*,\s*"parameters"\s*:\s*\[)(.*?)(\]\s*,)`
	re := regexp.MustCompile(pattern)
	changed := false

	out := re.ReplaceAllStringFunc(content, func(m string) string {
		if changed {
			return m
		}
		parts := re.FindStringSubmatch(m)
		if len(parts) != 4 {
			return m
		}
		prefix := parts[1]
		body := parts[2]
		suffix := parts[3]

		baseObj := `(?s)\{\s*"name"\s*:\s*"` + regexp.QuoteMeta(paramName) + `"\s*,\s*"in"\s*:\s*"path"\s*,.*?\}`
		newBody := body

		leadComma := regexp.MustCompile(`(?s),\s*` + baseObj)
		if leadComma.MatchString(newBody) {
			newBody = leadComma.ReplaceAllString(newBody, "")
		} else {
			trailComma := regexp.MustCompile(baseObj + `\s*,`)
			if trailComma.MatchString(newBody) {
				newBody = trailComma.ReplaceAllString(newBody, "")
			} else {
				objOnly := regexp.MustCompile(baseObj)
				newBody = objOnly.ReplaceAllString(newBody, "")
			}
		}
		if newBody != body {
			changed = true
		}

		return prefix + newBody + suffix
	})

	return out, changed
}

func summaryToCamelCase(summary string) string {
	splitRe := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	words := splitRe.Split(summary, -1)
	var b strings.Builder
	first := true
	for _, w := range words {
		if w == "" {
			continue
		}
		if first {
			b.WriteString(strings.ToLower(w))
			first = false
		} else {
			b.WriteString(strings.ToUpper(w[:1]))
			if len(w) > 1 {
				b.WriteString(strings.ToLower(w[1:]))
			}
		}
	}
	return b.String()
}

func renameOperationIds(content string) (string, int) {
	opPat := regexp.MustCompile(`"operationId"\s*:\s*"([^"]+)"`)
	sumPat := regexp.MustCompile(`"summary"\s*:\s*"([^"]+)"`)

	opMatches := opPat.FindAllStringSubmatchIndex(content, -1)
	sumMatches := sumPat.FindAllStringSubmatchIndex(content, -1)

	// Collect summary positions
	type posText struct {
		pos  int
		text string
	}
	sums := make([]posText, 0, len(sumMatches))
	for _, sm := range sumMatches {
		sums = append(sums, posText{sm[0], content[sm[2]:sm[3]]})
	}

	// For each operationId, find the nearest preceding summary
	type entry struct {
		start, end int
		oldId      string
		summary    string
	}
	var entries []entry

	for _, om := range opMatches {
		opPos := om[0]
		oldId := content[om[2]:om[3]]

		bestSum := ""
		for _, s := range sums {
			if s.pos < opPos {
				bestSum = s.text
			} else {
				break
			}
		}
		if bestSum == "" {
			continue
		}

		entries = append(entries, entry{
			start:   om[2],
			end:     om[3],
			oldId:   oldId,
			summary: bestSum,
		})
	}

	// Generate camelCase names from summaries
	baseNames := make([]string, len(entries))
	for i, e := range entries {
		baseNames[i] = summaryToCamelCase(e.summary)
	}

	// Deduplicate: first occurrence keeps base name, subsequent get numeric suffix
	finalNames := make([]string, len(entries))
	used := map[string]bool{}
	for i, name := range baseNames {
		if !used[name] {
			finalNames[i] = name
			used[name] = true
		} else {
			n := 2
			for {
				candidate := fmt.Sprintf("%s%d", name, n)
				if !used[candidate] {
					finalNames[i] = candidate
					used[candidate] = true
					break
				}
				n++
			}
		}
	}

	// Replace from end to start so earlier positions stay valid
	count := 0
	for i := len(entries) - 1; i >= 0; i-- {
		e := entries[i]
		newId := finalNames[i]
		if newId != "" && newId != e.oldId {
			content = content[:e.start] + newId + content[e.end:]
			count++
		}
	}

	return content, count
}

func autoAddOmadacID(content string) (string, int) {
	pat := regexp.MustCompile(`(?s)("operationId"\s*:\s*"[^"]+"\s*,\s*"parameters"\s*:\s*\[)(.*?)(\]\s*,)`)
	count := 0

	out := pat.ReplaceAllStringFunc(content, func(m string) string {
		parts := pat.FindStringSubmatch(m)
		if len(parts) != 4 {
			return m
		}
		prefix := parts[1]
		body := parts[2]
		suffix := parts[3]

		hasSite := regexp.MustCompile(`"name"\s*:\s*"siteTemplateId"\s*,\s*"in"\s*:\s*"path"`).MatchString(body)
		hasDevice := regexp.MustCompile(`"name"\s*:\s*"deviceTemplateId"\s*,\s*"in"\s*:\s*"path"`).MatchString(body)
		hasOmada := regexp.MustCompile(`"name"\s*:\s*"omadacId"\s*,\s*"in"\s*:\s*"path"`).MatchString(body)
		if !(hasSite && hasDevice) || hasOmada {
			return m
		}

		count++
		b := strings.TrimRight(body, " \t\r\n")
		if !strings.HasSuffix(b, ",") {
			b += ","
		}

		ins := `
          {
            "name": "omadacId",
            "in": "path",
            "description": "Omada ID",
            "required": true,
            "schema": { "type": "string" }
          }
`
		return prefix + b + ins + suffix
	})

	return out, count
}
