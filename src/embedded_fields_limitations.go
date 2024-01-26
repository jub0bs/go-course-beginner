package main

// START OMIT
type (
	Shape    int
	ShapePtr *Shape // underlying type *Shape is a pointer type
	ShapeI   any    // underlying type is an interface
)

type Foo struct {
	*Shape   // ok
	ShapePtr // compilation error: embedded field type cannot be a pointer
	*ShapeI  // compilation error: embedded field type cannot be a pointer to an interface
}

// END OMIT

func main() {}
