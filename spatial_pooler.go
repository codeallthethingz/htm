package main

import (
	"fmt"
	"math/rand"
)

// Cell is the connection from a spatial pooler cell to many input pixels
type Cell struct {
	Coordinates []int
	CoordLookup map[int]int
	Permanences []int

	Score  int
	ID     string
	Active bool
}

// SpatialPooler is a set of cells connecting to an input space
type SpatialPooler struct {
	Cells            []*Cell
	ActivatedCells   map[int]bool
	InputSpaceWidth  int
	InputSpaceHeight int
}

// Activate the cells in the spatial pooler for an encoded input
func (sp *SpatialPooler) Activate(encoded string, connectionThreshold int, overlap int, learning bool) {
	for i, cell := range sp.Cells {
		score := 0
		hits := ""
		for j, coord := range cell.Coordinates {
			if encoded[coord] == "X"[0] {
				if cell.Permanences[j] > connectionThreshold {
					hits += fmt.Sprintf("%d[0.%d] ", coord, cell.Permanences[j])
					score++
				}
			}
		}
		cell.Score = score
		if score > overlap {
			sp.ActivatedCells[i] = true
			cell.Active = true

			// learn
			if learning {
				for j, coord := range cell.Coordinates {
					if encoded[coord] == "X"[0] {
						if cell.Permanences[j] < 9 {
							cell.Permanences[j]++
						}
					} else if cell.Permanences[j] > 0 {
						cell.Permanences[j]--
					}
				}
			}
		}
	}
}

// NewSpatialPooler create a new pooler.
func NewSpatialPooler(spatialPoolerSize int, inputSpacePotentialPoolPercent int, inputSpaceWidth int, inputSpaceHeight int) *SpatialPooler {
	spatialPooler := &SpatialPooler{
		Cells:            make([]*Cell, spatialPoolerSize),
		ActivatedCells:   map[int]bool{},
		InputSpaceWidth:  inputSpaceWidth,
		InputSpaceHeight: inputSpaceHeight,
	}

	inputSpaceRandom := NewUniqueRand(inputSpaceWidth * inputSpaceHeight)
	for i := 0; i < len(spatialPooler.Cells); i++ {
		inputSpaceRandom.Reset()
		spatialPooler.Cells[i] = &Cell{
			ID:          fmt.Sprintf("c%d", i),
			CoordLookup: map[int]int{},
			Coordinates: []int{},
			Permanences: []int{},
		}
		position := 0
		for j := 0; j < inputSpaceWidth*inputSpaceHeight; j++ {
			if rand.Int()%100 < inputSpacePotentialPoolPercent {

				newCoord := inputSpaceRandom.Int()
				spatialPooler.Cells[i].CoordLookup[newCoord] = position
				spatialPooler.Cells[i].Coordinates = append(spatialPooler.Cells[i].Coordinates, newCoord)
				spatialPooler.Cells[i].Permanences = append(spatialPooler.Cells[i].Permanences, rand.Int()%10)
				position++
			}
		}
	}

	return spatialPooler
}
