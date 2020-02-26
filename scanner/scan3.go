package main

import (
	"fmt"
	"os"
	"text/scanner"
	"unicode"
)

func main() {
	name := "/tmp/a.s"
	f, _ := os.Open(name)
	defer f.Close()

	sc := new(scanner.Scanner)
	sc.Init(f)
	sc.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	sc.Position.Filename = name

	sc.Mode = scanner.ScanChars |
		scanner.ScanFloats |
		scanner.ScanIdents |
		scanner.ScanInts |
		scanner.ScanStrings |
		scanner.ScanComments
	sc.IsIdentRune = func(ch rune, i int) bool {
		if unicode.IsLetter(ch) {
			return true
		}

		switch ch {
		case '_': // Underscore; traditional.
			return true
		case '\u00B7': // Represents the period in runtime.exit. U+00B7 '·' middle dot
			return true
		case '\u2215': // Represents the slash in runtime/debug.setGCPercent. U+2215 '∕' division slash
			return true
		}
		return i > 0 && unicode.IsDigit(ch)
	}
	var tok rune
	tok = sc.Scan()
	fmt.Println("text-intokenizer:", sc.TokenText(), ",pos:", sc.Pos(), ",peek:", sc.Peek())
	fmt.Printf("tok=%c,%d\n\n", tok, tok)

}
