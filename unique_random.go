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
func NewUniqueRand(scope int) *UniqueRand {
	if scope < 2 {
		panic("Must be initialized with a value greater than 1")
	}
	return &UniqueRand{
		generated: map[int]bool{},
		scope:     scope,
	}
}

// Reset empties out the chosen values
func (u *UniqueRand) Reset() {
	u.generated = map[int]bool{}
}

// Int get a new unique random number within scope
func (u *UniqueRand) Int() int {
	if len(u.generated) == u.scope {
		panic("filled all positions in scope")
	}
	for {
		i := rand.Int() % u.scope
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}
