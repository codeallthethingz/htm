package main

import (
	"fmt"
	"math/rand"
)

// SpatialPooler is a set of neurons connecting to an input space
type SpatialPooler struct {
	Neurons          []*Neuron
	ActivatedNeurons map[int]bool
}

// Activate the neurons in the spatial pooler for an encoded input
func (sp *SpatialPooler) Activate(encoded string, connectionThreshold int, overlap int, learning bool) {
	for i, neuron := range sp.Neurons {
		score := 0
		for _, coord := range neuron.Coordinates {
			if encoded[coord] == "X"[0] {
				if neuron.GetPermanence(coord) >= connectionThreshold {
					score++
				}
			}
		}
		neuron.Score = score
		if score >= overlap {
			sp.ActivatedNeurons[i] = true
			neuron.Active = true

			// learn
			if learning {
				for _, coord := range neuron.Coordinates {
					if encoded[coord] == "X"[0] {
						neuron.IncPermanence(coord)
					} else {
						neuron.DecPermanence(coord)
					}
				}
			}
		}
	}
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(spatialPoolerSize int, inputSpacePotentialPoolPercent int, inputSpaceSize int) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Neurons:          make([]*Neuron, spatialPoolerSize),
		ActivatedNeurons: map[int]bool{},
	}
	maxConnections := int(float32(spatialPoolerSize) * (float32(inputSpacePotentialPoolPercent) / 100))
	inputSpaceRandom := NewUniqueRand(inputSpaceSize)
	for i := 0; i < len(spatialPooler.Neurons); i++ {
		inputSpaceRandom.Reset()
		spatialPooler.Neurons[i] = &Neuron{
			ID:          fmt.Sprintf("c%d", i),
			CoordLookup: map[int]int{},
			Coordinates: []int{},
			Permanences: []int{},
		}
		position := 0
		for j := 0; j < inputSpaceSize && len(spatialPooler.Neurons[i].Coordinates) < maxConnections; j++ {
			if rand.Int()%100 < inputSpacePotentialPoolPercent {
				newCoord := inputSpaceRandom.Int()
				spatialPooler.Neurons[i].CoordLookup[newCoord] = position
				spatialPooler.Neurons[i].Coordinates = append(spatialPooler.Neurons[i].Coordinates, newCoord)
				spatialPooler.Neurons[i].Permanences = append(spatialPooler.Neurons[i].Permanences, rand.Int()%10)
				position++
			}
		}
	}

	return spatialPooler
}

// Print to the command line
func (sp *SpatialPooler) Print(width int, height int) {
	for i := 0; i < len(sp.Neurons); i++ {
		fmt.Printf("neuron: %d", i)
		for c := 0; c < width*height; c++ {
			if c%width == 0 {
				fmt.Print("\n")
			}
			if sp.Neurons[i].IsConnected(c) {
				fmt.Print(sp.Neurons[i].Permanences[sp.Neurons[i].CoordLookup[c]])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
