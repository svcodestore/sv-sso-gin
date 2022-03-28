package utils

import "bytes"

func Pad(str string, length int, repStr string) string {
	l := len(str)
	if l >= length {
		return str
	} else {
		repStrLen := len(repStr)
		buf := bytes.NewBufferString(str)
		needRepL := length - l
		for i := 0; i < needRepL/repStrLen; i++ {
			buf.Write([]byte(repStr))
		}
		restL := needRepL % repStrLen
		if restL != 0 {
			buf.Write([]byte(repStr[0:restL]))
		}
		return buf.String()
	}

}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}