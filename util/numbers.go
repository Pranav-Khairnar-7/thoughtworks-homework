package util

import "math/rand"

func GetRandomUpto(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max) + 1
}
