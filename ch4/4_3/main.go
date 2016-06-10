// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

func reverse(a *[6]int) {
  i, j := 0, len(a)-1

  for i < j {
    (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
    i += 1
    j -= 1
  }
}
