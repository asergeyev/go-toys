package potter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPotting(t *testing.T) {
	x := NewPot(10)
	if len(x.ln) != 13 {
		t.Error("Wrong initial size, expected 13 got ", len(x.ln))
	}

	seen := map[uint32]bool{}
	for i := byte(0); i < 10; i++ {
		num := x.Encode([]byte{'a' + i})
		if seen[num] {
			t.Errorf("Same number as seen before for new value!")
		} else {
			seen[num] = true
		}
		if x.points != uint32(10-i-1) {
			t.Errorf("Points counter is wrong should've had %d but got %d", 10-i-1, x.points)
		}
		if len(x.ln) != 13 {
			t.Error("Resize triggered when not expected!", len(x.ln), "slot", i)
		}

	}

	for i := byte(0); i < 10; i++ {
		bt := []byte{'a' + i}
		pt := x.Encode(bt)
		if !seen[pt] {
			t.Errorf("New number for previously seen value!")
		}
		if got := x.Decode(pt); !bytes.Equal(bt, got) {
			t.Errorf("Expected %s and got %s", pt, got)
		}
		if x.points != 0 {
			t.Errorf("Points counter is wrong should've had 0 but got %d", x.points)
		}
		if len(x.ln) != 13 {
			t.Error("Resize triggered when not expected!", len(x.ln), "slot", i)
		}
	}

	x.Encode([]byte("ok"))
	// should've resized
	if len(x.ln) != 21 {
		t.Error("Wrong resized size, expected 21 got ", len(x.ln))
	}
	if x.points != 6-1 { // added 6 and 1 was taken
		t.Errorf("Did not get enough points, expected %d got %d", 5, x.points)
	}
}

// benchmarks on 100K value population

func BenchmarkPottingNew(b *testing.B) {
	x := NewPot(99999) // one resize after things done
	for i := 0; i < 100000; i++ {
		x.Encode([]byte(fmt.Sprint(i)))
	}
	for i := 0; i < b.N; i++ { // treating i as uint32
		x.Encode([]byte(fmt.Sprint("abcde", i)))
	}
}

func BenchmarkPottingOld(b *testing.B) {
	x := NewPot(100000) // one resize after things done
	for i := 0; i < 100000; i++ {
		x.Encode([]byte(fmt.Sprint(i)))
	}
	for j := 0; j < b.N; j++ {
		i := j % 100000
		x.Encode([]byte(fmt.Sprint(i)))

	}
}

// func BenchmarkPottingOld(b *testing.B) {

// }

// func BenchmarkPottingRandom(b *testing.B) {

// }
