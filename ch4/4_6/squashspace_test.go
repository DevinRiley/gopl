package squashSpace

import (
  "testing"
  "fmt"
)

func TestSquashOneSpace(t *testing.T) {
  s := "hello world"
  b := []byte(s)

  b = SquashSpace(b)
  fmt.Printf("%s\n", b)

  for i, char := range s {
    if b[i] != byte(char) {
      t.Fail()
    }
  }
}

func TestSquashTwoSpaces(t *testing.T) {
  s := "hello  world"
  b := []byte(s)
  expected := "hello world"

  b = SquashSpace(b)
  fmt.Printf("%s\n", b)

  for i, char := range expected {
    if b[i] != byte(char) {
      t.Fail()
    }
  }
}

func TestSquashMultipleSpaceRuns(t *testing.T) {
  s := "a    b    c   d"
  b := []byte(s)
  expected := "a b c d"

  b = SquashSpace(b)
  fmt.Printf("%s\n", b)

  for i, char := range expected {
    if b[i] != byte(char) {
      t.Fail()
    }
  }
}
