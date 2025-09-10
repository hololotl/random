package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"test/internal"
	"test/internal/models"
	"testing"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Test_Algorithm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	host := "localhost"
	port := "64333"
	log := setupLogger(envLocal)
	rtp := rand.Float64()
	server, err := internal.InitHttpServer(log, host, port, rtp)
	if err != nil {
		log.Error("error creating http server", "error", err)
		return
	}
	go server.Start()
	time.Sleep(5 * time.Second)
	size := 10000
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = 1 + rand.Float64()*(10000-1)
	}
	res := compareConstPlusUniNoise(arr)
	fmt.Println("server rtp:", rtp)
	fmt.Println("client rtp:", res)

}

func GetConstPlusUniNoise() float64 {
	resp, err := http.Get("http://localhost:64333/get")
	if err != nil {
		log.Fatal(err)
		return 0.0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return 0.0
	}
	var r models.Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
		return 0.0
	}
	return r.Result
}

func compareConstPlusUniNoise(arr []float64) float64 {
	var sum float64
	c := 0
	for i := 0; i < len(arr); i++ {
		m := GetConstPlusUniNoise()
		if m > arr[i] {
			sum += arr[i]
			c += 1
		}
	}
	return sum / float64(len(arr))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
