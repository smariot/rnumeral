package rnumeral

import (
	"fmt"
	"testing"
)

func ExampleEncode() {
	fmt.Println(Encode(2006))
	// Output: MMVI
}

func ExampleDecode() {
	fmt.Println(Decode("MMVI"))
	// Output: 2006
}

func TestDecode(t *testing.T) {
	for value := Min; value <= Max; value++ {
		t.Run(fmt.Sprintf("i=%d", value), func(t *testing.T) {
			encode := Encode(value)
			decode := Decode(encode)
			if value != decode {
				t.Errorf("%d != Decode(Encode(%d) [%q]) [%d]", value, value, encode, decode)
			}
		})
	}

	t.Run("invalid decode", func(t *testing.T) {
		const numeral = "MIXI"
		decode := Decode(numeral)
		if decode != 0 {
			t.Errorf("0 != Decode(%q) [%d]", numeral, decode)
		}
	})
}

func TestEncode(t *testing.T) {
	t.Run("invalid encode", func(t *testing.T) {
		const value = 4000
		encode := Encode(value)
		if encode != "" {
			t.Errorf("\"\" != Encode(%d) [%q]", value, encode)
		}
	})
}
