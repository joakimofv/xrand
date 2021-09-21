package xrand

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSplit(t *testing.T) {
	for n := 0; n < 1000; n++ {
		t.Run(fmt.Sprintf("%v", n), func(t *testing.T) {
			parts := 1 + rand.Intn(n+1)
			pp := Split(n, parts)
			sum := 0
			for _, p := range pp {
				sum += p
			}
			if sum != n {
				t.Errorf("Expected sum %v, got %v.", n, sum)
			}

			seed := time.Now().UnixNano()
			r := rand.New(rand.NewSource(seed))
			pp = RandSplit(r, n, parts)
			sum = 0
			for _, p := range pp {
				sum += p
			}
			if sum != n {
				t.Errorf("Expected sum %v, got %v.", n, sum)
			}
			r.Seed(seed)
			pp2 := RandSplit(r, n, parts)
			for i, p := range pp {
				if p != pp2[i] {
					t.Fatal("RandSplit did not get the same result at reseed.")
				}
			}
		})
	}
}

func TestSplitUint64(t *testing.T) {
	for _, n := range []uint64{0, 1, 2, 10, 100, 1000, 10000, math.MaxInt64 - 1, math.MaxInt64, math.MaxUint64 - 1, math.MaxUint64} {
		t.Run(fmt.Sprintf("%v", n), func(t *testing.T) {
			for count := 0; count < 100; count++ {
				parts := 1 + rand.Intn(5)
				pp := SplitUint64(n, parts)
				sum := uint64(0)
				for _, p := range pp {
					sum += p
				}
				if sum != n {
					t.Errorf("Expected sum %v, got %v.", n, sum)
				}

				seed := time.Now().UnixNano()
				r := rand.New(rand.NewSource(seed))
				pp = RandSplitUint64(r, n, parts)
				sum = uint64(0)
				for _, p := range pp {
					sum += p
				}
				if sum != n {
					t.Errorf("Expected sum %v, got %v.", n, sum)
				}
				r.Seed(seed)
				pp2 := RandSplitUint64(r, n, parts)
				for i, p := range pp {
					if p != pp2[i] {
						t.Fatal("RandSplitUint64 did not get the same result at reseed.")
					}
				}
			}
		})
	}
}
