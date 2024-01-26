package main

// START OMIT
type T int

func (t T) Val()  {} // method with value receiver ("value method")
func (t *T) Ptr() {} // method with pointer receiver ("pointer method")

func main() {
	var v T = 42
	p := &v // p is of type *T

	v.Val()
	p.Ptr()
	p.Val() // syntactic sugar: the compiler implicitly derefences p
	v.Ptr() // syntactic sugar: the compiler implicitly takes the address of variable f
	// T(42).Ptr() // compilation error! T(42) is not an addressable value
}

// END OMIT
