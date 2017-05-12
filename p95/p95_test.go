package p95

import (
	"math"
	"sort"
	"testing"
)

func find95usual(s []float64) float64 {
	if len(s) <= 1 {
		return s[0]
	}
	discard := float64(len(s)-1) * .05
	clone := make([]float64, len(s))
	copy(clone, s)
	sort.Sort(sort.Reverse(sort.Float64Slice(clone)))

	pos := int(discard)         // position on which p95 begins
	y := discard - float64(pos) // min is 0, never 1
	return clone[pos]*(1-y) + clone[pos+1]*y
}

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

	y := nlen - math.Floor(nlen) // min is 0, never 1
	return x[0]*(1-y) + x[1]*y
}

const NUM = 500

func TestP95Correctness(t *testing.T) {
	test := make([]float64, NUM)
	for i := 1; i < NUM; i++ {
		test[i] = float64(i) / 10
		pu := find95usual(test[:i])
		pf := find95(test[:i])
		if math.Abs(pu-pf) > 0.00001 {
			t.Errorf("Incorrect p95, %.5f != %.4f\nN %d %v", pu, pf, i, test[:i])
		}
	}
}

func BenchmarkP95Usual(b *testing.B) {
	test := make([]float64, NUM)
	for i := 0; i < NUM; i++ {
		test[i] = float64(i) / 10
	}
	for i := 0; i < b.N; i++ {
		p := find95usual(test)
		if math.Abs(p-47.4050) > 0.00001 {
			b.Fatalf("Incorrect p95, %.5f\nN %d %v", p, NUM, test)
		}
	}
}

func BenchmarkP95Short(b *testing.B) {
	test := make([]float64, NUM)
	for i := 0; i < NUM; i++ {
		test[i] = float64(i) / 10
	}
	for i := 0; i < b.N; i++ {
		p := find95(test)
		if math.Abs(p-47.4050) > 0.00001 {
			b.Fatalf("Incorrect p95, %.5f\nN %d %v", p, NUM, test)
		}
	}
}
