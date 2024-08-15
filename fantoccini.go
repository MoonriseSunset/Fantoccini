package main

import (
	"fmt"
	//"os"
)

//CONFIG:
const fromFile = false

//Encode and Decode modes: s for string, d for decimal/integer, h for hexidecimal
const encodeMode = "s"
const decodeMode = "s"

// ---------------------------


func encode(input string) string {

	var output string

	switch encodeMode {
		case "s":
			output = "string"
		case "d":
			output = "decimal"
		case "h":
			output = "hexidecimal"
		default:
			output = "unknown"

	}

	return output
}

func decode(input string) string {

	var output string

	switch decodeMode {
		case "s":
			output = "string"
		case "d":
			output = "decimal"
		case "h":
			output = "hexidecimal"
		default:
			output = "unknown"

	}

	return output
}

func main() {

	fmt.Println("Starting...")

	

}