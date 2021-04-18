package main

import "fmt"

func EarthAge() {
	age := 4.54 * 1000000000
	fmt.Printf("%0.0f\n", age)
	fmt.Printf("%0.0f\n", age+50000000)
	fmt.Printf("%0.0f\n", age-50000000)
	fmt.Printf("%0.0f\n", age*24.5)
}
