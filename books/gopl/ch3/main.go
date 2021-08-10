package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	// overflow
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+2, i*i)
	fmt.Printf("%08b\n", i+2)

	var j uint8 = 255
	var k uint8 = 25
	fmt.Println(j * k)

	// unary addition
	var p int
	fmt.Println(+p, -p)

	// binary operatior
	var a uint8 = 1
	var b uint8 = 1 << 1
	var c uint8 = 1<<1 | 1<<5
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println(int(a))
	fmt.Println(int(b))
	fmt.Println(int(c))

	// set theory with bit operator
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // the set {1, 5}
	fmt.Printf("%08b\n", y) // the set {1, 2}

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	// conversions
	var n = 1e100
	fmt.Printf("%T, %v\n", n, n)
	nint := int(n)
	fmt.Printf("%T, %v\n", nint, nint)

	// write and print integers
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	// hex
	h := int(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", h)
	// rune
	ascii := 'a'
	unicode := 'ㅂ'
	newline := '\n'
	fmt.Printf("%d %[1]c %#[1]q\n", ascii)
	fmt.Printf("%d %[1]c %#[1]q\n", unicode)
	fmt.Printf("%d %#[1]q\n", newline)

	// float
	var f float32 = 16777216
	fmt.Println(f == f+1)
	fmt.Println(1 << 24)

	// sci notation
	const Planck = 6.62606957e23
	fmt.Println(Planck)

	// verb
	for i := 0; i < 8; i++ {
		fmt.Printf("x = %d e^x = %.3f\n", i, math.Exp(float64(i)))
	}

	// special values
	var z float64
	fmt.Println(z, -z, 1/z, -1/z)

	// complex numbers
	var c1 complex128 = complex(1, 2)
	var c2 complex128 = complex(3, 4)
	fmt.Println(c1 * c2)
	fmt.Println(real(c1 * c2))
	fmt.Println(imag(c1 * c2))

	// strings
	str := "español"      // letras 7
	fmt.Println(len(str)) // "8"
	fmt.Println(str[0], str[7])

	// sub-string
	fmt.Println(str[0:2]) // "es"

	// string literals
	lit := "hello, world"
	fmt.Println(lit)

	// scape sequences
	fmt.Println("alert", "\a")

	// raw literal
	raw := `Hello
	world`
	fmt.Println(raw)

	// unicode
	// rune literals // "\uXXXX"
	symbol := '世'
	code := '\u4e16'
	code2 := '\U00004e16'
	fmt.Println(symbol) // "19990"
	fmt.Println(code)   // "19990"
	fmt.Println(code2)  // "19990"

	fmt.Printf("%b\n", code2) // "100111000010110"

	// escape secuences
	// A
	letter_a := '\x41'                   // hexa escape
	fmt.Println(letter_a)                // "65" int32
	fmt.Printf("%b\n", letter_a)         // "1000001" byte
	fmt.Printf("%s\n", string(letter_a)) // "A" string

	// check
	fmt.Println(HasPrefix("maracuya", "mara")) // "true"

	fmt.Println(HasSuffix("maracuya", "mara")) // "false"

	fmt.Println(Contains("hola", "ola")) // "true"

	// runes
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// utf8 decoder
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, r)
		i += size
	}

	// range support utf-8 decoding implicitly
	for i, r := range s {
		fmt.Printf("%d\t%q\t%x\n", i, r, r)
	}

	// unexpected input byte
	unex := '\uFFFD'
	fmt.Printf("%q\n", unex)

	// utf-8 bytes vs. runes
	text := "Dios"
	fmt.Printf("% x\n", text)
	r := []rune(text)
	fmt.Printf("%x\n", r)
	// converting to string
	fmt.Println(string(text))
	fmt.Println(string(r))

	korean := "바다"
	fmt.Printf("% x\n", korean)
	coding := []rune(korean)
	fmt.Printf("korean: %x\n", coding)

	esp := "hola"
	fmt.Printf("% x\n", esp)
	coding_esp := []rune(esp)
	fmt.Printf("español: %x\n", coding_esp)

	// unicode pure
	str_value := "g" // string value
	int_value := 103 // int value
	uint8_value := uint8(103)
	//fmt.Printf("%b\n", bits_value)  // "1100111"
	byte_value := byte(103) // "[103]"
	hex := "\x67"           // hexa escape secuense (utf8 encoding)
	g_unicode := '\u0067'   // unicode

	fmt.Println(
		str_value,
		hex,
		string(g_unicode),
		string(rune(int_value)),
		string(uint8_value),
		string(byte_value),
	)

	// playing with unicode
	hexa := 0x67                    // hexa literal
	fmt.Println(hexa)               // "103"
	fmt.Println(byte(hexa))         // "103"
	fmt.Println(rune(hexa))         // "103"
	fmt.Println(string(rune(hexa))) // "g"
	fmt.Printf("%o %[1]x\n", hexa)  // "147 67"

	// rune
	var rn rune = 103
	fmt.Println(rn)         // "103"
	fmt.Println(string(rn)) // "g"
	var rc rune = 'g'
	fmt.Println(rc) // "103"
}
