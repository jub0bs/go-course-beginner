# Packages

---

### Two kinds of packages

* main package
  * defines an executable
  * name of enclosing directory is arbitrary
* library package
  * name of enclosing directory must match package name

---

### Composition of a package

* possibly multiple source files (see stdlib)
* all source files go in a folder matching the package name
* at most one package per folder!

---

### Composition of a package (cont'd)

* package members accessible to all source files
* however, each file must import what it needs

---

### Subpackages

* packages defined in subdirectories of a package
* logical grouping of related functionalities
* example: `io/ioutil`
* completely isolated from superpackage

---

### Constraints on dependency graph

* no dependency cycle allowed
* otherwise, compilation error
* key design decision enabling fast compilation!

---

### `init` functions

* special function
* executed at package initialisation
* cannot be called in a program
* tip: avoid execessive side effects

---

### Lifecycle of packages

* consumed packages are imported
* package-level const and vars are initialised
* package's `init` function(s) are executed

---

### Lifecycle of packages (cont'd)

Package initialisation happens
* in a single goroutine
* sequentially
* one package at a time

---

## Consuming a package

---

### Normal import

```
import "fmt"
```
* using default package name

---

### Named import

```
import "math/rand"
import crand "crypto/rand"
```
* renaming a package to avoid conflicts

---

### Blank import

```
import _ "sql/postgres"
```
* import only for initialisation's side effects!
* members of package are not accessible
* used to register specific SQL drivers with `sql` package

---

### dot import

* all exported identifiers usable without qualification
```
import . "fmt"
// ...
Println("foo") // no 'fmt.' needed here
```
* seldom used

---

### relative import

```
import "../twitter"
```
* relative to location on disk
* expedient for quick & dirty prototype
* avoid at all costs!

---

### Access control

* package: the unit of encapsulation
* no encapsulation at the type level
* only two access levels
  * exported (public)
  * unexported (package-private)

---

### Exporting identifiers

* the case of an identifier's initial matters
  * upper case: exported
  * lower case: unexported

---

### Internal packages

* packages named "internal" are special
* an extra mechanism for encapsulation
* only visible to
  * the parent package
  * the parent's subpackages

---

### Project: list available checkers

```
func Checkers() []Checker {...}
```

---

## Designing packages

---

### Naming packages

* a well designed package starts with its name
* lowercase, alphanumeric
* a clue to the functionality provided
* concise, but unambiguous
* favour nouns over verbs

---

### Designing a package

* a focused set of related features
* strive for cohesion
* keep conceptual surface area small

---

### Naming identifiers

* avoid stutter
* some exceptions from stdlib
  * `regexp.Regexp`
  * `time.Time`
* naming factory methods
  * single-type package: `New`
  * multiple-type package: `NewFoo`

---

### Don't panic

* good libraries don't panic (after initialisation)
* report failures to clients as `error`s

---

## Organising packages

---

### Shape of package structure

* flat & wide
* _not_ tall & narrow

---

### Subpackages

* more specific functionalities
* no cycles allowed in dependency graph!
* dependency arrows point upward
* example: `sql` subpackages

---

### Project: CLI username checker

* write an executable
```
namecheck
├── cmd
│   └── cli
│       └── cli.go
...
```
* command-line argument: username
* check on each social network sequentially
* user experience?

---

# Concurrency

---

### Concurrency

* the art of composing a programme in terms of independenly executing computations
* arguably Go's most fascinating aspect
* very important in this era of multicores

---

### Concurrency is not parallelism

* a way of writing a program
* concurrent programme has value even if ultimately executed on a single core

---

### The origin of Go's concurrency

* Tony Hoare's _Communicating Sequential Processes_
* two fundamentals mechanisms:
  * goroutines
  * channels

---

### goroutines

* a concurrent function execution
* `main` function is itself a goroutine
* `go` keyword: used to spawn more goroutines
* Go scheduler: in charge of switching between goroutines

---

### The `go` keyword

* use keyword `go` to spawn a new goroutine
```
go prepareSoufflé()
go peelPotatoes()
```
* golden rule: before launching a goroutine, understand exactly how it will terminate

---

### Wait groups

* `main` does not wait for other goroutines to terminate before it itself terminates
* if tempted to use a sleep, step away from the keyboard and think again

---

###  Wait groups (cont'd)

* `sync.WaitGroup` can be used for
  * keeping track of the number of executing goroutines
  * waiting for that count to go down to 0

---

### Wait-group example

```go
func main() {
  var wg sync.WaitGroup // remark: usable zero value!
  wg.Add(1)             // one more goroutine created
  go func() {
    defer wg.Done()     // one done
    prepareSoufflé()
  }()
  wg.Add(1)             // one more goroutine created
  go func() {
    defer wg.Done()     // one done
    peelPotatoes()
  }()
  wg.Wait()             // ... for all goroutines to be done
}
```

---

### Gotcha: goroutine & func

```go
func printTenIntsConcurrently() {
  var wg sync.WaitGroup
  const n = 10
  wg.Add(n)
  for i := 0; i < n; i++ {
    go func() {
      defer wg.Done()
      fmt.Println(i)
    }()
  }
  wg.Wait()
}
```

* Do you get the output you expect?

---

### Gotcha: goroutine & func

```go
func printTenIntsConcurrently() {
  var wg sync.WaitGroup
  const n = 10
  wg.Add(n)
  for i := 0; i < n; i++ {
    go func(j int) {  // <----
      defer wg.Done()
      fmt.Println(j)
    }(i)              // <----
  }
  wg.Wait()
}
```

---

### Gotcha: goroutine & func

* if possible, avoid using func literals with `go`
* instead, define a dedicated function

```go
func printTenIntsConcurrently() {
  var wg sync.WaitGroup
  const n = 10
  wg.Add(n)
  for i := 0; i < n; i++ {
    go printConcurrently(i, &wg)
  }
  wg.Wait()
}
func printConcurrently(i int, wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Println(i)
}
```

---

## Channels

---

### Share memory by communicating

* _Do not communicate by sharing memory; instead, share memory by communicating._
* channel: mechanism for passing values of a certain type from one goroutine to another

---

### Channel type

* a channel of `T`s has type `chan T`
* channel of channels possible:

```go
var chandler chan chan man
```

---

### Channel: naming conventions

* "ch" (alone or as suffix)
* plural of element type:

```go
var tasks chan task
```

---

### Creating channels

* `ch := make(chan T, c)`
* `c`: optional capacity (defaults to zero)

---

### Channels are reference types

* reference type
* zero value: `nil`

---

### Channel capacity

* a channel has
  * a length: `len(ch)` (seldom used)
  * a capacity: `cap(ch)`

---

### Channel capacity (cont'd)

* a channel of
  * zero capacity is called an _unbuffered channel_
  * positive capacity is called a _buffered channel_

---

### Channel capacity: metaphor

[passe-plats](https://i.skyrock.net/1885/75341885/pics/2950151991_2_3.jpg)

---

## Operations on a channel

* sending a value to channel
* receiving a value from a channel
* ranging over a channel
* closing a channel

---

### Closing a channel

* a (non-`nil`) channel starts its life as open
* a channel can be closed: `close(ch)`
* useful operation when ranging over a channel

---

### Closing a channel (cont'd)

* a closed channel cannot be re-opened
* attempting to close a closed channel: panic!

---

### Closing a channel (cont'd)

* closing a channel does not drain it from the elements it contains
* channels are not resources
  * only close them if the program requires it
  * otherwise, simply let them get GC'd

---

### Sending to a channel

* sending value `v` to channel `ch`:
```
ch <- v
```

---

### Send semantics

* ...depend on the state of `ch`:
  * nil: blocking
  * open
    * not at capacity: adds `v` to `ch`
    * at capacity: blocking
  * closed: panic!

---

### Receiving from a channel

* receiving a value from channel `ch`:
```
<-ch
```
* to store the received value in a variable:
```
v := <-ch
```

---

### Receive semantics

* ...depend on the state of `ch`:
  * nil: blocking
  * not empty: yields an element from `ch`
  * empty:
    * open: blocks
    * closed: yields the zero value of the element type

---

### Dispelling the ambiguity

```go
v, ok := <-ch
```

* `ok` is
  * true if `ch` was still open at the time of receiving `v` from it
  * false if `ch` is closed

---

### Ranging over a channel

* receiving from a channel in a loop
```go
for v := range ch {
  // do something with v
}
```
* loop only terminates once channel is closed

---

### Ranging over a channel (cont'd)

* equivalent to

```go
for {
  v, ok := <-ch
  if !ok {
    break
  }
  // do something with v
}
```

---

### Directional channels

* send-only channel of `T`s has type
```
chan<- T
```
* receive-only channel if `T`s has type
```
<-chan T
```

---

### Directional channels (cont'd)

* useful to what can be done with a channel
* the compiler disallows
  * receiving from a send-only channel
  * sending to a receive-only channel
  * closing a receive-only channel

---

### `chan struct{}`

* can be useful for signaling that some unambiguous event took place!

```go
ch := make(chan struct{})
```

and from another goroutine:

```go
ch <- struct{}{}
```

---

### `select` key word

* analogous to a switch in which the first case that is "ready" gets executed

```
select {
case v := <-ch1:
  // do something with v
case ch2 <- 42:
  // ...
default:
  // ...
}
```

* `default` case
  * is optional
  * executes if none of the other cases is ready

---

### Event loop

```
for {
  select {
  case ...:
    // ...
  case ...:
    // ...
  default:
    // ...
  }
}
```

---

### Concurrency patterns

* we're only scratching the surface of what's possible
* many, many different possibilities with goroutines and channels
* to get an idea, watch
  * [Dave Cheney - Concurrency made easy (GopherCon SG 2017)](https://www.youtube.com/watch?v=yKQOunhhf4A)
  * [Go Concurrency Patterns (Rob Pike, Google I/O 2012)](https://www.youtube.com/watch?v=f6kdp27TYZs)

---

### Communicating by sharing memory

* why? performance, simplicity
* two alternatives to channels
  * atomics
  * mutexes

---

### Atomics

* synchronized operations on fixed-sized integers
* see `sync/atomic` package

---

### Project: number of visits

* keep a count of the number of requests served

---

### Mutexes

* "mutual exclusion"
* allows you to acquire/release a lock on a variable
* see `sync` package

---

### Mutex example: concurrency-safe map

```
var mu sync.Mutex
m := map[string]int
// ...
// inside some goroutine
s := "foo"
mu.Lock()
{
  m[s]++
}
mu.Unlock()
```

* readability tip: enclose the critical section in braces to make it stand out

---

### Project: number of visits

* keep track of how many times each username has been checked

---

### What we haven't covered

* useful third-party libraries
* `context` package
* compiler options
* micro-benchmarks
* concurrency patterns

---

### What we haven't covered (cont'd)

* escape analysis
* profiling
* mechanics of the garbage collector
* mechanics of the Go scheduler

---

# Learning resources

---

### Books

* Kernighan & Donovan - The Go Programming Language
* Brian Ketelsen & Erik St. Martin - Go in Action
* Matt Ryer - Go Programming Blueprints (2nd ed.)

---

### Blogs

* Official Golang blog: https://blog.golang.org
* Dave Cheney: https://dave.cheney.net
* JBD/Rakyll: https://rakyll.org
* Jack Lindamood: https://medium.com/@cep21
* jubobs.com (shameless plug)

---

### Talks

* Anything by Rob Pike, Dave Cheney, Sameer Ajmani, or Matt Ryer
* Marcus Olsson - Building an enterprise service in Go (Golang UK Conference 2016)
* Dave Cheney - Concurrency made easy (GopherCon SG 2017)

---

### Vidcasts

* Francesc Campoy Youtube channel (JustForFunc)

---

### Video courses

* William Kennedy - Ultimate Go v2 (O'Reilly 2018)
* John Graham-Cumming - Introduction to Go Programming (O'Reilly 2014)
* Mike van Sickle's Pluralsight courses

---

### Podcasts

* Go Time: https://changelog.com/gotime

---

# Thank you!
