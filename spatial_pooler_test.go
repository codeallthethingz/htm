package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSpatialPooler(t *testing.T) {
	spatialPooler := NewSpatialPooler(4, 100, 4)
	require.Equal(t, 4, len(spatialPooler.Neurons))
	for i := 0; i < 4; i++ {
		require.Equal(t, 4, len(spatialPooler.Neurons[0].ProximalInputs))
	}
}

func TestNewSpatialPoolerConnectionPool(t *testing.T) {
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(4, 50, 4)
	require.Equal(t, 4, len(spatialPooler.Neurons))
	for i := 0; i < 4; i++ {
		require.Equal(t, 2, len(spatialPooler.Neurons[0].ProximalInputs))
	}
}

func TestActivate(t *testing.T) {
	spatialPooler := NewSpatialPooler(4, 100, 4)
	spatialPooler.Activate("XXXX", 0, 2, false)
	for i := 0; i < 4; i++ {
		require.True(t, spatialPooler.Neurons[i].Active)
	}
}
func TestLearn(t *testing.T) {
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(4, 100, 4)
	initial1 := spatialPooler.Neurons[3].GetPermanence(1)
	initial2 := spatialPooler.Neurons[3].GetPermanence(3)
	spatialPooler.Print(4, 1)
	spatialPooler.Activate("XX  ", 4, 2, true)
	spatialPooler.Print(4, 1)
	after1 := spatialPooler.Neurons[3].GetPermanence(1)
	after2 := spatialPooler.Neurons[3].GetPermanence(3)
	fmt.Println(initial2, after2)
	require.True(t, after1 > initial1)
	require.True(t, after2 < initial2)
}
