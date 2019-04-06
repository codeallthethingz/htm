package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSpatialPooler(t *testing.T) {
	inputNeurons := []*Neuron{
		&Neuron{}, &Neuron{}, &Neuron{}, &Neuron{},
	}
	spatialPooler := NewSpatialPooler(4, 100, inputNeurons)
	require.Equal(t, 4, len(spatialPooler.Neurons))
	for i := 0; i < 4; i++ {
		require.Equal(t, 4, len(spatialPooler.Neurons[0].ProximalInputs))
	}
}

func TestNewSpatialPoolerConnectionPool(t *testing.T) {
	inputNeurons := []*Neuron{
		&Neuron{}, &Neuron{}, &Neuron{}, &Neuron{},
	}
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(4, 50, inputNeurons)
	require.Equal(t, 4, len(spatialPooler.Neurons))
	for i := 0; i < 4; i++ {
		require.Equal(t, 2, len(spatialPooler.Neurons[0].ProximalInputs))
	}
}

func TestActivate(t *testing.T) {
	inputNeurons := []*Neuron{
		&Neuron{Active: true}, &Neuron{Active: true}, &Neuron{Active: true}, &Neuron{Active: true},
	}
	spatialPooler := NewSpatialPooler(4, 100, inputNeurons)
	spatialPooler.Activate(0, 2, false)
	for i := 0; i < 4; i++ {
		require.True(t, spatialPooler.Neurons[i].Active)
	}
}

func TestLearn(t *testing.T) {
	inputNeurons := []*Neuron{
		&Neuron{Active: true}, &Neuron{Active: true}, &Neuron{Active: false}, &Neuron{Active: false},
	}
	rand.Seed(0)
	spatialPooler := NewSpatialPooler(4, 100, inputNeurons)
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
