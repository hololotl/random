package algoritms

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_algo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	TestConstant()
	fmt.Println()
	TestConstPlusUniNoise()
	fmt.Println()
	TestNormal()
}
