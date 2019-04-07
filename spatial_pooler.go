package main

import (
	"fmt"
	"strings"
)

// SpatialPooler is a set of neurons connecting to an input space
type SpatialPooler struct {
	Neurons           []*Neuron `json:"neurons"`
	MiniColumnNeurons []*Neuron `json:"miniColumnNeurons"`
	inputNeurons      []*Neuron //  reference list so you don't have to pass in the input everytime.
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(temporalPoolingSize int, spatialPoolerSize int, potentialPoolPercent float64, inputNeurons []*Neuron) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Neurons:           make([]*Neuron, spatialPoolerSize),
		MiniColumnNeurons: make([]*Neuron, spatialPoolerSize*temporalPoolingSize),
		inputNeurons:      inputNeurons,
	}
	// Create all the neurons
	for i := 0; i < len(spatialPooler.Neurons); i++ {
		id := fmt.Sprintf("c%d", i)
		// Create mini column abstraction - neurons that all fire (burst) when the primary neuron fire or
		// surpress if one is in a predictive state.
		for j := 0; j < temporalPoolingSize; j++ {
			spatialPooler.MiniColumnNeurons[i*temporalPoolingSize+j] = NewNeuron(fmt.Sprintf("%sm%d", id, j), 0, nil, nil)
		}
		spatialPooler.Neurons[i] = NewNeuron(id, potentialPoolPercent, inputNeurons, spatialPooler.MiniColumnNeurons[i*temporalPoolingSize:i*temporalPoolingSize+temporalPoolingSize])
	}

	// Create Distal Connections
	allNeurons := spatialPooler.getAllNeurons()
	for _, neuron := range allNeurons {
		neuron.ConnectDistal(allNeurons, potentialPoolPercent, temporalPoolingSize)
	}

	return spatialPooler
}

func (sp *SpatialPooler) getAllNeurons() []*Neuron {
	return append(sp.Neurons, sp.MiniColumnNeurons...)
}

// Activate the neurons in the spatial pooler for an enoded input
func (sp *SpatialPooler) Activate(connectionThreshold int, overlapThreshold int, learning bool) {
	for _, neuron := range sp.Neurons {
		score := 0
		for _, dendrite := range neuron.ProximalInputs {
			for _, inputNeuron := range sp.inputNeurons {
				if inputNeuron.Active && inputNeuron.ID == dendrite.ConnectedNeuronID {
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
						if inputNeuron.ID == dendrite.ConnectedNeuronID {
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
		for _, p := range sp.Neurons[i].proximalInputLookup {
			fmt.Printf(" %s ", p.ConnectedNeuronID)
		}
		for c, inputNeuron := range sp.inputNeurons {
			if c%width == 0 {
				fmt.Print("\n")
			}
			distal := "["
			for _, d := range sp.Neurons[i].DistalInputs {
				if len(d.ConnectedNeuronID) == 2 {
					distal += "  "
				}
				distal += d.ConnectedNeuronID + ","
			}
			distal += "]"
			if sp.Neurons[i].IsConnected(inputNeuron) {
				fmt.Printf("|%d%s| ", sp.Neurons[i].GetDendrite(inputNeuron).Permanence, distal)
			} else {
				fmt.Print("|" + strings.Repeat("_", len(distal)+1) + "| ")
			}
		}
		fmt.Print("\n")
	}
}
