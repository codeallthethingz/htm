package main

import (
	"math/rand"
)

// Cell is the connection from a spatial pooler cell to many input pixels
type Cell struct {
	Coordinates []int
	CoordLookup map[int]int
	Permenances []int
}

// SpatialPooler is a set of cells connecting to an input space
type SpatialPooler struct {
	Cells []Cell
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(spatialPoolerSize int, inputSpacePotentialPoolPercent int, inputSpaceSize int) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Cells: make([]Cell, spatialPoolerSize),
	}

	inputSpaceRandom := NewUniqueRand(inputSpaceSize)
	for i := 0; i < len(spatialPooler.Cells); i++ {
		inputSpaceRandom.Reset()
		spatialPooler.Cells[i] = Cell{
			CoordLookup: map[int]int{},
			Coordinates: []int{},
			Permenances: []int{},
		}
		position := 0
		for j := 0; j < inputSpaceSize; j++ {
			if rand.Int()%100 < inputSpacePotentialPoolPercent {

				newCoord := inputSpaceRandom.Int()
				spatialPooler.Cells[i].CoordLookup[newCoord] = position
				spatialPooler.Cells[i].Coordinates = append(spatialPooler.Cells[i].Coordinates, newCoord)
				spatialPooler.Cells[i].Permenances = append(spatialPooler.Cells[i].Permenances, rand.Int()%10)
				position++
			}
		}
	}

	return spatialPooler
}
