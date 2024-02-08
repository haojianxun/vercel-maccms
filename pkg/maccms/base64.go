package maccms

import "strings"

var base64EncodeChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var base64DecodeChars = []int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, 64, -1, -1,
	-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, -1, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1,
}

func Base64encode(str string) string {
	var out strings.Builder
	var c1, c2, c3 int

	length := len(str)
	i := 0
	for i < length {
		c1 = int(str[i]) & 0xff
		i++
		if i == length {
			out.WriteByte(base64EncodeChars[c1>>2])
			out.WriteByte(base64EncodeChars[(c1&0x3)<<4])
			out.WriteString("==")
			break
		}
		c2 = int(str[i])
		i++
		if i == length {
			out.WriteByte(base64EncodeChars[c1>>2])
			out.WriteByte(base64EncodeChars[(c1&0x3)<<4|((c2&0xF0)>>4)])
			out.WriteByte(base64EncodeChars[(c2&0xF)<<2])
			out.WriteByte('=')
			break
		}
		c3 = int(str[i])
		i++
		out.WriteByte(base64EncodeChars[c1>>2])
		out.WriteByte(base64EncodeChars[(c1&0x3)<<4|((c2&0xF0)>>4)])
		out.WriteByte(base64EncodeChars[((c2&0xF)<<2)|((c3&0xC0)>>6)])
		out.WriteByte(base64EncodeChars[c3&0x3F])
	}
	return out.String()
}

func Base64decode(str string) string {
	var c1, c2, c3, c4 int
	i, l := 0, len(str)
	var out []byte

	for i < l {
		for i < l && base64DecodeChars[str[i]] == -1 {
			i++
		}
		if i == l {
			break
		}
		c1 = base64DecodeChars[str[i]]
		i++

		for i < l && base64DecodeChars[str[i]] == -1 {
			i++
		}
		if i == l {
			break
		}
		c2 = base64DecodeChars[str[i]]
		i++

		out = append(out, byte((c1<<2)|(c2>>4)))

		for i < l && str[i] == '=' {
			i++
		}
		if i == l {
			break
		}
		c3 = base64DecodeChars[str[i]]
		i++

		out = append(out, byte(((c2&0x0f)<<4)|(c3>>2)))

		for i < l && str[i] == '=' {
			i++
		}
		if i == l {
			break
		}
		c4 = base64DecodeChars[str[i]]
		i++

		out = append(out, byte(((c3&0x03)<<6)|c4))
	}
	return string(out)
}
