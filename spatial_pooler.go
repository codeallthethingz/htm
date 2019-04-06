package main

import (
	"fmt"
)

// SpatialPooler is a set of neurons connecting to an input space
type SpatialPooler struct {
	Neurons      []*Neuron `json:"neurons"`
	inputNeurons []*Neuron //  reference list so you don't have to pass in the input everytime.
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(spatialPoolerSize int, inputSpacePotentialPoolPercent int, inputNeurons []*Neuron) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Neurons:      make([]*Neuron, spatialPoolerSize),
		inputNeurons: inputNeurons,
	}
	for i := 0; i < len(spatialPooler.Neurons); i++ {
		spatialPooler.Neurons[i] = NewNeuron(fmt.Sprintf("c%d", i), inputSpacePotentialPoolPercent, inputNeurons)
	}
	return spatialPooler
}

// Activate the neurons in the spatial pooler for an enoded input
func (sp *SpatialPooler) Activate(connectionThreshold int, overlapThreshold int, learning bool) {
	for _, neuron := range sp.Neurons {
		score := 0
		for _, dendrite := range neuron.ProximalInputs {
			for _, inputNeuron := range sp.inputNeurons {
				if inputNeuron.Active && inputNeuron == dendrite.ConnectedNeuron {
					if dendrite.Permanence >= connectionThreshold { // TODO this line could be 2 lines up speeding up things
						score++
					}
				}
			}
		}
		neuron.Score = score
		if score >= overlapThreshold {
			neuron.Active = true

			// learn
			if learning {
				for _, dendrite := range neuron.ProximalInputs {
					for _, inputNeuron := range sp.inputNeurons {
						if inputNeuron == dendrite.ConnectedNeuron {
							if inputNeuron.Active {
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
		for c, inputNeuron := range sp.inputNeurons {
			if c%width == 0 {
				fmt.Print("\n")
			}
			if sp.Neurons[i].IsConnected(inputNeuron) {
				fmt.Print(sp.Neurons[i].GetDendrite(inputNeuron).Permanence)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
