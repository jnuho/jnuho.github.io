package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5,2,6,3,1,4}
	sort.Ints(s)
	fmt.Println(s)

	s2 := []float64{3.14, 2.22, 5.55, 22.2, 10, 1.1101}
	sort.Float64s(s2)
	fmt.Println(s2)
}