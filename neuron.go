package main

import (
	"math/rand"
)

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	proximalInputLookup map[*Neuron]int
	ProximalInputs      []*Dendrite `json:"proximalInputs"`
	Score               int         `json:"score"`
	ID                  string      `json:"id"`
	Active              bool        `json:"active"`
	Predictive          bool        `json:"predictive"`
}

// NewNeuron creates an initialized neuron
func NewNeuron(id string, inputSpacePotentialPoolPercent int, inputNeurons []*Neuron) *Neuron {
	connectionPoolSize := int(float32(len(inputNeurons)) * (float32(inputSpacePotentialPoolPercent) / 100))
	n := &Neuron{
		ID:                  id,
		proximalInputLookup: map[*Neuron]int{},
		ProximalInputs:      make([]*Dendrite, connectionPoolSize),
	}
	inputNeuronsRandom := NewUniqueRand(len(inputNeurons))
	for j := range n.ProximalInputs {
		inputNeuron := inputNeurons[inputNeuronsRandom.Int()]
		permanence := rand.Int() % 10
		n.ProximalInputs[j] = NewDendrite(inputNeuron, permanence)
		n.proximalInputLookup[inputNeuron] = j
	}
	return n
}

// IsConnected is this neuron connected to the input
func (n *Neuron) IsConnected(inputNeuron *Neuron) bool {
	_, ok := n.proximalInputLookup[inputNeuron]
	return ok
}

// GetDendrite get a dendrite connected to this coordinate
func (n *Neuron) GetDendrite(inputNeuron *Neuron) *Dendrite {
	index, _ := n.proximalInputLookup[inputNeuron]
	return n.ProximalInputs[index]
}
