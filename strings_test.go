package main

import (
	"fmt"
	"html"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {

	fmt.Println("\n================== String functions ==================")

	// Trims the characters in "x " from start and end of string
	fmt.Println("s:", strings.Trim("x abc de x", "x "))

	fmt.Println("\n================== String literals (escape characters) ==================")

	fmt.Println("")                      // Default zero value for type string
	fmt.Println("Japan 日本")              // Japan 日本	Go code is Unicode text encoded in UTF‑8
	fmt.Println("\xe6\x97\xa5")          // 日	\xNN specifies a byte
	fmt.Println("\u65E5")                //	日	\uNNNN specifies a Unicode value
	fmt.Println("\\")                    // \	Backslash
	fmt.Println("\"")                    // "	Double quote
	fmt.Println("\t")                    // Tab
	fmt.Println(`\xe6`)                  // \xe6	Raw string literal*
	fmt.Println(html.EscapeString("<>")) // &lt;&gt;	HTML escape for <, >, &, ' and "
	fmt.Println(url.PathEscape("A B"))   // A%20B	URL percent-encoding net/url

	fmt.Println("\n================== Concatenate ==================")

	// Clean and simple string building
	s := fmt.Sprintf("Size: %d MB.", 85)
	fmt.Println(s) // s == "Size: 85 MB."

	// High-performance string concatenation using String Builder
	var b strings.Builder
	b.Grow(32)
	for i, p := range []int{2, 3, 5, 7, 11, 13} {
		fmt.Fprintf(&b, "%d:%d, ", i+1, p) // or b.WriteString(fmt.Sprintf("%d:%d", i+1, p))
	}
	b.WriteRune(' ')
	s = b.String()    // no copying
	s = s[:b.Len()-2] // no copying (removes trailing ", ")
	fmt.Println(s)    // 1:2, 2:3, 3:5, 4:7, 5:11, 6:13

	fmt.Println("\n================== Equal and compare (ignore case) ==================")

	fmt.Println("Japan" == "Japan")                  //	true	Equality
	fmt.Println(strings.EqualFold("Japan", "JAPAN")) // true	Unicode case folding
	fmt.Println("Japan" < "japan")                   // true	Lexicographic order

	fmt.Println("\n================== Length in bytes or runes ==================")

	fmt.Println(len("日"))                    // 3	Length in bytes
	fmt.Println(utf8.RuneCountInString("日")) //	1	in runes unicode/utf8
	fmt.Println(utf8.ValidString("日"))       // true	UTF-8? unicode/utf8

	fmt.Println("\n================== Index, substring, iterate ==================")

	fmt.Println("Japan"[2])   //	'p'	Byte at position 2
	fmt.Println("Japan"[1:3]) //	ap	Byte indexing
	fmt.Println("Japan"[:2])  //	Ja
	fmt.Println("Japan"[2:])  //	pan

	// A Go range loop iterates over UTF-8 encoded characters (runes):
	for i, ch := range "Japan 日本" {
		fmt.Printf("%d:%q ", i, ch)
	}
	// Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'

	// Iterating over bytes produces nonsense characters for non-ASCII text:
	s = "Japan 日本"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%q ", s[i])
	}
	// Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'

	fmt.Println("\n\n================== Search (contains, prefix/suffix, index) ==================")

	fmt.Println(strings.Contains("Japan", "abc"))     //	false	Is abc in Japan?
	fmt.Println(strings.ContainsAny("Japan", "abc"))  //	true	Is a, b or c in Japan?
	fmt.Println(strings.Count("Banana", "ana"))       //	1	Non-overlapping instances of ana
	fmt.Println(strings.HasPrefix("Japan", "Ja"))     //	true	Does Japan start with Ja?
	fmt.Println(strings.HasSuffix("Japan", "pan"))    //	true	Does Japan end with pan?
	fmt.Println(strings.Index("Japan", "abc"))        //	-1	Index of first abc
	fmt.Println(strings.IndexAny("Japan", "abc"))     //	1	a, b or c
	fmt.Println(strings.LastIndex("Japan", "abc"))    //	-1	Index of last abc
	fmt.Println(strings.LastIndexAny("Japan", "abc")) //	3	a, b or c

	fmt.Println("\n================== Replace (uppercase/lowercase, trim) ==================")

	fmt.Println(strings.Replace("foo", "o", ".", 2)) // f..	Replace first two “o” with “.” Use -1 to replace all
	f := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f, "ab"))            //bc	Apply function to each character
	fmt.Println(strings.ToUpper("Japan"))        //JAPAN	Uppercase
	fmt.Println(strings.ToLower("Japan"))        //japan	Lowercase
	fmt.Println(strings.TrimSpace(" foo\n"))     //foo	Strip leading and trailing white space
	fmt.Println(strings.Trim("foo", "fo"))       //	Strip leading and trailing f:s and o:s
	fmt.Println(strings.TrimLeft("foo", "f"))    //oo	only leading
	fmt.Println(strings.TrimRight("foo", "o"))   //f	only trailing
	fmt.Println(strings.TrimPrefix("foo", "fo")) //o
	fmt.Println(strings.TrimSuffix("foo", "o"))  //fo

	fmt.Println("\n================== Split by space or comma ==================")
	fmt.Println(strings.Fields(" a\t b\n"))     // ["a" "b"]	Remove white space
	fmt.Println(strings.Split("a,b", ","))      // ["a" "b"]	Remove separator
	fmt.Println(strings.SplitAfter("a,b", ",")) // ["a," "b"]	Keep separator

	fmt.Println("\n================== Join strings with separator ==================")
	fmt.Println(strings.Join([]string{"a", "b"}, ":")) // a:b	Add separator
	fmt.Println(strings.Repeat("da", 2))               // dada	2 copies of “da”

	fmt.Println("\n================== Format and convert ==================")
	fmt.Println(strconv.Itoa(-42))          //	"-42"	Int to string
	fmt.Println(strconv.FormatInt(255, 16)) // "ff"	Base 16
}
