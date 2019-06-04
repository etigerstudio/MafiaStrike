package util

import (
	"math/rand"
	"time"
)

func InitRandSeed()  {
	rand.Seed(time.Now().UTC().UnixNano())
}