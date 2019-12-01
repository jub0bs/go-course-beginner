### Installation

* Go (`go version`: 1.13.4)
* Visual Studio Code
* Go VSCode extension
* Git
* Wifi

---

### [A little bit about me](https://twitter.com/jub0bs)

* @jub0bs
* freelance developer
* Go trainer
* security researcher & bug-bounty hunter

---

### Who are you?

---

### What to expect from the course

* an invitation to Go
* emphasis on practice over theory
* interaction
* regular breaks

---

### Project: username checker

* checking a username on social networks
  * validity
  * availability
* business case
  * brand consistency
  * username squatting (?)
* a forcing function for learning Go

---

### Go's history

* created at Google in 2009
* ...by Rob Pike, Robert Griesemer, Ken Thompson
* Go 1.0 released in 2012
* Go 1.13.4 (latest release)
* powers Docker, Kubernetes, and many more

---

### What makes Golang special

* syntax reminiscent of C
* strong typing
* fast compilation
* garbage collector (fast & easy to tune)
* pointers, but no pointer arithmetic

---

### What makes Golang special (cont'd)

* first-class functions
* object-oriented, but no inheritance
* built-in concurrency (channels, goroutines)
* no operator overloading
* no generics

---

### What makes Go's ecosystem special

* BSD licence
* an agenda of simplicity
* no centralised package repository
* rich standard library
* rich command-line interface (gofmt!)
* vibrant community

---

### Why learn Go now

* built-in concurrency
* slowly evolving language
* dependency management (finally) solved
* good serverless support (AWS, GCP, Azure)

---

### Go playground

* https://play.golang.org
* de facto online REPL
* automatic imports
* automatic formatting
* permalinks

---

### Go playground limitations

* a single, main package
* no networking
* no command-line arguments
* only standard-library imports
* concurrency is weird...

---

# First steps in Go

---

### [Hello, World!](https://play.golang.org/p/xY2BPIIxBbk)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

---

### Hello, World!

* all source files start with a package clause
* `main` package: defines an executable program
* `main` function: program's entrypoint
* import-declaration syntax
* no semicolons needed to terminate statements

---

### Building & running a Go program

* `$ go build main.go`
* `$ go run main.go`

---

### Keywords

```txt
break       default       func      interface   select
case        defer         go        map         struct
chan        else          goto      package     switch
const       fallthrough   if        range       type
continue    for           import    return      var
```

---

### Predeclared names

* Constants
```txt
false true nil iota
```
* Types (more about that later)
* Built-in functions
```txt
make    len     cap     new
append  copy    delete
complex real    imag
panic   recover
```

---

### Comments

* end-of-line comments:

```go
var foo = "foo" // this is a comment
```

* multiline comments (do not nest)

```go
/*
  this is
  a multiline comment
*/
```

---

## Imports

---

### Package path

* globally unique
* examples
  * `fmt`
  * `io/ioutil`
  * `github.com/jubobs/username/twitter`

---

### Default package name

* the last component of the package path
* examples
  * `fmt`
  * `ioutil`
  * `twitter`
* picking a different name is possible
```go
import tw "github.com/jubobs/username/twitter"
```

---

### Import block

* possible to factor multiple import declarations

```go
import (
  "encoding/json"
  "fmt"
  "import io"
)
```

---

### Elements of style: import order

```go
import (
  // standard-library packages
  //
  // 3rd-party packages
  //
  // your packages
)
```

---

### The `fmt` package

* `fmt.Println`
* `fmt.Sprintf`
* `fmt.Printf`

---

### Elements of style

* source files in UTF-8
* indentation: tabs, not spaces

---

# Variables

A variable is an _addressable value_.

---

## Variable declaration

```go
var <name> <type> = <expression>
```

* type is optional...
* initialising expression is optional...
* ...but not both!

---

### Variable blocks

possible to factor multiple `var` declarations

```go
var (
	legalRegexp   = "^[A-Za-z0-9]*$"
	illegalPattern = "--"
)
```

---

### Elements of style: variable naming

* strive for conciseness
* camel case (no underscores)
* the smaller the scope, the shorter the name

---

### Short variable declaration

```go
<name> := <expression>
```

* type is inferred from expression
* there must be new values on the left-hand side!
* unused declared variable => compilation error!
* use the _blank identifier_ (`_`) when needed
* shadowing (due to lexical scoping) can lead to bugs

---

### Initial value

* no uninitialised variable in Go!
* implicitly assigned the _zero value_ of its type
```go
var i int
fmt.Printf("value: %d\n", i) // "0"
```

---

### `var` or `:=`

* `:=` only allowed for local variables
* favour `var` over `:=` when
  * the initial value must be the zero value
  * the initial value doesn't matter


---

# Basic types

---

### Booleans

* type `bool`
* values: `false` or `true`
* zero value: `false`
* not: `!`
* and: `&&`
* or: `||`

---

### Signed integers

* `int` (size is platform-dependent)
* `int8`
* `int16`
* `int32` (a.k.a. `rune`)
* `int64`

---

### Unsigned integers

* `uint` (size is platform-dependent)
* `uint8` (a.k.a. `byte`)
* `uint16`
* `uint32`
* `uint64`
* `uintptr` (size is platform-dependent)

---

### Floating-point numbers

* `float32`
* `float64`
* `complex64`
* `complex128`

---

### Numeric operators

* arithmetic: `+`, `-`, `*`, `/`, `%`
* comparison: `==`, `!=`, `<`, `<=`, `>`, `>=`
* bitwise (integers): `&`, `|`, `^`, `&^`

---

### Increment/decrement statements

* only work with integer types
* increment `i` by 1: `i++`
* decrement `i` by 1: `i--`
* no prefix inc/dec operators
* `i++` is a statement, not an expression!

---

### Runes

* `rune` is an alias for `int32`
* represents a single Unicode code point
* UTF-8 encoding (1 to 4 bytes per rune)
* number of bytes in rune `r`: `utf8.RuneLen(r)`

---

### Rune literals

* delimited by single quotes: `'a'`, `'汉'`
* some escape sequences:
  * `'\n'`
  * `'\uhhhh'`
  * `'\Uhhhhhhhh'`

---

### Strings

* immutable data structure
* under the hood
    * a pointer to an array of bytes
    * an int corresponding to the length
* zero value: empty string
* concatenation operator: `+`

---

### String literals

* delimited by double quotes...

```go
s := "Hello World!"
```

* ...or backticks (_raw string_)

```go
`json:"Name"`
```

---

### Useful packages for string manipulation

* `strings`
* `unicode/utf8`
* `strconv`
* `regexp`

---

### String length

```go
s := "Hello, World!"
len(s)                    // number of bytes
utf8.RuneCountInString(s) // number of runes
```

---

### Indexing a string

* 0-based indexing
* `s[i]` is the `i+1`th byte in string `s`
* if `i` out of bounds, _panic_!

---

### Substring operator

* syntax: `s[lo:hi]`
* semi-open interval: bytes from index `lo` to index `hi` (exclusive)
* structural sharing between string and its substrings
* if _hi_ or _lo_ out of bounds or _hi < lo_, panic!

---

### Substring operator (cont'd)

`lo` and `hi` are optional
* `s[:hi]` is equivalent to `s[0:hi]`
* `s[lo:]` is equivalent to `s[lo:len(s)]`
* `s[:]` is equivalent to `s`

---

### Ranging over a string

* iterates over the _runes_ of a string
```
for i, r := range s {
 ...
}
```
* exercise: print each rune and its index

---

### Named types

```go
type <name> <underlying_type>
```

* underlying type can be arbitrary
* user-defined types
* `type` declaration
* _methods_ can be defined on named types!

---

### Conversion

* strong typing
* no implicit casting in Go
* (almost) all conversions must be explicit

---

### Constants

* values fixed at compile time
* `const <name> = <expression>`
* no function calls allowed on right-hand side
* allow optimisations by the compiler
* two categories of constants: typed and untyped


---

### Typed constants

```go
const Monday int = 1
```

---

### Untyped constants

* constants not committed to a particular type
* flexibility!
* six _kinds_ of untyped constants
* numeric untyped constants have at least 256 bits of precision
* [example from the `math` package](https://golang.org/pkg/math/#pkg-constants)

---

### Constant blocks

possible to factor multiple const declarations

```go
const (
  Un = 1
  Deux = 2
  Trois = 3
)
```

---

### `iota` constant generator

* syntactic sugar for creating numeric constants

```go
const (
  Un = iota + 1
  Deux
  Trois
)
```

* RHS expression is implicitly reused within the constant block
* `iota` evaluates to `0`, then `1`, then `2`, etc.

---

### Enum pattern

* no native enumeration types
* instead, named type based on `int`
* not very typesafe
* exercise: write a `Month` enum

---

### Enum pattern: example

```go
type Month int
const (
  _ Month = iota
  January
  February
  //...
  December
)
```

---

## Pointers

* a value that holds the address of a variable
* `*T` denotes the type of a pointer to a variable of type `T`
* zero value: `nil`

---

## Pointers (cont'd)

* `&`: address-of operator
* `*`: dereference operator
* `nil` dereference: panic!

---

# Basic control-flow structures

---

### `if` (basic form)
```go
if <condition> {
  <body>
}
```
* condition is of boolean type
* no parentheses around condition
* braces are mandatory

---

### `if` (with optional declarations)
```go
if [short-var-declarations;] <condition> {
  <body>
}
```
* variables scoped to the `if` statement
* examples in Go playground

---

### `if`/`else`
```go
if <condition> {
  <if-body>
} else {
  <else-body>
}
```
* `else` is seldom used in Go
* avoid long `if`/`else` chains

---

### `switch`

```go
switch <condition> {
case <expression>:
  ...
case <expression>:
  ...
...
default:
  ...
}
```
exercise: write a signum function

---

### `switch`

* cases are evaluated from top to bottom
* the first match is executed
* if none match, the (optional) default case is executed
* `default` can occur at any place among the cases

---

### Tagless `switch`

```
switch {
case <condition>:
  ...
case <condition>:
  ...
...
default:
  ...
}
```
* favour a tagless switch to `if`/`else` chains

---

### `for` loop

```
for <condition> {
  <body>
}
```
* no `while` in Go

---

### `for` infinite loop

```
for {
  <body>
}
```

---

### C-style `for` loop

```
for <init>; <condition>; <post> {
  <body>
}
```

---

### `range`-based `for` loop

```
for <variables> := range <data-structure> {
  <body>
}
```
* `range` works with
  * strings
  * arrays
  * slices
  * maps
  * channels

---

# Functions

---

## Function declaration

```
func <name>(<param-list>) <result-list> {
  <body>
}
```

---

### Multiple results

```
func CountWords(path string) (int, error) {
  // ...
}
```

---

### Named results

```
func CountWords(path string) (c int, err error) {
  // ...
}
```
* useful for
  * clarifying the roles of results
  * simplifying some implementations

---

### Type factorisation

```
func first(a, b int) int {
  return a
}
```

---

### Function call

* no default parameter value
* no optional parameters
* all arguments must be specified

---

### Parameter evaluation

* call-by-value evaluation
* However...
  * Some types are glorified pointers (e.g. strings).
  * These types are known as "reference types".

---

### Functions are first-class values

* a function can be
    * another function's parameter
    * another function's result
* function literals

---

### Recursion

* functions can call themselves
* exercise: mutually recursive functions
    * `isEven(uint) bool`
    * `isOdd(uint) bool`

---

### Closures

* functions can capture variables in their environment
* exercise: stateful function

---

### Project: validating the length of a Twitter username

* `isLongEnough(username string) bool`
* `isShortEnough(username string) bool`

---

# Error handling & reporting

---

### Signaling a failure to the caller

* a function that can fail returns an additional value
* this value comes last, by convention
* only one possible cause: `bool`
* multiple possible causes: `error` (an interface type)
* non-nil error: something bad happened

---

### Error-check idiom

```
if err := fou(); err != nil {
  // unhappy path :(
  // early return
}
// happy path
```
* convention: keep the happy path "on the left"
* (not nested in if-else statements)

---

### Beyond `err != nil`

* you might want to find out more about the nature of the error
* when possible, assert on the error's behaviour
* more on that later (when we know what interfaces are)

---

### Error handling

* always check for errors
* when appropriate, wrap errors in custom errors before reporting them
* few errors should be ignored

---

## Project: validating legal chars/patterns

* `containsOnlyLegalChars(username string) ???`
* `containsNoIllegalPattern(username string) ???`

---

## Project: validate a prospective Twitter username

* `isValid(username string) bool`

---

# Packages

---

# Unit testing

---

### `testing` package

* tests for package `foo` should be written in `foo_test.go`
* `foo_test.go` must be placed inside folder `foo`
* the package declaration of `foo_test.go` matters:
    * `import foo`: white-box testing
    * `import foo_test`: black-box testing
* test methods start with prefix "Test" (e.g. `TestValidity`)
* Run `go test ./...` at the root of your module

---

### Code coverage
```
go test -coverprofile="coverprofile.tmp" ./...
go tool cover --html="coverprofile.tmp"
```
