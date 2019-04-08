package main

import (
	"fmt"
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

// Deactivate reset all the neurons
func (sp *SpatialPooler) Deactivate() {
	for _, neuron := range sp.getAllNeurons() {
		if neuron.Active {
			neuron.PreviouslyActive = true
		}
		neuron.Active = false
	}
}

// Depredict reset all the neurons predictive state
func (sp *SpatialPooler) Depredict() {
	for _, neuron := range sp.getAllNeurons() {
		neuron.Predictive = false
		neuron.PreviouslyActive = false
	}
}

// Activate the neurons in the spatial pooler for an enoded input
func (sp *SpatialPooler) Activate(connectionThreshold int, overlapThreshold int, learning bool) {
	sp.Deactivate()
	currentlyActive := map[string]bool{}
	previouslyActive := sp.getPreviouslyActive()

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

			predictive := neuron.GetPredictive()
			if len(predictive) > 0 {
				for _, n := range predictive {
					n.Active = true
					n.Predictive = false
					for _, d := range n.DistalInputs {
						_, ok := previouslyActive[d.ConnectedNeuronID]
						if ok {
							d.IncPermanence()
						} else {
							d.DecPermanence()
						}
					}
				}
			} else {
				currentlyActive[neuron.ID] = true
				neuron.Active = true
				for _, miniNeuron := range neuron.MiniColumnNeurons {
					currentlyActive[miniNeuron.ID] = true
					miniNeuron.Active = true
				}
			}
			// learn
			// this loop can probably be sped up by using the GetDendrite function
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

	sp.Depredict()
	sp.activatePredictive(currentlyActive, connectionThreshold, overlapThreshold)

}

func (sp *SpatialPooler) getPreviouslyActive() map[string]bool {
	result := map[string]bool{}
	for _, n := range sp.getAllNeurons() {
		if n.PreviouslyActive {
			result[n.ID] = true
		}
	}
	return result
}
func (sp *SpatialPooler) activatePredictive(currentlyActive map[string]bool, connectionThreshold int, overlapThreshold int) {
	for _, neuron := range sp.getAllNeurons() {
		score := 0
		for _, d := range neuron.DistalInputs {
			_, ok := currentlyActive[d.ConnectedNeuronID]
			if ok && d.Permanence >= connectionThreshold {
				score++
			}
		}
		if score >= overlapThreshold {
			neuron.Predictive = true
		}
	}

}

// Print to the command line
func (sp *SpatialPooler) Print(width int, height int) {
	for i, n := range sp.Neurons {
		for _, nm := range append(append([]*Neuron{}, n), n.MiniColumnNeurons...) {
			fmt.Printf("%s", nm.ID)
			if nm.Active {
				fmt.Printf(" A")
			}
			if nm.Predictive {
				fmt.Printf(" P")
			}
			distal := ""
			for _, d := range nm.DistalInputs {
				distal += fmt.Sprintf("%s:%d,", d.ConnectedNeuronID, d.Permanence)
			}
			fmt.Printf(" (%s) ", distal)
		}
		for c, inputNeuron := range sp.inputNeurons {
			if c%width == 0 {
				fmt.Print("\n")
			}
			if sp.Neurons[i].IsConnected(inputNeuron) {
				fmt.Printf("|%d| ", sp.Neurons[i].GetDendrite(inputNeuron).Permanence)
			} else {
				fmt.Print("|_| ")
			}
		}
		fmt.Print("\n")
	}
}
