package main

import "math/rand"

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	ProximalInputs []*Dendrite
	CoordLookup    map[int]int
	Score          int
	ID             string
	Active         bool
}

// NewNeuron creates an initialized neuron
func NewNeuron(id string, inputSpacePotentialPoolPercent int, inputSpaceSize int) *Neuron {
	n := &Neuron{
		ID:             id,
		CoordLookup:    map[int]int{},
		ProximalInputs: []*Dendrite{},
	}
	maxConnections := int(float32(inputSpaceSize) * (float32(inputSpacePotentialPoolPercent) / 100))
	position := 0
	inputSpaceRandom := NewUniqueRand(inputSpaceSize)
	for j := 0; j < inputSpaceSize && len(n.ProximalInputs) < maxConnections; j++ {
		if rand.Int()%100 < inputSpacePotentialPoolPercent {
			newCoord := inputSpaceRandom.Int()
			permanence := rand.Int() % 10
			n.ProximalInputs = append(n.ProximalInputs, &Dendrite{
				inputCoordinate: newCoord,
				Permanence:      permanence,
			})
			n.CoordLookup[newCoord] = position
			position++
		}
	}
	return n
}

// IsConnected is this neuron connected to the coordinate input
func (n *Neuron) IsConnected(coord int) bool {
	_, ok := n.CoordLookup[coord]
	return ok
}

// GetPermanence get the permanance value for this connection
func (n *Neuron) GetPermanence(coord int) int {
	index, _ := n.CoordLookup[coord]
	return n.ProximalInputs[index].Permanence
}

// IncPermanence increases the permanance value for this connection
func (n *Neuron) IncPermanence(coord int) {
	index, _ := n.CoordLookup[coord]
	n.ProximalInputs[index].IncPermanence()
}

// DecPermanence decrease the permanance value for this connection
func (n *Neuron) DecPermanence(coord int) {
	index, _ := n.CoordLookup[coord]
	n.ProximalInputs[index].DecPermanence()
}
