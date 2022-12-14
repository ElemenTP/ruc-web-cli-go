package rucweb

import (
	"testing"
)

func TestGetxEncode(t *testing.T) {
	msg := `{"username":"201626203044@cmcc","password":"15879684798qq","ip":"10.128.96.249","acid":"1","enc_ver":"srun_bx1"}`
	key := "e6843f26b8544327a3a25978dd3c5f89e6b745df1732993b88fe082c13a34cb9"
	str := GetxEncode(msg, key)
	res := []byte{102, 146, 239, 107, 228, 117, 59, 64, 183, 155, 138, 100, 238, 24, 148, 185, 252, 199, 111, 21, 207, 24, 229, 185, 99, 43, 9, 9, 50, 110, 60, 114, 67, 133, 216, 96, 156, 106, 22, 190, 65, 222, 169, 129, 25, 92, 240, 237, 132, 76, 57, 126, 194, 40, 214, 56, 111, 186, 109, 14, 107, 11, 153, 57, 136, 210, 252, 143, 96, 70, 70, 203, 70, 229, 145, 214, 134, 101, 55, 230, 169, 89, 186, 10, 155, 146, 135, 38, 99, 141, 249, 119, 39, 211, 3, 160, 166, 42, 159, 76, 233, 96, 4, 206, 218, 100, 216, 225, 94, 86, 68, 137, 76, 17, 83, 31}
	for i, v := range []byte(str) {
		if v != res[i] {
			t.Fail()
		}
	}
}
