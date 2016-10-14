package main

import (
	"fmt"
	"sort"
)

func find95(s []float64) float64 {
	nlen := float64(len(s)-1) * .95
	cut := len(s) - int(nlen)
	if cut <= 1 {
		return s[0] // panic for empty array, as designed
	}
	x := append(make([]float64, 0, cut), s[:cut]...) // a cut with 5% of elements
	sort.Float64s(x)
	for _, next := range s[cut:] {
		if next > x[0] {
			x[0] = next
			for i := 1; i < cut; i++ {
				if x[i] < next {
					x[i-1], x[i] = x[i], x[i-1]
				} else {
					break
				}
			}
		}
	}

	y := nlen - float64(int(nlen)) // min is 0, never 1
	return x[0]*(1-y) + x[1]*y
}

func main() {
	fmt.Println(find95([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(find95([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(find95([]float64{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(find95([]float64{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(find95([]float64{1, 2, 3, 4, 5, 6}))
	fmt.Println(find95([]float64{1, 2, 3, 4, 5}))
	fmt.Println(find95([]float64{1, 2, 3, 4}))
	fmt.Println(find95([]float64{1, 2, 3}))
	fmt.Println(find95([]float64{1, 2}))
	fmt.Println(find95([]float64{1}))
	fmt.Println(find95([]float64{}))
}
