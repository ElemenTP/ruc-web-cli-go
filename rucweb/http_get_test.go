package rucweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	addr := "127.0.0.1:65533"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("URL: %s\nHEADER: %s\n", r.RequestURI, fmt.Sprintln(r.Header))))
	})
	go http.ListenAndServe(addr, nil)
	params := map[string]string{
		"username": "username",
		"password": "password",
		"ip":       "ip",
		"acid":     "c_acid",
		"enc_ver":  "c_enc_ver",
	}
	res, err := HttpGet("http://"+addr, params, HEADERS)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
