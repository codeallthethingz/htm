package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBigInput(t *testing.T) {
	inputNeurons := MakeInputNeurons(19, 11)
	spatialPooler := NewSpatialPooler(20, 100, 0.4, inputNeurons)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	fmt.Println("about to show phone")
	Encode(phone, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	Encode(cup, inputNeurons, 0.04, 19)
	spatialPooler.Activate(5, 5, 100, true)
	countPredictive(spatialPooler)
	t.Fail()
}

func countPredictive(spatialPooler *SpatialPooler) {
	countP, countA, countT, countPP := 0, 0, 0, 0
	for _, n := range spatialPooler.getAllNeurons() {
		if n.Predictive {
			countP++
		}
		if n.Active {
			countA++
		}
		if n.PreviouslyPredictive {
			countPP++
		}

		countT++
	}
	fmt.Println(len(spatialPooler.Neurons), len(spatialPooler.MiniColumnNeurons))
	fmt.Printf("pre: %d, act: %d prev: %d of %d\n", countP, countA, countPP, countT)
}
func TestPrediction(t *testing.T) {
	rand.Seed(3)
	inputNeurons := MakeInputNeurons(2, 2)

	spatialPooler := NewSpatialPooler(2, 4, .5, inputNeurons)
	for i := 0; i < 4; i++ {
		require.Equal(t, 2, len(spatialPooler.Neurons[i].MiniColumnNeurons))
	}
	spatialPooler.Print(4, 1)
	fmt.Println()

	inputNeurons[0].Active = true
	inputNeurons[1].Active = true
	inputNeurons[2].Active = false
	inputNeurons[3].Active = false
	spatialPooler.Activate(2, 2, true)
	for _, neuron := range spatialPooler.Neurons {
		for _, miniNeuron := range neuron.MiniColumnNeurons {
			require.Equal(t, neuron.Active, miniNeuron.Active)
		}
	}
	distalPermenanceBefore := spatialPooler.Neurons[0].MiniColumnNeurons[1].DistalInputs[0].Permanence
	spatialPooler.Print(4, 1)
	require.Equal(t, 3, distalPermenanceBefore)
	require.Equal(t, true, spatialPooler.Neurons[2].MiniColumnNeurons[1].Predictive)
	fmt.Println()

	inputNeurons[0].Active = true
	inputNeurons[1].Active = true
	inputNeurons[2].Active = false
	inputNeurons[3].Active = true
	spatialPooler.Activate(1, 2, true)
	spatialPooler.Print(4, 1)

	distalPermenanceAfter := spatialPooler.Neurons[0].MiniColumnNeurons[1].DistalInputs[0].Permanence
	require.Equal(t, 4, distalPermenanceAfter)
	require.Equal(t, true, spatialPooler.Neurons[0].MiniColumnNeurons[1].Active)
	require.Equal(t, false, spatialPooler.Neurons[0].Active)
	fmt.Println()

	inputNeurons[0].Active = false
	inputNeurons[1].Active = false
	inputNeurons[2].Active = true
	inputNeurons[3].Active = true
	spatialPooler.Activate(1, 2, true)
	spatialPooler.Print(4, 1)

	require.Equal(t, 8, spatialPooler.Neurons[2].MiniColumnNeurons[1].DistalInputs[0].Permanence)
	fmt.Println()
}
func TestSetupMiniColumns(t *testing.T) {
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
	require.True(t, after1 > initial1)
	require.True(t, after2 < initial2)
}
