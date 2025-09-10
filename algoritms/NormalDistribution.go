package algoritms

import (
	"fmt"
	"math"
	"math/rand"
)

func GetNormal(rtp float64, std float64) float64 {
	res := math.Sqrt(rtp*2*9999 + 1)
	random := rand.NormFloat64()*res*std + res
	return random
}

func TestNormal() {
	var MAE float64  // absolute error
	var MAPE float64 // Mean Absolute Percentage Error
	var PWT float64  // Percent within tolerance ∣y−x∣≤ t
	tryes := 10000
	std := 0.05
	for i := 0; i < tryes; i++ {
		rtp := rand.Float64()
		size := 10000
		arr := make([]float64, size)
		for i := 0; i < size; i++ {
			arr[i] = 1 + rand.Float64()*(10000-1)
		}
		res := compareNormal(arr, rtp, std)
		MAE += math.Abs(rtp - res)
		MAPE += math.Abs(rtp-res) / rtp
		if math.Abs(rtp-res) < rtp*0.15 {
			PWT += 1
		}
	}
	MAE = MAE / float64(tryes)
	fmt.Println("Normal Distribution. MAE:", MAE, "Std:", std)
	fmt.Println("Normal Distribution", PWT/float64(tryes)*100)
	fmt.Println("Normal Distribution", MAPE/float64(tryes)*100)
}

func compareNormal(arr []float64, rtp, std float64) float64 {
	var sum float64
	c := 0
	for i := 0; i < len(arr); i++ {
		m := GetNormal(rtp, std)
		if m > arr[i] {
			sum += arr[i]
			c += 1
		}
	}
	return sum / float64(len(arr))
}
