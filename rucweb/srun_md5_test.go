package rucweb

import "testing"

func TestGetMD5(t *testing.T) {
	password := "15879684798qq"
	token := "711ab370231392679fe06523b119a8fe096f5ed9bd206b4de8d7b5b994bbc3e5"
	str, err := GetMD5(password, token)
	if err != nil {
		t.Fatal(err)
	}
	if str != "b7cc5da95734d0161fadc8ad87855e75" {
		t.Fatalf("%s != b7cc5da95734d0161fadc8ad87855e75", str)
	}
}
