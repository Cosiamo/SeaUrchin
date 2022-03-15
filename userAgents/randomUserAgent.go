package userAgents

import (
	"math/rand"
	"time"
)

// this is used so that Google thinks the requests are coming from different browsers
// need this so that Google doesn't think anything shady is going on
func RandomUserAgent() string {
	// select a random number
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(UserAgents)
	// to access a particular value in a slice there needs to be an index passed into an array
	return UserAgents[randNum]
}