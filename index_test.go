package xrand

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestIndex(t *testing.T) {
	for _, useRng := range []bool{false, true} {
		name := "default"
		if useRng {
			name = "rand"
		}
		t.Run(name, func(t *testing.T) {
			for name, tc := range map[string]struct {
				weightedSlice []float64
				expectError   bool
			}{
				"basic":           {weightedSlice: []float64{0, 0.1, 0.2}},
				"small":           {weightedSlice: []float64{0.00000000001}},
				"zeroes":          {weightedSlice: []float64{0, 0, 0}, expectError: true},
				"empty":           {weightedSlice: []float64{}, expectError: true},
				"nan":             {weightedSlice: []float64{math.NaN(), math.NaN()}, expectError: true},
				"nan-and-numbers": {weightedSlice: []float64{math.NaN(), 1}},
			} {
				t.Run(name, func(t *testing.T) {
					seed := time.Now().Unix()
					t.Log("seed", seed)
					rng := rand.New(rand.NewSource(seed))

					savedScore := make([]int, len(tc.weightedSlice))
					for n := 0; n < 2; n++ {
						rand.Seed(seed)
						rng.Seed(seed)

						score := make([]int, len(tc.weightedSlice))
						for n := 0; n < 10000; n++ {
							idx, err := Index(tc.weightedSlice)
							if useRng {
								idx, err = IndexWithRand(rng, tc.weightedSlice)
							}
							if tc.expectError {
								if err == nil {
									t.Fatal("Expected error, got nil.")
								}
								return
							}
							if err != nil {
								t.Fatal(err)
							}
							score[idx] += 1
						}

						for idx, elem := range tc.weightedSlice {
							if elem <= 0 || math.IsNaN(elem) {
								if score[idx] > 0 {
									t.Errorf("selected invalid idx %v (%v times).", idx, score[idx])
								}
							} else {
								if score[idx] == 0 {
									t.Errorf("never selected valid idx %v.", idx)
								}
							}
						}

						if n > 0 {
							if diff := cmp.Diff(score, savedScore); diff != "" {
								t.Error(diff)
							}
						}
						copy(savedScore, score)
					}
				})
			}
		})
	}
}
