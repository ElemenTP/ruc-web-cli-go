package rucweb

import (
	"bytes"
	"fmt"
)

const _PADCHAR = '='
const _ALPHA = "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"

func getByte(str string, index int) int {
	return int(str[index])
}

func GetBase64(str string) string {
	if len(str) == 0 {
		return str
	}
	i, b10 := 0, 0
	var x bytes.Buffer
	imax := len(str) - len(str)%3
	for i = 0; i < imax; i += 3 {
		b10 = (getByte(str, i) << 16) | (getByte(str, i+1) << 8) | getByte(str, i+2)
		x.WriteByte(_ALPHA[(b10 >> 18)])
		x.WriteByte(_ALPHA[((b10 >> 12) & 63)])
		x.WriteByte(_ALPHA[((b10 >> 6) & 63)])
		x.WriteByte(_ALPHA[(b10 & 63)])
	}
	i = imax
	switch len(str) - imax {
	case 1:
		b10 = getByte(str, i) << 16
		x.WriteString(fmt.Sprintf("%c%c%c%c", _ALPHA[(b10>>18)], _ALPHA[((b10>>12)&63)], _PADCHAR, _PADCHAR))
	case 2:
		b10 = (getByte(str, i) << 16) | (getByte(str, i+1) << 8)
		x.WriteString(fmt.Sprintf("%c%c%c%c", _ALPHA[(b10>>18)], _ALPHA[((b10>>12)&63)], _ALPHA[((b10>>6)&63)], _PADCHAR))
	}
	return x.String()
}
