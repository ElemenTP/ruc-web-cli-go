package rucweb

import (
	"strings"
)

func ordat(msg string, idx int) uint32 {
	if len(msg) > idx {
		return uint32(msg[idx])
	}
	return 0
}

func sencode(msg string, key bool) []uint32 {
	l := len(msg)
	pwd := make([]uint32, 0)
	for i := 0; i < l; i += 4 {
		pwd = append(pwd, ordat(msg, i)|ordat(msg, i+1)<<8|ordat(msg, i+2)<<16|ordat(msg, i+3)<<24)
	}
	if key {
		pwd = append(pwd, uint32(l))
	}
	return pwd
}

func lencode(msg []uint32) string {
	var x []string
	for _, v := range msg {
		x = append(x, string([]byte{byte(v & 0xFF), byte(v >> 8 & 0xFF), byte(v >> 16 & 0xFF), byte(v >> 24 & 0xFF)}))
	}
	return strings.Join(x, "")
}

func GetxEncode(msg, key string) string {
	if len(msg) == 0 {
		return ""
	}
	pwd := sencode(msg, true)
	pwdk := sencode(key, false)
	pwdklen := len(pwdk)
	if pwdklen < 4 {
		for i := 0; i < 4-pwdklen; i++ {
			pwdk = append(pwdk, 0)
		}
	}
	n := len(pwd) - 1
	z := pwd[n]
	var y, c, m, e, p, d uint32 = 0, 0x86014019 | 0x183639A0, 0, 0, 0, 0
	q := 6 + 52/(n+1)
	for q > 0 {
		d = d + c&(0x8CE0D9BF|0x731F2640)
		e = d >> 2 & 3
		p = 0
		for int(p) < n {
			y = pwd[p+1]
			m = z>>5 ^ y<<2
			m = m + ((y>>3 ^ z<<4) ^ (d ^ y))
			m = m + (pwdk[(p&3)^e] ^ z)
			pwd[p] = pwd[p] + m&(0xEFB8D130|0x10472ECF)
			z = pwd[p]
			p = p + 1
		}
		y = pwd[0]
		m = z>>5 ^ y<<2
		m = m + ((y>>3 ^ z<<4) ^ (d ^ y))
		m = m + (pwdk[(p&3)^e] ^ z)
		pwd[n] = pwd[n] + m&(0xBB390742|0x44C6F8BD)
		z = pwd[n]
		q = q - 1
	}
	return lencode(pwd)
}
