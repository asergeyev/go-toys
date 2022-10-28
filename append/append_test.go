package append

import "testing"

// inspired by https://marin-basic.com/posts/tips-for-using-make-in-go/
const size = 1_000_000

func BenchmarkSliceWithSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, size)
		for j := 0; j < size; j++ {
			slice[j] = j
		}
	}
}

func BenchmarkSliceWithoutSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var slice []int
		for j := 0; j < size; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkSliceWithSizeAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, size)
		for j := 0; j < size; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkSliceWithoutSizeWithMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0)
		for j := 0; j < size; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkSliceSmallSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, 343)
		for j := 0; j < size; j++ {
			slice = append(slice, j)
		}
	}
}
