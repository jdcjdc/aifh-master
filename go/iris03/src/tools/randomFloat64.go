package tools

import (
	"time"
	"math/rand"
)

func RandomFloat64(min int32, max int32) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
 	rndNbr := rand.Float64()
	diff := max - min
	rndNbr *= float64(diff)
	rndNbr += float64(min)
	return rndNbr
}