package p95

import (
	"math/rand"
	"time"
)

const (
	NUM              = 8000
	NUMSTREAMING     = 30000 // much longer now
	NOTDETERMINISTIC = false
)

var (
	DUMP, TESTSET []float64
	EXPECTED      float64
)

func init() {
	// yes, confusing but I know what's happening here...
	//
	if !NOTDETERMINISTIC && len(DUMP) > 0 { // needs test data file with values in init()
		TESTSET = DUMP
	} else {
		TESTSET = make([]float64, NUMSTREAMING)
		if NOTDETERMINISTIC {
			rand.Seed(int64(time.Now().Nanosecond()))
		}
		for i := 0; i < NUMSTREAMING; i++ {
			// we assume distribution has no impact on quality of streaming eastimators
			// range here is like 1-10GB
			TESTSET[i] = float64(rand.Int63n(9E7)*100) + 1E9
		}
	}
	EXPECTED = find95usual(TESTSET)
}
