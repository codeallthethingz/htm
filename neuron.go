package main

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	Coordinates []int
	CoordLookup map[int]int
	Permanences []int
	Score       int
	ID          string
	Active      bool
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
