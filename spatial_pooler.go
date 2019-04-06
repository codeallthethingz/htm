package main

import (
	"fmt"
)

// SpatialPooler is a set of neurons connecting to an input space
type SpatialPooler struct {
	Neurons          []*Neuron
	ActivatedNeurons map[int]bool
	InputSpace       []*Neuron // handy reference list.
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(spatialPoolerSize int, inputSpacePotentialPoolPercent int, inputSpace []*Neuron) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Neurons:          make([]*Neuron, spatialPoolerSize),
		ActivatedNeurons: map[int]bool{},
		InputSpace:       inputSpace,
	}
	for i := 0; i < len(spatialPooler.Neurons); i++ {
		spatialPooler.Neurons[i] = NewNeuron(fmt.Sprintf("c%d", i), inputSpacePotentialPoolPercent, inputSpace)
	}
	return spatialPooler
}

// Activate the neurons in the spatial pooler for an enoded input
func (sp *SpatialPooler) Activate(connectionThreshold int, overlap int, learning bool) {
	for i, neuron := range sp.Neurons {
		score := 0
		for _, dendrite := range neuron.ProximalInputs {
			for _, encode := range sp.InputSpace {
				if encode.Active && encode == dendrite.ConnectedNeuron {
					if dendrite.Permanence >= connectionThreshold {
						score++
					}
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
					for _, encode := range sp.InputSpace {
						if encode == dendrite.ConnectedNeuron {
							if encode.Active {
								dendrite.IncPermanence()
							} else {
								dendrite.DecPermanence()
							}
						}
					}
				}
			}
		}
	}
}

// Print to the command line
func (sp *SpatialPooler) Print(width int, height int) {
	for i := 0; i < len(sp.Neurons); i++ {
		fmt.Printf("neuron: %d", i)
		for c, neuron := range sp.InputSpace {
			if c%width == 0 {
				fmt.Print("\n")
			}
			if sp.Neurons[i].IsConnected(neuron) {
				fmt.Print(sp.Neurons[i].GetDendrite(neuron).Permanence)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
