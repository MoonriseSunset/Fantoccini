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

# Function Documentation:

**WIP**