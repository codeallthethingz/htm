package main

import (
	"sort"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInt(t *testing.T) {
	results := make([]int, 10)
	rand := NewUniqueRand(len(results))
	for i := 0; i < len(results); i++ {
		results[i] = rand.Int()
	}
	require.Panics(t, func() { rand.Int() }, "should fail")
	sort.Ints(results)
	for i := 0; i < len(results); i++ {
		require.Equal(t, i, results[i])
	}
}

func TestEdges(t *testing.T) {
	require.Panics(t, func() { NewUniqueRand(-1) }, "should fail")
	require.Panics(t, func() { NewUniqueRand(0) }, "should fail")
	require.Panics(t, func() { NewUniqueRand(1) }, "should fail")
}

func TestReset(t *testing.T) {
	rand := NewUniqueRand(2)
	rand.Int()
	rand.Int()
	require.Panics(t, func() { rand.Int() }, "should fail")
	rand.Reset()
	rand.Int()
	rand.Int()
	require.Panics(t, func() { rand.Int() }, "should fail")
}
