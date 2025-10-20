package main

import "fmt"

func main() {
	var smallPositionValue uint8
	smallPositionValue = 10
	fmt.Println(smallPositionValue)

	var smallNegaPositionValue int8
	smallNegaPositionValue = -10
	fmt.Println(smallNegaPositionValue)

	var myInt = 23000
	fmt.Println(myInt)

	myInt = int(smallNegaPositionValue)
	myInt = int(smallPositionValue)

	var myByte byte = 'A'
	fmt.Println(myByte)

	var myRune rune = '🤡'
	fmt.Println(myRune)

}
