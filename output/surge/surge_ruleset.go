package surge

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	RuleDomain        = "DOMAIN"
	RuleDomainSuffix  = "DOMAIN-SUFFIX"
	RuleDomainKeyword = "DOMAIN-KEYWORD"
	RuleIPCIDR        = "IP-CIDR"
	RuleIPCIDR6       = "IP-CIDR6"
)

// DomainRule formats a domain rule for a Surge external RULE-SET.
func DomainRule(ruleType string, value string) string {
	return ruleType + "," + value
}

// IPCIDRRule formats an IPv4 or IPv6 prefix for a Surge external RULE-SET.
func IPCIDRRule(cidr string) (string, error) {
	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", fmt.Errorf("parse CIDR %q: %w", cidr, err)
	}
	ruleType := RuleIPCIDR
	if ip.To4() == nil {
		ruleType = RuleIPCIDR6
	}
	return ruleType + "," + cidr + ",no-resolve", nil
}

// SaveRuleSet writes one policy-less Surge rule per line.
func SaveRuleSet(rules []string, outputPath string) error {
	content := strings.Join(rules, "\n")
	if len(rules) > 0 {
		content += "\n"
	}
	return os.WriteFile(outputPath, []byte(content), 0666)
}
