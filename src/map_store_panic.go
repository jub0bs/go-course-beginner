package main

func main() {
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "green": "#00ff00"}
	// START OMIT
	m = nil
	m["white"] = "#ffffff" // HL
	// END OMIT
}
