package internal

import (
	"math"
	"math/rand"
)

type UniAlgorithm struct {
	rtp float64
	std float64
}

func NewUniAlgorithm(rtp float64, std float64) *UniAlgorithm {
	return &UniAlgorithm{rtp: rtp, std: std}
}

func (u *UniAlgorithm) GetMultiplier() float64 {
	res := math.Sqrt(u.rtp*2*9999 + 1)
	random := rand.Float64()*2 - 1
	v := res * u.std * random
	res += v
	return res
}
