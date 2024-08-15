package main

import (
	"fmt"			//For printing

	//"unicode"
	"math"			//For Ternary to Decimal conversion

	"strconv" 		//Conversion to and from string type
	"strings" 		//Addtl string functions
	"slices" 		//Slice manipulation

	//"reflect" 	//For returning the type of a variable

	"os"			//For file I/O
)

//CONFIG:
const fromFile = true

//Encode modes: s for string, n for decimal/integer/hex/octal/binary
const encodeMode = "s"

//Decode modes: s for string, d for decimal, h for hex, o for octal, b for binary
const decodeMode = "s"

// ---------------------------

//INPUT:
// if opening a file, place the raw txt in the same directory as the go file, otherwise, just type text into the input variable
var input = "output.txt"

//name of the output file
var outputName = "decoded.txt"

// ---------------------------

//Misc variables:

var(

	charmap = []string{"▌", "▖", "▘"}
	decodemap = []string{"▖", "▘", "▌"}

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
		fmt.Println("Error in string to number conversion! ")
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

func dollcodeToNum(input string, mode string) string {

	ternary := []int{}
	p := 0
	dec := 0.0
	base := 10

	switch (mode){
		case "d":
			base = 10
		case "h":
			base = 16
		case "o":
			base = 8
		case "b":
			base = 2
		default:
			base = 10

		}

	for _, d := range slices.DeleteFunc(strings.Split(input, ""), func(x string) bool { return x == " " }) {
		ternary = append(ternary, slices.Index(decodemap,d)+1)
	}

	for i := len(ternary)-1; i >= 0; i-- {
		dec += math.Pow(3,float64(p))*float64(ternary[i])
		p += 1
	}

	return strconv.FormatInt(int64(dec),base)
}

func dollcodeToString(input string) string {
	var unit []rune
	word := strings.Split(input, " ")

	for _, w := range word{

		U, _ := strconv.ParseInt(dollcodeToNum(w,"d"),10,64)

		unit = append(unit, rune(U))
	}
	return string(unit)
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
		case "d":
			output = dollcodeToNum(input,"d")
		case "h":
			output = dollcodeToNum(input,"h")
		case "o":
			output = dollcodeToNum(input,"o")
		case "b":
			output = dollcodeToNum(input,"b")
		case "s":
			output = dollcodeToString(input)
		default:
			output = "Error: unknown type"

	}

	//fmt.Println(input)

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
	if(fromFile) {
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Println(err)
		} else {
			input = string(data)
		}

		file, err := os.OpenFile(outputName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		file.WriteString(translate(input))
	} else{
		fmt.Println(translate(input))
	}
	
	fmt.Println("Done.")

}