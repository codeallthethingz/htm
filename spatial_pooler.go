package main

import (
	"fmt"
)

// SpatialPooler is a set of neurons connecting to an input space
type SpatialPooler struct {
	Neurons          []*Neuron
	ActivatedNeurons map[int]bool
}

// Activate the neurons in the spatial pooler for an encoded input
func (sp *SpatialPooler) Activate(encoded string, connectionThreshold int, overlap int, learning bool) {
	for i, neuron := range sp.Neurons {
		score := 0
		for _, dendrite := range neuron.ProximalInputs {
			if encoded[dendrite.InputCoordinate] == "X"[0] {
				if dendrite.Permanence >= connectionThreshold {
					score++
				}
			}
		}
		neuron.Score = score
		if score >= overlap {
			sp.ActivatedNeurons[i] = true
			neuron.Active = true

			// learn
			if learning {
				for _, dendrite := range neuron.ProximalInputs {
					if encoded[dendrite.InputCoordinate] == "X"[0] {
						dendrite.IncPermanence()
					} else {
						dendrite.DecPermanence()
					}
				}
			}
		}
	}
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(spatialPoolerSize int, inputSpacePotentialPoolPercent int, inputSpaceSize int) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Neurons:          make([]*Neuron, spatialPoolerSize),
		ActivatedNeurons: map[int]bool{},
	}
	for i := 0; i < len(spatialPooler.Neurons); i++ {
		spatialPooler.Neurons[i] = NewNeuron(fmt.Sprintf("c%d", i), inputSpacePotentialPoolPercent, inputSpaceSize)
	}
	return spatialPooler
}

// Print to the command line
func (sp *SpatialPooler) Print(width int, height int) {
	for i := 0; i < len(sp.Neurons); i++ {
		fmt.Printf("neuron: %d", i)
		for c := 0; c < width*height; c++ {
			if c%width == 0 {
				fmt.Print("\n")
			}
			if sp.Neurons[i].IsConnected(c) {
				fmt.Print(sp.Neurons[i].GetDendrite(c).Permanence)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
