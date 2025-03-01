package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

/* Create a function that takes two colors as input and returns the resulting color when they are mixed. The colors will be provided as strings representing primary colors (red, blue, yellow) and their combinations (purple, orange, green, black).

Mixing Rules:

Same Color:

Mixing the same color yields that color:

red + red results in red.

Primary Colors:

Mixing two primary colors results in a secondary color:

red + blue results in purple.
red + yellow results in orange.
blue + yellow results in green.

Primary with Secondary:

Mixing a primary color with a secondary color which is partly composed of the main color gives precedence to the secondary color:

red + orange results in orange.
blue + purple results in purple.
blue + green results in green.
red + purple results in purple.
yellow + green results in green.
yellow + orange results in orange.


Black color:

mix a primary color with a secondary color that is not partly composed of the main colour returns black.
Input
Line 1: The first color.
Line 2: The second color.

Output
The resulting color.
Constraints

Example

Input		Output

red 		purple

blue
 */

func mixColors(color1, color2 string) string {
    // Normalize input to lowercase
    color1 = strings.ToLower(color1)
    color2 = strings.ToLower(color2)

    // Define mixing rules
    if color1 == color2 {
        return color1 // Mixing the same color yields that color
    }

    // Primary color combinations
    if (color1 == "red" && color2 == "blue") || (color1 == "blue" && color2 == "red") {
        return "purple"
    }
    if (color1 == "red" && color2 == "yellow") || (color1 == "yellow" && color2 == "red") {
        return "orange"
    }
    if (color1 == "blue" && color2 == "yellow") || (color1 == "yellow" && color2 == "blue") {
        return "green"
    }

    // Primary with secondary color combinations
    if (color1 == "red" && color2 == "orange") || (color1 == "orange" && color2 == "red") {
        return "orange"
    }
    if (color1 == "blue" && color2 == "purple") || (color1 == "purple" && color2 == "blue") {
        return "purple"
    }
    if (color1 == "blue" && color2 == "green") || (color1 == "green" && color2 == "blue") {
        return "green"
    }
    if (color1 == "red" && color2 == "purple") || (color1 == "purple" && color2 == "red") {
        return "purple"
    }
    if (color1 == "yellow" && color2 == "green") || (color1 == "green" && color2 == "yellow") {
        return "green"
    }
    if (color1 == "yellow" && color2 == "orange") || (color1 == "orange" && color2 == "yellow") {
        return "orange"
    }

    // Black color combinations
    return "black"
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    scanner.Scan()
    color1 := scanner.Text()
    _ = color1 // to avoid unused error
    scanner.Scan()
    color2 := scanner.Text()
    _ = color2 // to avoid unused error

    result := mixColors(color1, color2)
    fmt.Println(result) // Write answer to stdout
}
