package rucweb

import (
	"testing"
)

func TestGetBase64(t *testing.T) {
	str := GetBase64("13245667")
	if str != "9F9x0JHI0kM=" {
		t.Fatalf("%s != 9F9x0JHI0kM=", str)
	}
}
