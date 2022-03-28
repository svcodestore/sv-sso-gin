package utils

import (
	"math/rand"
	"time"
)

func RandRange(min, max int) (n int) {
	rand.Seed(time.Now().Unix())
	n = rand.Intn(max-min) + min
	return
}
