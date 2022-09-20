package rucweb

import "testing"

func TestGetSHA1(t *testing.T) {
	str, err := GetSHA1("123456")
	if err != nil {
		t.Fatal(err)
	}
	if str != "7c4a8d09ca3762af61e59520943dc26494f8941b" {
		t.Fatalf("%s != 7c4a8d09ca3762af61e59520943dc26494f8941b", str)
	}
}
