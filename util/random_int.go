package util

import (
	"math/rand"
	"time"
)

func getRandomInt(maxInt int) (index int) {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(maxInt)
}
