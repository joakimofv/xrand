package xrand

import (
	"errors"
	"math"
	"math/rand"
)

func Index(weightedSlice []float64) (int, error) {
	sum := float64(0)
	for _, elem := range weightedSlice {
		if elem <= 0 || math.IsNaN(elem) {
			continue
		}
		sum += elem
	}
	r := rand.Float64() * sum
	accum := float64(0)
	for _, idx := range rand.Perm(len(weightedSlice)) {
		elem := weightedSlice[idx]
		if elem <= 0 || math.IsNaN(elem) {
			continue
		}
		accum += elem
		if r < accum {
			return idx, nil
		}
	}
	return 0, errors.New("Nothing can be selected.")
}

func IndexWithRand(rng *rand.Rand, weightedSlice []float64) (int, error) {
	if rng == nil {
		return 0, errors.New("*rand.Rand is nil")
	}
	sum := float64(0)
	for _, elem := range weightedSlice {
		if elem <= 0 || math.IsNaN(elem) {
			continue
		}
		sum += elem
	}
	r := rng.Float64() * sum
	accum := float64(0)
	for _, idx := range rng.Perm(len(weightedSlice)) {
		elem := weightedSlice[idx]
		if elem <= 0 || math.IsNaN(elem) {
			continue
		}
		accum += elem
		if r < accum {
			return idx, nil
		}
	}
	return 0, errors.New("Nothing can be selected.")
}
