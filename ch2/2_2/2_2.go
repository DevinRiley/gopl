package main

import (
  "gopl/ch2/2_2/unitconv"
  "strconv"
  "os"
  "fmt"
)

func main() {
  for _, temp := range os.Args[1:] {
    t, err := strconv.ParseFloat(temp, 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(1)
    }

    k := unitconv.Kelvin(t)
    f := unitconv.KToF(k)
    fmt.Printf("%s in Fahrenheit: %s\n", k, f)
  }
}
