package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// fmt.Printf("Non-Generic Sums: %v and %v\n",
	// 	SomaInts(ints),
	// 	SomaFloats(floats))

	// fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
	// 	SomaIntsOrFloats(ints),
	// 	SomaIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SomaNumbers(ints),
		SomaNumbers(floats))
}

// Recebe um mapa onde as chaves são string e os valores são int64
// func SomaInts(m map[string]int64) int64 {
// 	var s int64

// 	for _, v := range m {
// 		s += v
// 	}

// 	return s
// }

// func SomaFloats(m map[string]float64) float64 {
// 	var s float64

// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
// func SomaIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
// 	var s V
// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }

func SomaNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
