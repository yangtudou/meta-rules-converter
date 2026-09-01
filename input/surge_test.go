package input

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/metacubex/mihomo/component/geodata/router"
	"google.golang.org/protobuf/proto"
)

func TestConvertSiteToSurge(t *testing.T) {
	attribute := &router.Domain_Attribute{Key: "ads"}
	list := &router.GeoSiteList{Entry: []*router.GeoSite{
		{
			CountryCode: "TEST",
			Domain: []*router.Domain{
				{Type: router.Domain_Full, Value: "www.example.com", Attribute: []*router.Domain_Attribute{attribute}},
				{Type: router.Domain_Domain, Value: "example.org"},
				{Type: router.Domain_Plain, Value: "example-keyword"},
				{Type: router.Domain_Regex, Value: `^api[0-9]+\.example\.net$`, Attribute: []*router.Domain_Attribute{attribute}},
			},
		},
	}}

	inputPath := filepath.Join(t.TempDir(), "geosite.dat")
	writeProtoFile(t, inputPath, list)
	outputDir := filepath.Join(t.TempDir(), "geosite")
	if err := ConvertSite(nil, inputPath, "surge", outputDir, false); err != nil {
		t.Fatal(err)
	}

	assertFileContent(t, filepath.Join(outputDir, "test.list"),
		"DOMAIN,www.example.com\n"+
			"DOMAIN-SUFFIX,example.org\n"+
			"DOMAIN-KEYWORD,example-keyword\n")
	assertFileContent(t, filepath.Join(outputDir, "test@ads.list"),
		"DOMAIN,www.example.com\n")
}

func TestConvertIPToSurge(t *testing.T) {
	list := &router.GeoIPList{Entry: []*router.GeoIP{
		{
			CountryCode: "TEST",
			Cidr: []*router.CIDR{
				{Ip: net.ParseIP("192.0.2.0").To4(), Prefix: 24},
				{Ip: net.ParseIP("2001:db8::").To16(), Prefix: 32},
			},
		},
	}}

	inputPath := filepath.Join(t.TempDir(), "geoip.dat")
	writeProtoFile(t, inputPath, list)
	outputDir := filepath.Join(t.TempDir(), "geoip")
	if err := ConvertIP(nil, inputPath, "surge", outputDir); err != nil {
		t.Fatal(err)
	}

	assertFileContent(t, filepath.Join(outputDir, "test.list"),
		"IP-CIDR,192.0.2.0/24,no-resolve\n"+
			"IP-CIDR6,2001:db8::/32,no-resolve\n")
}

func writeProtoFile(t *testing.T, path string, message proto.Message) {
	t.Helper()
	data, err := proto.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
}

func assertFileContent(t *testing.T, path string, want string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("%s contains %q, want %q", path, got, want)
	}
}
