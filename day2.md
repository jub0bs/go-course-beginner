# Composite types

---

## Arrays

---

### Arrays

* fixed-length, homogenous container
* elements are contiguous in memory
* array length is encoded in the type (e.g. `[3]int`)
* zero value: array full of element type's zero value

---

### Array literals
```go
words := [3]string{"foo", "bar", "baz"}
moreWords := [...]string { // size is optional
  "qux",
  "quux", // the last comma is mandatory
}
```

---

### Operations on array

* indexing  (same as for strings)
* array elements are variables: `&arr[i]` is legal
* get a slice from an array: `words[:]`
* ranging over an array
```go
for i, v := range arr {
  // ...
}
```

---

### Array are no reference types

* fully copied when passed as params
* arrays are seldom used directly

---

## Slices

---

### Slices

* fundamental, ubiquitous type in Go
* a resizable view inside an array
* a slice of `T` has type `[]T`

---

### Slices: under the hood

* a slice is composed of
  * a pointer to the underlying array
  * a length (`len`)
  * a capacity (`cap`): size of underlying array

---

### Slice is a reference type

* zero value: `nil`
* a function with a slice parameter gets a copy of it
* but the underlying array is not copied!

---

### Allocating a slice

```go
s := make([]string, l)    // l: length
s := make([]string, l, c) // c: capacity
```

---

### Slice literal

```
words := []string{"foo", "bar", "baz"}
moreWords := []string {
  "qux",
  "quux", // the last comma is mandatory
}
```

---

### Slices are not comparable

* equality comparison disallowed by compiler
* a slice can only be compared to `nil`

---

### Indexing into a slice

* same syntax as for strings
* indexing into a `nil` slice: panic!
* slice elements are variables: `&s[i]` is legal

---

### Slicing operator

* `s1 := s[i:j]`
* if `j` exceeds `cap(s)`, panic!
* yields a different view into the same array
* also: `s1 := s[i:j:k]` gives you more control over the resulting slice's capacity

---

### Ranging over a slice

```
for i, v := range s {
  // ...
}
```

---

### Appending an element to a slice

* use the built-in function `append`
* `s = append(s, elem)`
* built-in `append` copies elements to a larger underlying array if capacity has been reached

---

### Appending another slice to a slice

* `s = append(s, s1...)`

---

### Aliasing

* multiple slices may use the same underlying array
* careful...

---

## Maps

---

### Maps

* collection of key-value pairs
* a map of `K`s to `V`s has type `map[K]V`
* the key type must be comparable
* implemented as a hash table

---

### Maps are reference types

* zero value: `nil`
* accessing elements of a nil map: zero value of value type
* attempting to add an entry to a nil map => panic!

---

### Maps are reference types (cont'd)

* a function that takes a map argument gets a copy of the reference
* the underlying hash table does not get copied

---

### Allocating a map

```
m := make(map[string]int)
```

---

### Map literals

```
enToFr := map[string]string {
  "one": "un",
  "two": "deux",
  "three": "trois", // mandatory comma
}
```

---

### Accessing the elements of a map

* `v := m[k]`
* key `k` needs not be in the map
* ... in which case `v` will be the zero value for the value type!

---

### Testing for membership

```go
v := m[k] // then test that v is not the zero value?
```

---

### Testing for membership

```go
v := m[k] // ambiguous!

// solution
v, ok := m[k]
if ok {
  // the key is present
}
```

---

### Ranging over a map

```
for k, v := range m {
  // do something with k and v
}
```
* iteration order is nondeterministic!

---

### Ranging over a map in a deterministic order

```
// create a slice of keys
keys := make([]string, 0, len(m))
for k := range m {
  keys = append(keys, k)
}

// sort it
sort.Strings(keys)

// then range over the sorted slice of keys
for k := range keys {
  // do something with k and m[k]
}
```

---

### Sets

* No native set type in Go!
* Usually implement as `map[T]bool`
* ... or `map[T]struct{}`

---

### Project: availability check

```
func Available(username string) ???
```

---

### Mechanics of the `defer` keyword

* delays a function call to when the enclosing function terminates
* ... whether normally or not (panic)
* the arguments of the deferred function call are evaluated at the `defer` statement!
* `defer` statements are executed in reverse order in which they appear in the enclosing function

---

### Cleaning up resources with `defer`

```
func Available(username string) (bool, error) {
  resp, err := http.Get("https://twitter.com/" + username)
  if err != nil {
    return false, errors.New("Unknown availability")
  }
  if resp.StatusCode != http.StatusNotFound {
    resp.Body.Close() // <--- reclaim resource no longer needed
    return false, nil
  }
  resp.Body.Close() // <--- needed here too!
  return true, nil
}
```

---

### Cleaning up resources with `defer`

```
func Available(username string) (bool, error) {
  resp, err := http.Get("https://twitter.com/" + username)
  if err != nil {
    return false, errors.New("Unknown availability")
  }
  defer resp.Body.Close() // <--- once and for all!
  if resp.StatusCode != http.StatusNotFound {
    return false, nil
  }
  return true, nil
}
```

---

### Gotcha: `defer` in a loop

* Can you spot the problem?

```
func foo() {
  for i := range ... {
    r, err := obtainResource(i)
    if err != nil {
      // deal with error
    }
    defer r.Close()
    // do something with r...
  }
}
```

---

### Gotcha: `defer` in a loop

```go
func foo() {
  for i := range is {
    if err := process(i); err != nil {
      // deal with error
    }
  }
}
func process(i int) error {
  r, err := obtainResource(i)
  if err != nil {
    // deal with error
  }
  defer r.Close()
  // do something with r...
}
```

---

### Gotcha: errors in deferred function calls

* Deferred statements return no result.
* What if the function call returns an error?

```
func foo() error {
  // ...
  defer badIfThisFails() // but error is lost!
  // ...
}
```

---

### Gotcha: errors in deferred function calls

* workaround: use a named return and a func literal

```go
func foo() (err error) {
  // ...
  defer func() {
    if err1 := gobadIfThisFails(); err1 != nil {
      err = err1
  }()
  // ...
}
```

* see https://www.joeshaw.org/dont-defer-close-on-writable-files/

---

### Project: write tests for `Available`

* any issue?

---

## Structs

Live demo in Playground

---

## Project: custom Twitter & GitHub types

* in package `twitter`, declare a `Twitter` struct type
* in package `github`, declare a `GitHub` struct type

---

# Methods

Live demo in Playground

---

### Project: declare methods on custom types

* turn `Validate` into a method
* turn `Available` into a method
* declare a `String() string` method
* adjust your tests

---

# Interfaces

---

### Interfaces

* abstract type
* describes a type that has certain methods
* an interface is composed of
  * the name of a concrete type
  * a pointer to a value of that type
* THE mechanism for polymorphism in Go

---

### Interfaces are reference types

* zero value: `nil`
* pointers to interfaces? rarely (if ever) useful

---

### Gotcha: `nil` and interfaces

* An interface that holds a `nil` pointer to a concrete type is not itself `nil`!
* see https://golang.org/doc/faq#nil_error

---

### Interface satisfaction

* easy: declare the right methods on the concrete type
* verified at compile time
* implicit!
* compare to Java, PHP, C#, etc.
```
class PokerHand implements Comparable<? super PokerHand> // Java
class Template implements iTemplate                      // PHP
class BoomerangCollection : IEnumerable                  // C#
```
---

### Interfaces as contracts

* some interfaces ask implementors more than satisfying a signature
* example: [`io.Reader`](https://golang.org/pkg/io/#Reader)
* not enforcible by the compiler!
* write tests to check that the implementor is lawful!

---

### Keep interfaces small

* favour one-method interfaces
* easier to satisfy
* many examples in the standard library

---

### Naming of single-method interfaces

* method + "er"
```
type Climber interface {
    Climb(int)
}
```

---

### Interface composition

```
type ReadWriteCloser interface {
  Reader
  Writer
  Closer
}
```
* small interfaces are convenient for composition!

---

## Notable interface types

---

### The empty interface

* `interface{}`
* satisfied by _any_ type
* useful in some cases (`fmt.Println`)
* but not very typesafe
* avoid, if possible

---

### `fmt.Stringer`

```
type Stringer interface {
  String() string
}
```
* governs how `fmt` functions print values of a concrete type

---

## Project: satisfy `fmt.Stringer`

* Do `twitter.Twitter` and `github.GitHub` satisfy it?
* If not, make it so!

---

### `error`

```
type error interface {
  Error() string
}
```
* for any type meant to represent an error value

---

## Project: define a custom error type

* define (where?)
```go
type ErrUnknownAvailability struct {
  Username string
  Platform string
}
```
* make it satisfy the `error` interface
* use it in `twitter.Available`

---

### `sort.Interface`

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```
* analogous to Java's `Comparable` interface

---

### `io.Reader`

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```
* source of `[]byte`
* satisfied by `*os.File`, and many others

---

### `io.Writer`

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```
* sink of `[]byte`
* satisfied by `*os.File`, and many others

---

### Project: define `Validator` interface

```go
type Validator interface {
  Validate(username string) bool
}
```

* Do `twitter.Twitter` and `github.GitHub` satisfy it?
* If not, make it so!


---

### Project: define `Availabler` interface

```go
type Validator interface {
  Validate(username string) bool
}
```

* Do `twitter.Twitter` and `github.GitHub` satisfy it?
* If not, make it so!

---

## Project: define `Checker` interface

* ...by composition of
  * a `fmt.Stringer`
  * a `Validator`
  * an `Availabler`
* Do `twitter.Twitter` and `github.GitHub` satisfy it?

---

### Favour interface parameters

* if possible, take interface types as function parameters, rather than concrete types

```go
type Tree struct {...}
func (t *Tree) Save(f *os.File) error
```

---

### Favour interface parameters

* if possible, take interface types as function parameters, rather than concrete types
* more flexible
* easier to test!

```go
type Tree struct {...}
func (t *Tree) Save(w io.Writer) error
```

---

### Require no more, promise no less

* only ask for the required behaviour from interface parameters

```go
func (t *Tree) Save(rw ReadWriter) error // good?
```

---

### Require no more, promise no less

* only ask for the required behaviour
* similar to DDD principle of "keeping ports small"

```go
func (t *Tree) Save(w Writer) error // better!
```

---

### Wrapping errors

* error-translation idiom
* present a high-level error...
* ... but allow programmatic access to lower-level cause
* standardised in Go 1.13: see https://golang.org/pkg/errors/

---

### Project: wrap low-level http error into custom error type

* augment `ErrUnknownAvailability` with a `Cause error` field
* make it satisfy the following interface
```go
type wrapper interface {
  Unwrap() error
}
```

---

### Asserting on behaviour

* given an interface value, you can ask whether
  the concrete type also satisfy another interface
```go
var b Barista = &Ray{}
b.PrepareCoffee()
c, ok := ray.(Cashier)
if ok {
  c.AcceptPayment()
}
```

---

## Project: figure out why twitter.Available failed

* in `main.go`
```go
available, err := tw.Available("babar")
err, ok := err.(wrapper) // does err satisfy wrapper?
if ok {
  cause := err.Unwrap()
  // inspect cause...
}
```

---

### Type switch

```
func do(i interface{}) {
  switch v := i.(type) { // what's the type of i?
  case int:
    //...
  case string:
    //...
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}
```

* strong coupling with the concrete types :(
* avoid type switches on types you don't own

---

### Project: list available checkers

* in `namecheck.go`, define
```
func Checkers() []Checker {...}
```
