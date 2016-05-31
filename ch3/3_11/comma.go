package comma

import (
	"bytes"
	"strings"
)

func getSign(str string) (string, string) {
	if strings.HasPrefix(str, "+") {
		return "+", str[1:]
	}

	if strings.HasPrefix(str, "-") {
		return "-", str[1:]
	}

	return "", str
}

func getDecimal(str string) (string, string) {
	strs := strings.Split(str, ".")

	if len(strs) < 2 {
		return str, ""
	} else {
		return strs[0], strs[1]
	}
}

func stripSignAndDecimal(orig string, buf *bytes.Buffer) (string, string) {
	sign, s := getSign(orig)
	buf.Write([]byte(sign))
	return getDecimal(s)
}

func addDecimal(buf *bytes.Buffer, dec string) {
	if len(dec) > 0 {
		buf.WriteByte('.')
		buf.Write([]byte(dec))
	}
}

func addCommas(s string, buf *bytes.Buffer) {
	i := 0
	b := []byte(s)
	leading := len(s) % 3

	if leading > 0 {
		buf.Write(b[:leading])
		i += leading
	}

	for {
		if i == len(s) {
			break
		}

		if i > 0 {
			buf.WriteByte(',')
		}

		buf.Write(b[i : i+3])
		i += 3
	}

}

func Comma(orig string) string {
	var buf bytes.Buffer
	s, dec := stripSignAndDecimal(orig, &buf)

	if len(s) < 4 {
		return orig
	}

	addCommas(s, &buf)
	addDecimal(&buf, dec)
	return buf.String()
}
