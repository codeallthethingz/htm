package main

import (
	"math/rand"
)

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	ProximalInputs      []*Dendrite
	proximalInputLookup map[int]int
	Score               int
	ID                  string
	Active              bool
}

// NewNeuron creates an initialized neuron
func NewNeuron(id string, inputSpacePotentialPoolPercent int, inputSpaceSize int) *Neuron {
	connectionPoolSize := int(float32(inputSpaceSize) * (float32(inputSpacePotentialPoolPercent) / 100))
	n := &Neuron{
		ID:                  id,
		proximalInputLookup: map[int]int{},
		ProximalInputs:      make([]*Dendrite, connectionPoolSize),
	}
	inputSpaceRandom := NewUniqueRand(inputSpaceSize)
	for j := 0; j < len(n.ProximalInputs); j++ {
		inputCoordinate := inputSpaceRandom.Int()
		permanence := rand.Int() % 10
		n.ProximalInputs[j] = NewDendrite(inputCoordinate, permanence)
		n.proximalInputLookup[inputCoordinate] = j
	}
	return n
}

// IsConnected is this neuron connected to the coordinate input
func (n *Neuron) IsConnected(coord int) bool {
	_, ok := n.proximalInputLookup[coord]
	return ok
}

// GetDendrite get the permanance value for this connection
func (n *Neuron) GetDendrite(coord int) *Dendrite {
	index, _ := n.proximalInputLookup[coord]
	return n.ProximalInputs[index]
}
