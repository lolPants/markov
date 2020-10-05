package markov

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func weightedRandomIdx(weights []int) int {
	var totalWeight int
	// var i int
	// var random float64

	for _, weight := range weights {
		totalWeight += weight
	}

	random := rand.Float64() * float64(totalWeight)
	for i := range weights {
		f64 := float64(weights[i])
		if random < f64 {
			return i
		}

		random -= f64
	}

	return -1
}
