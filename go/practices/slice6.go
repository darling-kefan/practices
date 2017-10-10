package main

import (
	"fmt"
	// "reflect"
	"unicode/utf8"
)

func main() {
	// Declare a string with both chinese and english characters.
	s := "世界 means world"

	fmt.Printf("%q\n", s[0:6])

	// UTFMax is 5 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	for i, r := range s {
		// fmt.Println(i, string(r), reflect.TypeOf(r))

		rl := utf8.RuneLen(r)
		ui := i + rl

		copy(buf[:], s[i:ui])

		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
