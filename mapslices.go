package main

import "fmt"

func main() {
	// Make new maps.
	fmt.Println(make(map[string]int)) // map[]
	m := make(map[string]int, 2)
	fmt.Println(m, len(m)) // map[] 0
	m["C"] = 2000
	m["Go"] = 1888
	fmt.Println(m, len(m)) // map[C:1972 Go:2009] 2

	// Make new slices.
	s := make([]int, 2, 4)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 2 4
	s = make([]int, 1)
	fmt.Println(s, len(s), cap(s)) // [0 0] 1 1
}