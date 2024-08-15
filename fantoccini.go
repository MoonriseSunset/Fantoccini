package main

import (
	"fmt"

	//"unicode"
	"strings"
	"strconv"

	//"os"
	
)

//CONFIG:
//const fromFile = false

//Encode and Decode modes: s for string, n for decimal/integer/hex/octal/binary
const encodeMode = "n"
const decodeMode = "s"

// ---------------------------

//INPUT:
var input = "0b101010"

// ---------------------------

//Misc variables:

var(

	charmap = [3]string{"▌", "▖", "▘"}
	decodemap = [3]string{"▖", "▘", "▌"}

)

//Functions

func numToDollcode(input string) string {

	var output string

	intermediate := make([]string, 0)
	//multiplier, err := strconv.ParseFloat(input,32);
	multiplier, err := strconv.ParseInt(input,0,32)
	
	if (err != nil) {
		fmt.Println("Error in string to float conversion! ")
		fmt.Println(err)
		return "error, check console"
	}

	for (multiplier >= 1) {

		var mod = int64(multiplier) % 3

		if(mod == 0) {
			multiplier = ((multiplier -3)/3)
		} else{
			multiplier = ((multiplier-mod)/3)
		}
		intermediate = append(intermediate, charmap[mod])
	}

	for i := len(intermediate)-1; i >= 0; i-- {
		output = output + intermediate[i]
	}

	return output
}	

func encode(input string) string {

	var output string

	switch encodeMode {
		case "n":
			output = numToDollcode(input)
		case "s":
			output = "string"
		default:
			output = "Error: unknown type"

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
			output = "Error: unknown type"

	}

	fmt.Println(input)

	return output
}

func translate(input string) string {

	var output string

	if (strings.Contains(input,decodemap[0]) || strings.Contains(input,decodemap[1]) || strings.Contains(input,decodemap[2])) {
		output = decode(input)
	} else{
		output = encode(input)
	}

	return output
}

func main() {

	fmt.Println("Starting...")
	fmt.Println(translate(input))
	fmt.Println("Done.")

}