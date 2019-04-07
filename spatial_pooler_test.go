package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMiniColumns(t *testing.T) {
	inputNeurons := MakeInputNeurons(2, 2)
	spatialPooler := NewSpatialPooler(2, 4, 1, inputNeurons)
	for i := 0; i < 4; i++ {
		require.Equal(t, 2, len(spatialPooler.Neurons[i].MiniColumnNeurons))
	}
}
func TestDistalConnections(t *testing.T) {
	rand.Seed(5)
	inputNeurons := MakeInputNeurons(3, 3)
	temporalMemorySize := 4
	spatialPoolerSize := 9
	potentialPoolPercent := 0.4
	spatialPooler := NewSpatialPooler(temporalMemorySize, spatialPoolerSize, potentialPoolPercent, inputNeurons)
	for i := 0; i < spatialPoolerSize; i++ {
		require.Equal(t, 16, len(spatialPooler.Neurons[i].DistalInputs))
	}
	spatialPooler.Print(3, 3)
}

func TestNewSpatialPooler(t *testing.T) {
	inputNeurons := MakeInputNeurons(2, 2)
	spatialPooler := NewSpatialPooler(0, 4, 1, inputNeurons)
	require.Equal(t, 4, len(spatialPooler.Neurons))
	for i := 0; i < 4; i++ {
		require.Equal(t, 4, len(spatialPooler.Neurons[0].ProximalInputs))
	}
}

func TestNewSpatialPoolerConnectionPool(t *testing.T) {
	inputNeurons := MakeInputNeurons(2, 2)
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(0, 4, 0.5, inputNeurons)
	require.Equal(t, 4, len(spatialPooler.Neurons))
	for i := 0; i < 4; i++ {
		require.Equal(t, 2, len(spatialPooler.Neurons[0].ProximalInputs))
	}
}

func TestActivate(t *testing.T) {
	inputNeurons := MakeInputNeurons(2, 2)

	for _, neuron := range inputNeurons {
		neuron.Active = true
	}
	spatialPooler := NewSpatialPooler(0, 4, 1, inputNeurons)
	spatialPooler.Activate(0, 2, false)
	for i := 0; i < 4; i++ {
		require.True(t, spatialPooler.Neurons[i].Active)
	}
}

func TestLearn(t *testing.T) {
	inputNeurons := MakeInputNeurons(2, 2)
	inputNeurons[0].Active = true
	inputNeurons[1].Active = true
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(0, 4, 1, inputNeurons)
	initial1 := spatialPooler.Neurons[3].GetDendrite(inputNeurons[1]).Permanence
	initial2 := spatialPooler.Neurons[3].GetDendrite(inputNeurons[3]).Permanence
	spatialPooler.Print(2, 2)
	spatialPooler.Activate(4, 2, true)
	spatialPooler.Print(2, 2)
	after1 := spatialPooler.Neurons[3].GetDendrite(inputNeurons[1]).Permanence
	after2 := spatialPooler.Neurons[3].GetDendrite(inputNeurons[3]).Permanence
	fmt.Println(initial2, after2)
	require.True(t, after1 > initial1)
	require.True(t, after2 < initial2)
}
