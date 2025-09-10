package algoritms

import (
	"fmt"
	"math"
	"math/rand"
)

func GetConstant(rtp float64) float64 {
	res := math.Sqrt(rtp*2*9999 + 1)
	return res
}

func TestConstant() {
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
		res := compare(arr, m)
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

func compare(arr []float64, m float64) float64 {
	var sum float64
	c := 0
	for i := 0; i < len(arr); i++ {
		if m > arr[i] {
			sum += arr[i]
			c += 1
		}
	}
	return sum / float64(len(arr))
}
