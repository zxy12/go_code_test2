package main

import (
	"fmt"
	"os"
	"text/scanner"
	"unicode"
)

func main() {
	f := "/tmp/a.s"
	s := new(scanner.Scanner)
	fd, _ := os.Open(f)
	defer fd.Close()
	s.Init(fd)
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	s.Position.Filename = f
	s.Mode = scanner.ScanChars |
		scanner.ScanFloats |
		scanner.ScanIdents |
		scanner.ScanInts |
		scanner.ScanStrings |
		scanner.ScanComments
	s.IsIdentRune = func(ch rune, i int) bool {
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
		// Digits are OK only after the first character.
		return i > 0 && unicode.IsDigit(ch)
	}

	var tok rune
	tok = s.Scan()
	fmt.Println("text-intokenizer:", s.TokenText(), ",pos:", s.Pos(), ",peek:", s.Peek(), ",tok", tok)

}
