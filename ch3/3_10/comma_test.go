package comma

import (
	"fmt"
	"gopl.io/ch3/3_10"
	"strings"
	"testing"
)

func TestShort(t *testing.T) {
	strs := []string{"1", "12", "123"}
	for _, str := range strs {
		res := comma.Comma(str)
		if strings.Contains(res, ",") {
			fmt.Println(res)
			t.Fail()
		}
	}
}

func TestFourDigit(t *testing.T) {
	res := comma.Comma("1234")
	if res != "1,234" {
		fmt.Println(res)
		t.Fail()
	}
}

func TestFiveDigit(t *testing.T) {
	res := comma.Comma("12345")
	if res != "12,345" {
		fmt.Println(res)
		t.Fail()
	}
}

func TestSixDigit(t *testing.T) {
	res := comma.Comma("123456")
	if res != "123,456" {
		fmt.Println(res)
		t.Fail()
	}
}

func TestLong(t *testing.T) {
	var res string

	res = comma.Comma("123456789")
	if res != "123,456,789" {
		fmt.Println(res)
		t.Fail()
	}
	res = comma.Comma("1234567891")
	if res != "1,234,567,891" {
		fmt.Println(res)
		t.Fail()
	}

	res = comma.Comma("12345678910")
	if res != "12,345,678,910" {
		fmt.Println(res)
		t.Fail()
	}
}
