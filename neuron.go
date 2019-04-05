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
