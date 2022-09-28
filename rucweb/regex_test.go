package rucweb

import (
	"regexp"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
	configstr := `var CONFIG = { 
ip : "10.114.51.4"
}`
	regex, err := regexp.Compile(`var CONFIG = (?s:{.*})`)
	if err != nil {
		t.Fatal(err)
	}
	configSubmatch := regex.FindAllStringSubmatch(configstr, -1)
	if len(configSubmatch) == 0 {
		t.Fatalf("no config find\n")
	}
	config := configSubmatch[0][len(configSubmatch[0])-1]
	if !strings.Contains(config, `ip : "10.114.51.4"`) {
		t.Fatalf(`%s does not contain ip`, config)
	}
	regex, err = regexp.Compile(`ip.*?:.*?"(.*?)"`)
	if err != nil {
		t.Fatal(err)
	}
	ipSubmatch := regex.FindAllStringSubmatch(config, -1)
	if len(ipSubmatch) == 0 {
		t.Fatalf("no ip find\n")
	}
	ip := ipSubmatch[0][len(ipSubmatch[0])-1]
	if ip != `10.114.51.4` {
		t.Fatalf(`%s != 10.114.51.4`, ip)
	}
}
