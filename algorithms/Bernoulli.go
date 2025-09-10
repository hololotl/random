package algorithms

import (
	"fmt"
	"math"
	"math/rand"
)

func GetBernoulli(rtp float64) float64 {
	t := (2 * 9999 * rtp) / (10000*10000 - 1)
	if rand.Float64() < t {
		return 10000
	}
	return 0
}

func TestBernoulli() {
	var MAE float64  // absolute error
	var MAPE float64 // Mean Absolute Percentage Error
	var PWT float64  // Percent within tolerance ∣y−x∣≤ t

	tryes := 10000
	for i := 0; i < tryes; i++ {
		rtp := rand.Float64()
		m := GetConstant(rtp)
		size := 10000
		arr := make([]float64, size)
		for i := 0; i < size; i++ {
			arr[i] = 1 + rand.Float64()*(10000-1)
		}
		res := compareBernoulli(arr, m)
		MAE += math.Abs(rtp - res)
		MAPE += math.Abs(rtp-res) / rtp
		if math.Abs(rtp-res) < rtp*0.15 {
			PWT += 1
		}
	}
	fmt.Println("Constant. Mean absolute error:", MAE/float64(tryes))
	fmt.Println("Constant. Percent within tolerance", PWT/float64(tryes)*100)
	fmt.Println("Constant. Mean Absolute Percentage Error", MAPE/float64(tryes)*100)
}

func compareBernoulli(arr []float64, rtp float64) float64 {
	var sum float64
	c := 0
	for i := 0; i < len(arr); i++ {
		m := GetBernoulli(rtp)
		if m > arr[i] {
			sum += arr[i]
			c += 1
		}
	}
	return sum / float64(len(arr))
}
