# Packages

---

### Two kinds of packages

* main package
  * defines an executable
  * name of enclosing directory is arbitrary
* library package
  * name of encloding directory must match package name

---

### Composition of a package

* possibly multiple source files (see stdlib)
* all source files go in a folder matching the package name
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
* one package at a time.

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

* good libraries don't panic
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

### Project: list social networks

```
func SocialNetworks() []SocialNetwork {...}
```

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


