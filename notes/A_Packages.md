# Go Projects

In Go, there's a concept of _**modules**_ and _**packages**_

A package is just a folder that contains go files

A module is a collection of packages

## In Practise

To start a new project we need to do the following:

1. Create (or be in) the directory where we want our project to go and make it pwd
2. Initialise a module using the CLI and give it a name

> It is a convention to provide where the project is hosted as the name

To follow convention the cli command for this is:

```console
go mod init github.com/aalnussairi/lets_go
```

This creates a go.mod file, below is it's contents:

```txt
module github.com/aalnussairi/lets_go

go 1.23.0

```

We can see it's just got the module name at the moment and the go version

It seems kinda like a package.json file but using the .mod file syntax, whatever that is

We'll see what happens when we add dependencies later

To create a package, we'll create a folder, I'll just use `src/main.go` and put all of my
'main' package stuff in there.

In the main.go file, I'll identify that this is part of the main package by identifying
the package name at the start of the file

```go
package main
```

Our entrypoint will be a function called main within the main package.

```go
func main() {
    fmt.Println("Hello, World!")
}
```

Let's disect this code block.

- `func`: function declaration, equivalent to function in JS
- `main`: function name, in this case it's main, our entrypoint
- `()`: function arguments
- `{...}`: function block scope
- `fmt`: the name of a package that contains some console logging stuff
- `Println`: the name of the function we're using from this package
- `"Hello, World!"`: the string input argument to the Println function

Key things to note:

1. `fmt` is a package that will need to be imported
2. `Println`:
   - appends a `\n` to the end of the input and prints it
   - starts with an uppercase letter, the convention is expored identifiers are upper case
3. You are guarenateed to not have noticed the fact that go uses tabs for spacing by convention

To import fmt, we need the following line

```go
import "fmt"
```

To combine everything together, here's how the file should look:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> How do we run this though?

Note that go is a compiled language unlike JavaScript.

If you've only experienced programming from an interpreted language then you may be confused. You may even wonder what about JIT and stuff like that.

This isn't the most important thing in the world though, so feel free to skip the next bit (at this stage).

> Go is a strongly typed, statically compiled language.
>
> Focusing on the statically compiled stuff for now, go
> uses a compiler to turn the code it writes into machine
> code (NOT bytecode).
>
> It takes the source and generates a binary that can be run
> by the operating system without the need for a runtime or an
> interpreter.

if we want to run the code we just wrote, we can use the following
command:

```console
go run src/main.go
```

This will compile the source into machine code then run it

If we want to compile the source into a re-runable executable,
we can use the following command

```console
go build src/main.go
```

or

```console
go build -o dist/lesson_A src/main.go
```

we can then run the binary the way we normally run binaries on our OS (depends on the OS)

> You can compile for other OS's and architectures using
> the GOOS and GOARCH env variables
>
> env GOOS=target-os GOARCH=target-architecture go build `<package>`
>
> [Check docs for more info](https://pkg.go.dev/internal/platform)
