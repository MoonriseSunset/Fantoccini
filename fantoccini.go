/*
Fantoccini: A Dollcode Translation Suite in Go
Written By MoonriseSunset

Documentation here: https://github.com/MoonriseSunset/Fantoccini/blob/main/README.md

Last edited: 8/20/2024
*/
package main
import ("fmt"; "math"; "time"; "strconv"; "slices"; "strings"; "os")			//For Ternary to Decimal conversion

//INPUT:
// if opening a file, place the raw txt in the same directory as the go file 
// otherwise, just type text into the input variable
var input = ""

//name of the output file, if applicable
var outputName = ""

//CONFIG:

//To have simple input and output, set simpleMode to true, which disables using files entirely
const simpleMode = true

var fromFile = true
var toFile = true

//Verbose Console Output
const verboseConsole = false

//Encode modes: s for string, n for decimal/integer/hex/octal/binary
const encodeMode = "s"

//Decode modes: s for string, d for decimal, h for hex, o for octal, b for binary
const decodeMode = "s"

//NOTE: If you are NOT translating from file, Fantoccini will DISABLE multithreading on the input
//2ND NOTE: The decode functionality is SINGLE THREADED regardless, due to how fast it can decode.
//Number of processes to run
const processes = 1



// END OF CONFIG, DO NOT EDIT BELOW THIS POINT
//Primary Variables:
var(

	charmap = []string{"▌", "▖", "▘"}
	decodemap = []string{"▖", "▘", "▌"}

)
//Helper functions

func Timer(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s", name, elapsed)
}

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

func threadedTranslate(input string, outputChannel chan string) {

	var output string

	if (strings.Contains(input,decodemap[0]) || strings.Contains(input,decodemap[1]) || strings.Contains(input,decodemap[2])) {
		output = decode(input)
	} else{
		output = encode(input)
	}
	outputChannel <- output
	if(verboseConsole) {
		fmt.Println("process finished!")
	}
}

func main() {

	//Config Parsing:

	if(simpleMode){
		fromFile = false
		toFile = false
	}

	// ------------------------

	var(
		frame string
		outputs []string
		finalOutput string
	)

	processNum := 0

	channels := make([]chan string, processes)

	fmt.Println("Starting...")

	defer Timer(time.Now(), "Total file and translation time")

	if(fromFile) {
		//Read from file
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Println(err)
		} else {
			input = string(data)
			fmt.Println("Read from file successfully")

			if(processes > 1 && !strings.Contains(input,decodemap[0]) && !strings.Contains(input,decodemap[1]) && !strings.Contains(input,decodemap[2])) {
				fmt.Println("Splitting input in " + strconv.Itoa(processes) + " pieces")

				//We convert the input into a slice of runes for reliability with special characters
				loaf := []rune(input)
				for i, r := range loaf {
	
					//Convert rune of input back to string
					frame += string(r)
					if(i>0 && (i+1)%(len(loaf)/processes)==0) {
	
						if(verboseConsole) {
							fmt.Println("Starting process " + strconv.Itoa(processNum+1))
							fmt.Println("Process is operating on " + strconv.Itoa(len(frame)) + " characters")
						}
						
						//Creating GoRoutine and Starting it
						channels[processNum] = make(chan string)
						go threadedTranslate(frame, channels[processNum])
						outputs = append(outputs, <-channels[processNum])
	
						//Clear frame and increment index
						frame = ""
						processNum += 1
					}
				}

				finalOutput = condense(outputs, false, false)

			} else{

				fmt.Println("Performing single thread translation")
				finalOutput = translate(input)
			}	
		}
	} else{
		fmt.Println("Performing single thread translation")
		finalOutput = translate(input)
	}

	if(toFile){
		file, err := os.OpenFile(outputName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}
		//Close the file once we're done
		defer file.Close()
		fmt.Println("Created File, writing to it")
		file.WriteString(finalOutput)

	} else{
		fmt.Println(finalOutput)
	}
}