# Fantoccini
Marionette adaptation in Go

Okok, funky name I know. I just ran out of name ideas to go with the whole "Marionette" trope, so I just looked through thesaurus.com for synonyms.

Documentation WIP, see [Marionette](https://github.com/MoonriseSunset/Marionette) for more general info :)

## Features:
- Automatic switching between encoding to dollcode and decoding back to unicode/text.

- **Converts:**

    - Decimal, Hexidecimal, Octal, and Binary
    - Strings and single characters  


# Usage Instructions:
All the config needed to tweak Fantoccini's functionality lies between Lines ``4-31``. Fantoccini uses the standard library *only*, so no additional packages need to be installed.

**Instructions:**

1. Make sure *Go* is installed and properly set up.
2. Download ``fantoccini.go`` and place it in whatever directory you want. All translation capabilities are housed within this one file.
3. Convert away by running ``go run fantoccini.go`` in the terminal while in the directory Fantoccini is in.
4. If converting from a file, make sure to put the ``.txt`` file in the **same** directory as the program, or else it will not know what to do. Fantoccini will automatically create an output file if enabled with the name specified in the ``outputName`` variable. **Be careful**, Fantoccini *will overwrite* the output file, regardless of if there are existing contents in the file. 

**Notes:**

- ``simpleMode`` is on by default, which means Fantoccini will only read the input string and output via commandline, *it will not read any files nor write to any files*.


- To **encode** to additional bases (binary, octal, and hex), put ``0b``, ``0o``, or ``0x`` in front of the number. This tells the string-to-number converter Fantoccini uses to use a different base. 
- To **decode** to other bases (same selection as above), set the ``decodeMode`` variable to: ``"b"`` for binary, ``"o"`` for octal, ``"d"`` for decimal, and ``"h"`` for hexidecimal conversion.


- By default, both ``encodeMode`` and ``decodeMode`` are set to ``"s"``, this means that Fantoccini will treat the input as a string when encoding, and will encode *to a string* if decoding.

- If you want to see more info when translating the dollcode, set ``verboseConsole = true``.

# References/Acknowledgement:

The timer function used in Fantoccini originated from [this post on coderwall](https://coderwall.com/p/cp5fya/measuring-execution-time-in-go).

Marionette, Puppeteer, and Fantoccini Dollcode generators are heavily inspired by [noe's dollcode standard and generator](https://noe.sh/dollcode/).

**I am not a doll, nor claim responsibility for any implications with the lore of dolls or the Voidgoddess site**. A close friend of mine *is* a doll and introduced me to dollcode. I found the modified ternary structure interesting, and decided to create the 3 translation suites for practice and as an exploration of the capabilities of these coding languages.

# TO-DO:

- [ ] Simplify file I/O

# Customization Documentation:

**WIP**

# Function Documentation:

## ``condense(input, reversed, spaced)``
- ``input`` (type: ``[]string``): String slice to merge
- ``reversed`` (type: ``bool``): output reversed string
- ``spaced`` (type: ``bool``): insert spaces between slice indicies 
- **RETURNS condensed output as a ``string``**

I made this function without knowing Go had a function in the standard library that could join slices, however said function didn't have the additional functionality this one provides, so I kept it. In essence just iterates through all slices in the ``input``, and appends the contents of each slice index into a string, and spits them out. Setting ``reversed`` to ``true`` causes the function to iterate from *final index to 0* of the slice, effectively reversing it. Setting ``spaced`` to ``true`` inserts spaces between each slice, which is used when separating the dollcodes for different characters.

## ``numToDollcode(input)``
- ``input`` (type: ``string``): numerical value to be encoded
- **RETURNS dollcode as a ``string``**

The reason the input is a ``string`` type is so the function can perform string operations **and** be able to convert multiple bases. ``numToDollcode`` uses the ``strings.ParseInt()`` function to convert the input into numbers, which has the neat side effect of allowing the function to also convert *binary, octal, and hexidecimal* simply by putting ``0b``, ``0o``, or ``0x`` respectively in front of the number being inputted.

## ``stringToDollcode(input)``
- ``input`` (type: ``string``): String to be encoded
- **RETURNS dollcode as a ``string``**

Leverages the ``numToDollcode()`` function along with Go runes to translate **full unicode** into dollcode.

## ``dollcodeToNum(input, mode)``
- ``input`` (type: ``string``): Dollcode to be decoded to number
- ``mode`` (type: ``string``): What base should be used as the output
- **RETURNS ``string``**

Decodes dollcode to a number with base indicated by the ``mode`` parameter. ``mode`` can be set to: ``"d"`` for decimal, ``"h"`` for hexidecimal, ``"o"`` for octal, or ``"b"`` for binary.


## ``dollcodeToString(input)``
- ``input`` (type: ``string``): Dollcode to be decoded to string
- **RETURNS ``string``**

Decodes dollcode to a string, uses ``dollcodeToNum()`` in decimal mode in addition to Go rune and string operations to then turn the unicode points into characters.

## ``encode(input)``
- ``input`` (type: ``string``): Input to be encoded to dollcode
- **RETURNS ``string``**

Encodes the input into dollcode. The input can either be a number or a ``string``. The function will return an error if the ``encodeMode`` config variable does not match what is being inputted, so make sure ``encodeMode`` is set to ``"n"`` for number encoding, or ``"s"`` for string translation.

## ``decode(input)``
- ``input`` (type: ``string``): Dollcode string to be decoded
- **RETURNS ``string``**

Decodes dollcode input into output specified by ``decodeMode`` config variable. The function will return an error if the ``decodeMode`` config variable does not match what is being inputted, so make sure ``decodeMode`` is set to ``"d"`` for decimal, ``"h"`` for hexidecimal, ``"o"`` for octal, ``"b"`` for binary, or ``"s"`` for string output.

## ``translate(input)``
- ``input`` (type: ``string``): String to translate
- **RETURNS ``string``**

Condenses ``encode()`` and ``decode()`` functions into one. Will automatically switch to decoding if it detects dollcode characters in the input, and defaults to encoding otherwise.

## ``threadedTranslate(input, outputChannel)``
- ``input`` (type: ``string``): String to translate
- ``outputChannel`` (type: ``chan string``): channel variable to send data to
- **RETURNS nothing, sends output to channel variable**

For multithreaded translation, ``threadedTranslate()`` runs exactly the same way ``transslate()`` does, with the only difference being the output being sent to the specified ``outputChannel`` instead of being returned.
