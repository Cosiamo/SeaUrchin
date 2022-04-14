package logic

import (
	"math/rand"
	"time"
)

var res int

func Backoff()int {
    rand.Seed(time.Now().UnixNano())
    for i := 1; i <= 6; i++ {
		i = rand.Intn(5)
		if i <= 2 {
			// 1, 2, 3
			res = i + 1
			return res
		} else {
			// 3, 4, 5
			res = i
			return res
		}
	}

	return res
}