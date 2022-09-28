package rucweb

import (
	"testing"
)

func Test_build_login_params(t *testing.T) {
	res := buildLoginParams("username", "password", "10.114.51.4", "1919810")
	fail := false
	if res["password"] != `{MD5}f75dbb40d096c30d7ec3c8900fab54ba` {
		t.Logf("'password' %s != {MD5}f75dbb40d096c30d7ec3c8900fab54ba\n", res["password"])
		fail = true
	}
	if res["chksum"] != `50c23b83b379c5f2f52b481deb53825f8eb10b5b` {
		t.Logf("'chksum' %s != 50c23b83b379c5f2f52b481deb53825f8eb10b5b\n", res["chksum"])
		fail = true
	}
	if res["info"] != `{SRBX1}1LVmRb1TkcX3RRMzSIfY30wECm2peDN5+LhPEKdAbcrxaWf/Au/CluUa9/ofCVtQo5B1Oa52uHf9UEoIYsdeWzDhMT6BTQIIvXQrMGbFCIvdFodNW3CJH438JLApcJQ52HLlmS==` {
		t.Logf("'info' %s != {SRBX1}1LVmRb1TkcX3RRMzSIfY30wECm2peDN5+LhPEKdAbcrxaWf/Au/CluUa9/ofCVtQo5B1Oa52uHf9UEoIYsdeWzDhMT6BTQIIvXQrMGbFCIvdFodNW3CJH438JLApcJQ52HLlmS==\n", res["info"])
		fail = true
	}
	if fail {
		t.Fail()
	}
}
