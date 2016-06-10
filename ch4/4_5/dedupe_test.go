package dedupe

import(
  "gopl/ch4/4_5"
  "testing"
)

func TestNoDupes(t *testing.T) {
  s := []string{"A", "B", "C"}
  res = Dedupe(s)
  if res != s {
    t.Fail()
  }
}
