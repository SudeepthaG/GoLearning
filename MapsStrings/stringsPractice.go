package mapsstrings

import (
	"fmt"
	"unicode/utf8"
)

func StringPractice() string {

	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	fmt.Println()
	printBytes(name)
	fmt.Println()
	printRunes(name)

	name = "Señor"
	fmt.Println()
	fmt.Printf("String: %s\n", name)
	fmt.Println()
	printBytes(name)
	fmt.Println()
	printRunes(name)

	createStringFromByteRune()

	stringOperations()
	return "end of strings"
}

func createStringFromByteRune() {
	// UTF-8 Encoded hex bytes of the string Café.
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println("\n" + str)
	byteSlice = []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println(str)
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\nCharacters with bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

}

func printRunes(s string) {
	fmt.Printf("Characters with runes: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}

	for index, rune := range s {
		fmt.Printf("\n%c starts at byte %d", rune, index)
	}

	fmt.Printf("\nLength: %d\n", utf8.RuneCountInString(s))
	fmt.Printf("Number of bytes: %d\n", len(s))

}

func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

// func mutate(s string)string {
//     s[0] = 'a'//any valid unicode character within single quote is a rune
//     return s
// }

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func stringOperations() {
	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)

	// Concatenation
	string1 = "Go"
	string2 = "is awesome"
	result := string1 + " " + string2
	fmt.Println(result)

	// Mutation- strings are immutable in go. Bypassed using slice of runes
	h := "hello"
	// fmt.Println(mutate(h))		//error
	fmt.Println("passed string after mutation:", mutate([]rune(h)))
	fmt.Println("original string after mutation:", h)

}

// Unrelated special case of variables

// func sum() (int, int) {
// 	fmt.Println("inside")
// 	return 2, 3
// }

// func Summain() {
// 	fmt.Println("Hello, playground")
// 	x := 1
// 	if true {
//
// 		// x,err:=sum() //wont change x value outside if block. so declare err before and assign using
// var err int
//  x,err=sum()
// or option2 is below.bu this wont work for assigning x,err val using sum()
// 		x = 2
// 		err := 3
// 		fmt.Printf("%d %d", x, err)
// 	}
// 	fmt.Printf("\noutside: %d", x)
// }
