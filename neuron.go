package main

import "math/rand"

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	Coordinates []int
	CoordLookup map[int]int
	Permanences []int
	Score       int
	ID          string
	Active      bool
}

// NewNeuron creates an initialized neuron
func NewNeuron(id string, inputSpacePotentialPoolPercent int, inputSpaceSize int) *Neuron {
	n := &Neuron{
		ID:          id,
		CoordLookup: map[int]int{},
		Coordinates: []int{},
		Permanences: []int{},
	}
	maxConnections := int(float32(inputSpaceSize) * (float32(inputSpacePotentialPoolPercent) / 100))
	position := 0
	inputSpaceRandom := NewUniqueRand(inputSpaceSize)
	for j := 0; j < inputSpaceSize && len(n.Coordinates) < maxConnections; j++ {
		if rand.Int()%100 < inputSpacePotentialPoolPercent {
			newCoord := inputSpaceRandom.Int()
			n.CoordLookup[newCoord] = position
			n.Coordinates = append(n.Coordinates, newCoord)
			n.Permanences = append(n.Permanences, rand.Int()%10)
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
	return n.Permanences[index]
}

// IncPermanence increases the permanance value for this connection
func (n *Neuron) IncPermanence(coord int) {
	index, _ := n.CoordLookup[coord]
	if n.Permanences[index] < 9 {
		n.Permanences[index]++
	}
}

// DecPermanence decrease the permanance value for this connection
func (n *Neuron) DecPermanence(coord int) {
	index, _ := n.CoordLookup[coord]
	if n.Permanences[index] > 0 {
		n.Permanences[index]--
	}
}
