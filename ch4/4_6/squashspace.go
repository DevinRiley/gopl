package squashSpace

import (
  "unicode"
)

func SquashSpace(b []byte) []byte {
  i, spaceCount := 0, 0
  j := 1

  for j < len(b) {
    for unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[j])) {
      spaceCount++

      if j == len(b)-1 {
        break
      } else {
        j++
      }
    }

    i += 1

    if spaceCount > 0 {
      b[i] = b[j]
    }

    j += 1
  }
  return b[:len(b)-spaceCount]
}
