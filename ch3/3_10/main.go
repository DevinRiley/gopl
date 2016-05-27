package main

import(
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

  i := 0
  leadingDigits := len(s) % 3
  if leadingDigits != 0 {
    buf.Write([]byte(s[:leadingDigits]))
    buf.WriteByte(',')
    i += leadingDigits
  }

  for range s {
    remaining := len(s[i:])

    if remaining < 3 {
      buf.Write([]byte(s[i:]))
      break
    }

    if i > 2 { 
      buf.WriteByte(',')
    }
    buf.Write([]byte(s[i:i+3]))
    i += 3
  }

  return buf.String()
}

func main() {
  nums := []string {"123", "1234", "12345", "123456", "1234567"}
  for _, str := range nums {
    res := comma(str)
    fmt.Printf("with commas: %s\n", res)
  }
}
