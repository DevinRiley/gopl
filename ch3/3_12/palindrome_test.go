package palindrome

import (
	"gopl.io/ch3/3_12"
	"testing"
)

func TestPalindrome(t *testing.T) {
	var res bool

	res = palindrome.IsPalindrome("ABC")
	if res != false {
		t.Fail()
	}

	res = palindrome.IsPalindrome("ABA")
	if res != true {
		t.Fail()
	}

	res = palindrome.IsPalindrome("ABBA")
	if res != true {
		t.Fail()
	}
}
