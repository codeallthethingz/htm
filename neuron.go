package main

import (
	"math/rand"
)

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	ProximalInputs      []*Dendrite
	proximalInputLookup map[*Neuron]int
	DistalInputs        []*Dendrite
	Score               int
	ID                  string
	Active              bool
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

// IsConnected is this neuron connected to the coordinate input
func (n *Neuron) IsConnected(inputCoordinate *Neuron) bool {
	_, ok := n.proximalInputLookup[inputCoordinate]
	return ok
}

// GetDendrite get a dendrite connected to this coordinate
func (n *Neuron) GetDendrite(inputCoordinate *Neuron) *Dendrite {
	index, _ := n.proximalInputLookup[inputCoordinate]
	return n.ProximalInputs[index]
}
