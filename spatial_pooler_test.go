package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSpatialPooler(t *testing.T) {
	spatialPooler := NewSpatialPooler(4, 100, 4)
	require.Equal(t, 4, len(spatialPooler.Cells))
	for i := 0; i < 4; i++ {
		require.Equal(t, 4, len(spatialPooler.Cells[0].Coordinates))
	}
}

func TestNewSpatialPoolerConnectionPool(t *testing.T) {
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(4, 50, 4)
	require.Equal(t, 4, len(spatialPooler.Cells))
	for i := 0; i < 4; i++ {
		require.Equal(t, 2, len(spatialPooler.Cells[0].Coordinates))
	}
}

func TestActivate(t *testing.T) {
	spatialPooler := NewSpatialPooler(4, 100, 4)
	spatialPooler.Activate("XXXX", 0, 2, false)
	for i := 0; i < 4; i++ {
		require.True(t, spatialPooler.Cells[i].Active)
	}
}
func TestLearn(t *testing.T) {
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(4, 100, 4)
	initial1 := spatialPooler.Cells[3].Permanences[1]
	initial2 := spatialPooler.Cells[3].Permanences[2]
	spatialPooler.Activate("XX  ", 4, 2, true)
	spatialPooler.Print(2, 2)
	after1 := spatialPooler.Cells[3].Permanences[1]
	after2 := spatialPooler.Cells[3].Permanences[2]
	require.True(t, after1 > initial1)
	require.True(t, after2 < initial2)
}
