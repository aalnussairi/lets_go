package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // Declaring Variables
    var number int

    // Assigning Values to Variables
    var number int = 2
    var number = 2

    // Integer Overflow Example
    var numberOverflow int8 = 127
    numberOverflow = numberOverflow + 1
    fmt.Println(numberOverflow) // -128

    // Arithmetic Operators
    var sum = 1 + 2 // 3
    var difference = 3 - 2 // 1
    var product = 2 * 3 // 6
    var quotient = 6 / 2 // 3

    // Type Mismatch Example
    // var sumMismatch = 1 + 2.0 // mismatched types int and float64

    // Integer Division Example
    var quotientInt = 2 / 3 // 0
    var quotientFloat = float64(2) / 3 // 0.6666666666666666
    var remainder = 2 % 3 // 2

    // Strings
    var name = "John Doe"
    var message = `Hello, World!`
    var rawMessage = `Hello,
World!`
    var concatenatedMessage = "Hello," + " " + "World!"

    // String Length
    var messageLength = "Hello, World!" // 13 characters
    fmt.Println(len(messageLength)) // 13

    var arabicMessage = "السلام عليكم" // 12 characters, means peace be upon you in Arabic
    fmt.Println(len(arabicMessage)) // 24

    // UTF-8 Rune Count
    fmt.Println(utf8.RuneCountInString(arabicMessage)) // 12

    // Runes
    var letter rune = 'A'
    var codePoint = rune('A')
    fmt.Println(codePoint) // 65

    // Booleans
    var isTrue = true
    var isFalse = false
    var and = isTrue && isFalse // false
    var or = isTrue || isFalse // true
    var not = !isTrue // false

    // Short Variable Declarations
    nameShort := "John Doe"

    // Multiple Declarations in One Line
    var nameMultiple, age = "John Doe", 30

    // Constants
    const piConst = 3.14159
    const (
        piMulti = 3.14159
        e       = 2.71828
    )

    // Example of using var outside a function and := inside a function
    var nameOutside = "John Doe"
    func() {
        ageInside := 30
        fmt.Println(nameOutside, ageInside)
    }()

    // Example of explicit type declaration for function return value
    var nameExplicit string = foo()
    fmt.Println(nameExplicit)
}

func foo() string {
    return "Hello from foo"
}