package main

import (
	"math/rand"
)

// UniqueRand keeps track of generated numbers
type UniqueRand struct {
	generated map[int]bool //keeps track of
	rng       *rand.Rand   //underlying random number generator
	scope     int          //scope of number to be generated
}

//NewUniqueRand unique rand less than N
//If N is less or equal to 0, the scope will be unlimited
//If N is greater than 0, it will generate (-scope, +scope)
//If no more unique number can be generated, it will return -1 forwards
func NewUniqueRand(scope int) *UniqueRand {
	return &UniqueRand{
		generated: map[int]bool{},
		scope:     scope,
	}
}

// Reset empties out the chosen values
func (u *UniqueRand) Reset() {
	u.generated = map[int]bool{}
}

// Int get a new unique random number
func (u *UniqueRand) Int() int {
	if u.scope > 0 && len(u.generated) >= u.scope {
		return -1
	}
	for {
		var i int
		if u.scope > 0 {
			i = rand.Int() % u.scope
		} else {
			i = rand.Int()
		}
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}
