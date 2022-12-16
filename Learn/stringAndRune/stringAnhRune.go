package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "日本語"

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}
