# Constants, Variables and Data Types

<details>
  <summary>Code Starting point</summary>

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

</details>

## Declaring Variables

To declare variables, we do the following

```go
var number int
```

in the above:

- `var` is the variable declaration (can change), similar to let in JS
- `number` is the variable name
- `int` is the data type. More on types later

You should only use this to declare variables and not constants.

Avoid this for example:

```go
package main

import "fmt"

var pi int

func main() {
  fmt.Println("Hello!")
}
```

Notice how pi is:

1. Not variable, i.e. it never varies, i.e. it never changes
2. Doesn't even get assigned once

In those cases, you really shouldn't be using a variable.

## Assigning Values to Variables

You can assign a value to a variable by:

```go
var pi int = 2
```

Or infer the type of pi by omitting the type.

```go
var pi = 2
```

You can't change the type of a variable, so if you tried to do the following
assignment:

```go
pi = "wrong"
```

after declaring pi is an int or inferring that it's an int, you'll get the
following error

`cannot use "wrong" (untyped string constant) as int value in assignment`

## Inference is great in typescript

But in go you gotta think about memory size.

`int` is actually `int64` or `int32` depending on your cpu architecture.

That means it takes up a full 64 bits of numbers, giving you a minimum value of
`-9,223,372,036,854,775,807` and a max value of `9,223,372,036,854,775,807`

If you're trying to store the values for rgb, that's obviously overkill. More on
that later.

## Some basic types

### Integers

`int` is an integer. That's any whole number like 1 or 12 or even 82.

`int` can come in different sizes:

- `int8` for integers that take up 8 bits
- `int16` for integers that take up 16 bits
- `int32` ... 32 bits
- `int64` ... 64 bits

if you try to assign a value that's too big for the type, you'll get an error.

```go
var number int8 = 128

// cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows) compiler NumericOverflow

```

Hare in mind that the NumericOverflow error only happens at compile time.

So it'll detect the error when you assign a value to the variable, but not when
you're manipulating the variable.

```go
var number int8 = 127
number = number + 1
fmt.Println(number) // -128
```

We get a good ole y2k style bug.

### Unsigned Integers

There's also the unsigned versions of the integers:

- `uint8` for unsigned integers that take up 8 bits
- `uint16` for unsigned integers that take up 16 bits
- `uint32` ... 32 bits
- `uint64` ... 64 bits

unsigned here means there's no minus sign, so the minimum value is 0.

That means you can double the max value of the signed integers, but you can't
store negative numbers.

### Floats

Floats are numbers with decimal points. That's any number like 3.14 or 2.71828
or even 1.61803398875.

They work using the [IEEE 754](https://en.wikipedia.org/wiki/IEEE_754) standard.

It essentially stores the number in scientific notation with a sign bit, an
exponent and a mantissa.

Bunch of nerdy stuff if you ask me.

You can have the following floaters:

- `float32` for floats that take up 32 bits
- `float64` for floats that take up 64 bits

> Remember that bug that people roast javascript for?
>
> The one where `0.1 * 0.2` is actually `0.0200000000000000004`?
>
> Well blame it on the IEEE 754 standard. In the words of Tupac, "That's just
> the way it is".

As always with any IEEE 754 implementations, it's best to round your numbers if
you wanna use them. Also never use floats for money calculations. Use integers
instead. Companies don't like it when you lose them money.

### Arithmetic Operators

You can do calculations with the basic arithmetic operators:

- `+` for addition
- `-` for subtraction
- `*` for multiplication
- `/` for division

```go
var sum = 1 + 2 // 3
var difference = 3 - 2 // 1
var product = 2 * 3 // 6
var quotient = 6 / 2 // 3
```

What's interesting is that you can't mix types in go.

```go
var sum = 1 + 2.0 // mismatched types int and float64
```

You'll get a mismatched types error.

Also, you can't do division with integers and expect a float.

```go
var quotient = 2 / 3 // 0
```

You'll get 0 because it's integer division. You can fix this by casting one of
the numbers to a float.

```go
var quotient = float64(2) / 3 // 0.6666666666666666
```

or, you can get the remainder of the division by using the `%` operator.

```go
var remainder = 2 % 3 // 2
```

That `float64(2)` is called a type conversion or a type cast in other languages.

### Strings

Strings are, well, a bunch of characters. In go, you can use double quotes or
backticks to create strings.

Their type is `string`.

```go
var name = "John Doe"
var message = `Hello, World!`
```

The difference between the two is that backticks are used for raw strings.

That means you can have newlines and tabs in your strings.

```go
var message = `Hello,
World!`
```

we can also concatenate strings using the `+` operator.

```go
var message = "Hello," + " " + "World!"
```

There's no `'` in go for strings, no more deciding between single or double
quotes.

One interesting thing about strings in go is that when getting the length of a
string, you get the number of bytes in the string.

```go
var message = "Hello, World!" // 13 characters
fmt.Println(len(message)) // 13

var message = "السلام عليكم" // 12 characters, means peace be upon you in arabic
fmt.Println(len(message)) // 24
```

The arabic message has 12 characters but 24 bytes because each character is
encoded in 2 bytes.

To get the number of characters in a string, you can use the `utf8` package.

```go
import "fmt"
import "unicode/utf8"

func main() {
  var message = "السلام عليكم"
  fmt.Println(utf8.RuneCountInString(message)) // 12
}
```

Looking at the function we're calling (`utf8.RuneCountInString`), one thing that
you'll notice is that go uses the term `rune`.

What's a rune?

### Runes

A rune is a character in go. It's an alias for `int32`.

If we go back to ASCII, a character is just a number. The letter `A` is 65, the
letter `B` is 66 and so on.

An ASCII character is represented by a byte, which is 8 bits. That means you can
only have 128 characters in ASCII or 256 characters if you include the extended
ASCII.

UTF-8 is a variable length encoding. That means that a character can be 1, 2, 3
or 4 bytes long. That's why the arabic message we had earlier had 24 bytes.

Utf-8 is a superset of ASCII, so the first 128 characters are the same. That
means that the letter `A` is still 65. But the arabic letter `س` is 1587.

In go, a rune is a character. It's an alias for `int32` because it can be any
unicode character.

We can declare a rune by using single quotes.

```go
var letter rune = 'A'
```

If you try to use double quotes, you'll get an error.

```go
var letter rune = "A" // cannot use "A" (type string) as type rune in variable declaration
```

You can also get the unicode code point of a character by casting it to a rune.

```go
var codePoint = rune('A')
fmt.Println(codePoint) // 65
```

Let's leave it there for now. We'll come back to strings and runes inevitably.

### Booleans

Booleans are super simple. They're either `true` or `false`.

Their type is `bool`.

```go
var isTrue = true
var isFalse = false
```

You can use the following operators with booleans:

- `&&` for and
- `||` for or
- `!` for not

```go
var isTrue = true
var isFalse = false

var and = isTrue && isFalse // false
var or = isTrue || isFalse // true
var not = !isTrue // false
```

Bools are the goats.

## Default Values

If you declare a variable without assigning a value to it, it'll get the default
value for the type.

Here's a table of the default values for the basic types:

| Type   | Default Value |
| ------ | ------------- |
| int    | 0             |
| uint   | 0             |
| float  | 0.0           |
| rune   | 0             |
| string | ""            |
| bool   | false         |

## Short Variable Declarations

You can also declare variables using the short variable declaration.

```go
name := "John Doe"
```

## Fancy Shpancy Multiple Declarations in One Line

You can declare multiple variables in one line.

```go
var name, age = "John Doe", 30
```

## When to use `var` and when to use `:=`

Use `var` when you're declaring a variable outside a function.

Use `:=` when you're declaring a variable inside a function.

i.e.

```go
package main

import "fmt"

var name = "John Doe"

func main() {
  age := 30
  fmt.Println(name, age)
}
```

Don't use the short hand declaration for the return value of a function.

It's better to be explicit when it's not obvious what the return value is.

```go
package main

import "fmt"

func main() {
  var name string = foo()
  fmt.Println(name)
}
```

## Constants

Constants are a lot like variables, but they can't change.

You declare them using the `const` keyword.

```go
const pi = 3.14159
```

You can also declare multiple constants in one line.

```go
const (
  pi = 3.14159
  e = 2.71828
)
```

you can't just declare a constant without assigning a value to it.

```go
const pi float64 // missing value in const declaration
```

As always, constants are preferred over variables when they're not going to
change.

## Conclusion

That's it for variables and constants in go.

We've covered:

- Declaring variables
- Assigning values to variables
- Basic types
- Arithmetic operators
- Strings
- Runes
- Booleans
- Default values
- Short variable declarations
- Multiple declarations in one line
- When to use `var` and when to use `:=`
- Constants

I find go interesting because it's super simple to grasp. There's a GC so you
don't have to worry about memory management. It's statically typed so you get
type safety. It's compiled so it's fast. It's got a lot of built in stuff so you
don't have to worry about dependencies.

I look forward to learning more about go and seeing how it contrasts with what I
usually use, which is TypeScript. I'd love to see how I can build a web server
with go and compare it to building a web server with Node.js.

Next up, we'll look at control structures in go.
