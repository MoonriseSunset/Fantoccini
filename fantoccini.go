package main

import (
	"fmt"

	//"unicode"
	"strconv"
	"strings"

	//"os"
)

//CONFIG:
//const fromFile = false

//Encode and Decode modes: s for string, n for decimal/integer/hex/octal/binary
const encodeMode = "s"
const decodeMode = "s"

// ---------------------------

//INPUT:
var input = "Hello, World!"

// ---------------------------

//Misc variables:

var(

	charmap = [3]string{"▌", "▖", "▘"}
	decodemap = [3]string{"▖", "▘", "▌"}

)

// ---------------------------

//Functions

//Helper functions

func condense(input []string, reversed bool, spaced bool) string {

	var output string
	var spacer string

	if(spaced) {
		spacer = " "
	} else{
		spacer = ""
	}

	if(reversed) {
		for i := len(input)-1; i >= 0; i-- {
			output = output + input[i] + spacer
		}

	} else{
		for i := range input {
			output = output + input[i] + spacer
		}
	}

	return output
}
func numToDollcode(input string) string {

	intermediate := []string{}

	
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


	return condense(intermediate, true, false)
}	

func stringToDollcode(input string) string {
	intermediate := []string{}
	units := []rune(input)

	for i := 0; i < len(units); i++ {
		intermediate = append(intermediate, numToDollcode(strconv.Itoa(int(units[i]))))
	}

	return condense(intermediate, false, true)
}

//Primary Functions
func encode(input string) string {

	var output string

	switch encodeMode {
		case "n":
			output = numToDollcode(input)
		case "s":
			output = stringToDollcode(input)
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