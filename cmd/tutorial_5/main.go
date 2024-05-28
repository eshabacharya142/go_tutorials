package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	// when u are dealing with strings in Go, you are dealing with a value whose underlying representation is in Bytes
	// r   é    s   u   m   é
	// 0 [1,2]  3   4   5  [6,7]
	fmt.Printf("The length of myString is %v\n", len(myString)) // length is in bytes and not the number of characters

	// runes are just an alias for int32

	var myString1 = []rune("résumé")
	var indexed1 = myString1[1]
	fmt.Printf("%v, %T\n", indexed1, indexed1)
	for i, v := range myString1 {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of myString is %v", len(myString1)) // the number of characters

	var myRune = 'a'
	fmt.Printf("\n myRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
