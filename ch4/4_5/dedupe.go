package dedupe

func Dedupe(s []string) []string {
  i, dupeCount := 0, 0
  j := 1

  for j < len(s) {
    for s[i] == s[j] {
      dupeCount++

      if j == len(s)-1 {
        break
      } else {
        j++
      }
    }

    i += 1
    if dupeCount > 0 {
      s[i] = s[j]
    }
    j += 1
  }
  return s[:len(s)-dupeCount]
}
