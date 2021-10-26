package xrand

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// Uint64n is like rand.Int63n for uint64.
func Uint64n(n uint64) uint64 {
	if n <= math.MaxInt64 {
		return uint64(rand.Int63n(int64(n)))
	} else {
		var ret uint64
		for ret = rand.Uint64(); ret >= n; ret = rand.Uint64() {
		}
		return ret
	}
}

// Uint64nWithRand is like Uint64n but uses rng *rand.Rand as source.
// panics if rng is nil.
func Uint64nWithRand(rng *rand.Rand, n uint64) uint64 {
	if rng == nil {
		panic("*rand.Rand is nil")
	}
	if n <= math.MaxInt64 {
		return uint64(rng.Int63n(int64(n)))
	} else {
		var ret uint64
		for ret = rng.Uint64(); ret >= n; ret = rng.Uint64() {
		}
		return ret
	}
}

// Split splits n into parts that add up to n.
// panics if n < 0.
func Split(n int, parts int) []int {
	return randSplit(nil, n, parts)
}

// SplitWithRand is like Split but uses rng *rand.Rand as source.
// panics if rng is nil.
func SplitWithRand(rng *rand.Rand, n int, parts int) []int {
	if rng == nil {
		panic("*rand.Rand is nil")
	}
	return randSplit(rng, n, parts)
}

func randSplit(rng *rand.Rand, n int, parts int) []int {
	if n < 0 {
		panic(fmt.Errorf("invalid n: %v < 0", n))
	}
	if parts < 1 {
		return nil
	}
	pp := make([]int, parts)
	for i, _ := range pp {
		if rng != nil {
			pp[i] = rng.Intn(n + 1)
		} else {
			pp[i] = rand.Intn(n + 1)
		}
	}
	pp[0] = 0
	sort.SliceStable(pp, func(i, j int) bool { return pp[i] < pp[j] })
	for i, p := range pp {
		next := n
		if i+1 < len(pp) {
			next = pp[i+1]
		}
		pp[i] = next - p
	}
	return pp
}

// SplitUint64 splits n into parts that add up to n.
func SplitUint64(n uint64, parts int) []uint64 {
	return randSplitUint64(nil, n, parts)
}

// SplitUint64WithRand is like SplitUInt64 but uses rng *rand.Rand as source.
// panics if rng is nil.
func SplitUint64WithRand(rng *rand.Rand, n uint64, parts int) []uint64 {
	if rng == nil {
		panic("*rand.Rand is nil")
	}
	return randSplitUint64(rng, n, parts)
}

func randSplitUint64(rng *rand.Rand, n uint64, parts int) []uint64 {
	if parts < 1 {
		return nil
	}

	pp := make([]uint64, parts)
	for i, _ := range pp {
		if rng != nil {
			if n == math.MaxUint64 {
				pp[i] = Uint64nWithRand(rng, n)
			} else {
				pp[i] = Uint64nWithRand(rng, n+1)
			}
		} else {
			if n == math.MaxUint64 {
				pp[i] = Uint64n(n)
			} else {
				pp[i] = Uint64n(n + 1)
			}
		}
	}
	pp[0] = 0
	sort.SliceStable(pp, func(i, j int) bool { return pp[i] < pp[j] })
	for i, p := range pp {
		next := n
		if i+1 < len(pp) {
			next = pp[i+1]
		}
		pp[i] = next - p
	}
	return pp
}
