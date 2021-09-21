package xrand

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// Split splits n into parts that add up to n.
// panics if n < 0 or parts < 1.
func Split(n int, parts int) []int {
	return randSplit(nil, n, parts)
}

// RandSplit is like Split but uses r *rand.Rand as source.
// panics if r is nil.
func RandSplit(r *rand.Rand, n int, parts int) []int {
	if r == nil {
		panic("r is nil")
	}
	return randSplit(r, n, parts)
}

func randSplit(r *rand.Rand, n int, parts int) []int {
	if n < 0 {
		panic(fmt.Errorf("invalid n: %v < 0", n))
	}
	if parts < 1 {
		panic(fmt.Errorf("invalid parts: %v < 1", parts))
	}
	pp := make([]int, parts)
	for i, _ := range pp {
		if r != nil {
			pp[i] = r.Intn(n + 1)
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
// panics if parts < 1.
func SplitUint64(n uint64, parts int) []uint64 {
	return randSplitUint64(nil, n, parts)
}

// RandSplitUint64 is like SplitUInt64 but uses r *rand.Rand as source.
// panics if r is nil.
func RandSplitUint64(r *rand.Rand, n uint64, parts int) []uint64 {
	if r == nil {
		panic("r is nil")
	}
	return randSplitUint64(r, n, parts)
}

func randSplitUint64(r *rand.Rand, n uint64, parts int) []uint64 {
	if parts < 1 {
		panic(fmt.Errorf("invalid parts: %v < 1", parts))
	}

	randUint64n := func() uint64 {
		return uint64(rand.Int63n(int64(n + 1)))
	}
	if r != nil {
		randUint64n = func() uint64 {
			return uint64(r.Int63n(int64(n + 1)))
		}
	}
	if n >= math.MaxInt64 {
		randUint64n = func() uint64 {
			var ret uint64
			for ret = rand.Uint64(); ret > n; ret = rand.Uint64() {
			}
			return ret
		}
		if r != nil {
			randUint64n = func() uint64 {
				var ret uint64
				for ret = r.Uint64(); ret > n; ret = r.Uint64() {
				}
				return ret
			}
		}
	}

	pp := make([]uint64, parts)
	for i, _ := range pp {
		if r != nil {
			pp[i] = randUint64n()
		} else {
			pp[i] = randUint64n()
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
