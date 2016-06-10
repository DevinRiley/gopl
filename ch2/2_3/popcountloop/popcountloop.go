// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcountloop

// pc[i] is the population count of i.
var pc [256]byte

// constructs a table of the hamming weight (population count)
// of an 8-bit integer.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
  var sum byte

  for i:= byte(0); i<7; i++ {
    sum += pc[byte(x>>(i*8))]
  }
  return int(sum)
}

//!-
