package p95

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	vhist "github.com/VividCortex/gohistogram"
	tdigest "github.com/caio/go-tdigest"
	chist "github.com/circonus-labs/circonusllhist"
	gkhist "github.com/dgryski/go-gk"
	kllhist "github.com/dgryski/go-kll"
	dhist "github.com/dgryski/go-linlog"

	"testing"
)

const NUMSTREAMING = 288 * 30 // much longer now

var (
	DUMP, TESTSET []float64
	EXPECTED      float64
)

func init() {
	if len(DUMP) > 0 { // needs test data file with values
		TESTSET = DUMP
	} else {
		TESTSET = make([]float64, NUMSTREAMING)
		rand.Seed(int64(time.Now().Nanosecond()))
	}
	for i := 0; i < NUMSTREAMING; i++ {
		// we assume distribution has no impact on quality of streaming eastimators
		// range here is like 1-10GB
		TESTSET[i] = float64(rand.Int63n(9E7)*100) + 1E9
	}
	EXPECTED = find95usual(TESTSET)
}

func TestVividcortexP95Correctness(t *testing.T) {
	runtime.ReadMemStats(m)
	before := m.TotalAlloc

	var x = vhist.NewHistogram(150)
	for _, v := range TESTSET {
		x.Add(v)
	}

	p95 := x.Quantile(.95)
	runtime.ReadMemStats(m)
	diff := p95 - EXPECTED

	t.Logf("Error in vhist-150 estimation: %.3f   %d allocated", diff, m.TotalAlloc-before)
}

func TestCirconuslabsP95Correctness(t *testing.T) {
	runtime.ReadMemStats(m)
	before := m.TotalAlloc
	var x = chist.NewNoLocks()
	for _, v := range TESTSET {
		x.RecordValue(v)
	}

	p95, _ := x.ApproxQuantile([]float64{.95})
	runtime.ReadMemStats(m)

	diff := p95[0] - find95usual(TESTSET)

	t.Logf("Error in chist-%d estimation: %.0f   %d allocated", len(x.DecStrings()), diff, m.TotalAlloc-before)
}

var m = new(runtime.MemStats)

func TestDamianP95Correctness(t *testing.T) {
	runtime.ReadMemStats(m)
	before := m.TotalAlloc
	var x = dhist.NewHistogram(10E9, 8, 8)
	for _, v := range TESTSET {
		x.AtomicInsert(uint64(v))
	}

	var p95 float64

	total := uint64(0)
	stop := uint64(float64(NUMSTREAMING) * .95)
	bins := x.Bins()
	for _, b := range bins {
		total += b.Count
		if total > stop {
			break
		}
		p95 = float64(b.Size)
	}

	runtime.ReadMemStats(m)

	diff := p95 - EXPECTED
	t.Logf("Error in dhist-8x8 estimation: %.0f  %d allocated", diff, m.TotalAlloc-before)
}

func TestGKP95Correctness(t *testing.T) {
	runtime.ReadMemStats(m)
	before := m.TotalAlloc
	var x = gkhist.New(0.5)
	for _, v := range TESTSET {
		x.Insert(v)
	}

	p95 := x.Query(.95)
	runtime.ReadMemStats(m)
	diff := p95 - EXPECTED

	t.Logf("Error in gk-0.5 estimation: %.0f   %d allocated", diff, m.TotalAlloc-before)
}

func TestKLLP95Correctness(t *testing.T) {
	runtime.ReadMemStats(m)
	before := m.TotalAlloc
	var x = kllhist.New(150)
	for _, v := range TESTSET {
		x.Update(v)
	}

	p95 := x.CDF().Query(.95)
	runtime.ReadMemStats(m)
	diff := p95 - EXPECTED

	t.Logf("Error in KLL-150 estimation: %.0f   %d allocated", diff, m.TotalAlloc-before)
	if testing.Short() {
		var m = new(runtime.MemStats)
		runtime.ReadMemStats(m)
		fmt.Println(m.TotalAlloc)
	}
}

func TestTDigestP95Correctness(t *testing.T) {
	runtime.ReadMemStats(m)
	before := m.TotalAlloc
	var x, _ = tdigest.New(tdigest.Compression(200))
	for _, v := range TESTSET {
		x.Add(v)
	}

	p95 := x.Quantile(.95)
	runtime.ReadMemStats(m)
	diff := p95 - EXPECTED

	t.Logf("Error in tdigest-190 estimation: %.0f   %d allocated", diff, m.TotalAlloc-before)
}
