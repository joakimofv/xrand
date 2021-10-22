package xrand

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestIdx(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		seed := time.Now().Unix()
		t.Log("seed", seed)
		rand.Seed(seed)
		weightedSlice := []float64{0, 0.1, 0.2}
		score := make([]int, len(weightedSlice))
		for n := 0; n < 10000; n++ {
			idx, err := Idx(weightedSlice)
			if err != nil {
				t.Error(err)
				continue
			}
			if idx < 1 || idx > 2 {
				t.Errorf("Unexpected idx %v.", idx)
				continue
			}
			score[idx] += 1
		}
		if score[1] > score[2] {
			t.Errorf("Weighting appearently didn't work. score %v.", score)
		} else {
			t.Log(score)
		}
		weightedSlice = []float64{0.00000000001}
		idx, err := Idx(weightedSlice)
		if err != nil {
			t.Error(err)
		}
		if idx != 0 {
			t.Errorf("Unexpected idx %v.", idx)
		}
	})
	t.Run("basic-fail", func(t *testing.T) {
		weightedSlice := []float64{0, 0, 0}
		_, err := Idx(weightedSlice)
		if err == nil {
			t.Error("Expected error, got nil.")
		} else {
			t.Logf("Got error as expected: %v", err)
		}
		weightedSlice = weightedSlice[:0]
		_, err = Idx(weightedSlice)
		if err == nil {
			t.Error("Expected error, got nil.")
		} else {
			t.Logf("Got error as expected: %v", err)
		}
	})
	t.Run("rand", func(t *testing.T) {
		seed := time.Now().Unix()
		t.Log("seed", seed)
		rng := rand.New(rand.NewSource(seed))
		weightedSlice := []float64{0, 0.1, 0.2}
		score := make([]int, len(weightedSlice))
		for n := 0; n < 10000; n++ {
			idx, err := RandIdx(rng, weightedSlice)
			if err != nil {
				t.Error(err)
				continue
			}
			if idx < 1 || idx > 2 {
				t.Errorf("Unexpected idx %v.", idx)
				continue
			}
			score[idx] += 1
		}
		if score[1] > score[2] {
			t.Errorf("Weighting appearently didn't work. score %v.", score)
		} else {
			t.Log(score)
		}
		// Same result on repeat.
		for k := 0; k < 10; k++ {
			scoreRepeat := make([]int, len(weightedSlice))
			rng.Seed(seed)
			for n := 0; n < 10000; n++ {
				idx, _ := RandIdx(rng, weightedSlice)
				scoreRepeat[idx] += 1
			}
			if diff := cmp.Diff(score, scoreRepeat); diff != "" {
				t.Error(diff)
			}
		}
		weightedSlice = []float64{0.00000000001}
		idx, err := RandIdx(rng, weightedSlice)
		if err != nil {
			t.Error(err)
		}
		if idx != 0 {
			t.Errorf("Unexpected idx %v.", idx)
		}
	})
	t.Run("rand-fail", func(t *testing.T) {
		rng := rand.New(rand.NewSource(0))
		weightedSlice := []float64{0, 0, 0}
		_, err := RandIdx(rng, weightedSlice)
		if err == nil {
			t.Error("Expected error, got nil.")
		} else {
			t.Logf("Got error as expected: %v", err)
		}
		weightedSlice = weightedSlice[:0]
		_, err = RandIdx(rng, weightedSlice)
		if err == nil {
			t.Error("Expected error, got nil.")
		} else {
			t.Logf("Got error as expected: %v", err)
		}
	})
}
