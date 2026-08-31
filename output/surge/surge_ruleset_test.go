package surge

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIPCIDRRule(t *testing.T) {
	tests := []struct {
		name string
		cidr string
		want string
	}{
		{name: "IPv4", cidr: "192.0.2.0/24", want: "IP-CIDR,192.0.2.0/24,no-resolve"},
		{name: "IPv6", cidr: "2001:db8::/32", want: "IP-CIDR6,2001:db8::/32,no-resolve"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IPCIDRRule(tt.cidr)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("IPCIDRRule() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIPCIDRRuleRejectsInvalidCIDR(t *testing.T) {
	if _, err := IPCIDRRule("not-a-cidr"); err == nil {
		t.Fatal("IPCIDRRule() accepted an invalid CIDR")
	}
}

func TestSaveRuleSet(t *testing.T) {
	path := filepath.Join(t.TempDir(), "example.list")
	rules := []string{
		DomainRule(RuleDomain, "www.example.com"),
		DomainRule(RuleDomainSuffix, "example.org"),
	}
	if err := SaveRuleSet(rules, path); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want := "DOMAIN,www.example.com\nDOMAIN-SUFFIX,example.org\n"
	if string(got) != want {
		t.Fatalf("SaveRuleSet() wrote %q, want %q", got, want)
	}
}
