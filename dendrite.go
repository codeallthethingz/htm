package main

// Dendrite is a connection to some axon somewhere.
type Dendrite struct {
	InputCoordinate int
	Permanence      int
}

// NewDendrite create a new dendrite
func NewDendrite(inputCoordinate int, permanence int) *Dendrite {
	return &Dendrite{
		InputCoordinate: inputCoordinate,
		Permanence:      permanence,
	}
}

// IncPermanence increases the permanance value for this connection
func (d *Dendrite) IncPermanence() {
	if d.Permanence < 9 {
		d.Permanence++
	}
}

// DecPermanence decrease the permanance value for this connection
func (d *Dendrite) DecPermanence() {
	if d.Permanence > 0 {
		d.Permanence--
	}
}
