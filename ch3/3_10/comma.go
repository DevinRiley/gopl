package comma

import (
	"bytes"
)

func Comma(s string) string {
	if len(s) < 4 {
		return s
	}

	var buf bytes.Buffer
	i := 0

	leading := len(s) % 3

	if leading != 0 {
		buf.Write([]byte(s[:leading]))
		buf.WriteByte(',')
		i += leading
	}

	for _, _ = range s {
		if i == len(s) {
			break
		}

		if i > 2 {
			buf.WriteByte(',')
		}

		buf.Write([]byte(s[i : i+3]))
		i += 3
	}

	return buf.String()
}
