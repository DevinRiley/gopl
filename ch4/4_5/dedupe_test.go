package dedupe

import(
  "testing"
)

func TestNoDupes(t *testing.T) {
  s := []string{"A", "B", "C"}
  res := Dedupe(s)
  for i, _ := range res {
    if res[i] != s[i] {
      t.Fail()
    }
  }
}

func TestOneDupe(t *testing.T) {
  s := []string{"A", "B", "B", "C"}
  expected := []string{"A", "B", "C"}

  res := Dedupe(s)

  for i, _ := range res {
    if res[i] != expected[i] {
      t.Fail()
    }
  }
}

func TestAdjacentDupes(t *testing.T) {
  s := []string{"A", "B", "B", "B", "C"}
  expected := []string{"A", "B", "C"}

  res := Dedupe(s)

  for i, _ := range res {
    if res[i] != expected[i] {
      t.Fail()
    }
  }
}

func TestDupesAtBeginning(t *testing.T) {
  s := []string{"A", "A", "B", "C"}
  expected := []string{"A", "B", "C"}

  res := Dedupe(s)

  for i, _ := range res {
    if res[i] != expected[i] {
      t.Fail()
    }
  }
}

func TestDupesAtEnd(t *testing.T) {
  s := []string{"A", "B", "C", "C"}
  expected := []string{"A", "B", "C"}

  res := Dedupe(s)

  for i, _ := range res {
    if res[i] != expected[i] {
      t.Fail()
    }
  }
}

func TestMultipleDupes(t *testing.T) {
  s := []string{"A", "A", "B", "B", "C", "C"}
  expected := []string{"A", "B", "C"}

  res := Dedupe(s)

  for i, _ := range res {
    if res[i] != expected[i] {
      t.Fail()
    }
  }
}

func TestAllDupes(t *testing.T) {
  s := []string{"A", "A", "A"}
  expected := []string{"A"}

  res := Dedupe(s)

  for i, _ := range res {
    if res[i] != expected[i] {
      t.Fail()
    }
  }
}

