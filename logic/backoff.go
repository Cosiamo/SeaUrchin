package logic

import (
	"math/rand"
	"time"
)

var res int

func Backoff()int {
    rand.Seed(time.Now().UnixNano())
    for i := 1; i <= 12; i++ {
		i = rand.Intn(10)
		if i <= 2 {
			// 4, 5, 6
			res = i + 4
			return res
		} else if i > 2 && i <= 6 {
			// 5, 6, 7, 8
			res = i + 2
			return res
		} else {
			// 7, 8, 9
			res = i
			return res
		}
	}

	return res
}