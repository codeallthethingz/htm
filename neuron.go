package main

import (
	"math/rand"
)

// Neuron is the connection from a spatial pooler neuron to many inputs
type Neuron struct {
	columnFamily        string
	proximalInputLookup map[string]*Dendrite
	ProximalInputs      []*Dendrite `json:"proximalInputs"`
	MiniColumnNeurons   []*Neuron   `json:"miniColumnNeurons"`
	DistalInputs        []*Dendrite `json:"distalInputs"`
	Score               int         `json:"score"`
	ID                  string      `json:"id"`
	Active              bool        `json:"active"`
	Predictive          bool        `json:"predictive"`
	PreviouslyActive    bool        `json:"previouslyActive"`
}

// NewNeuron creates an initialized neuron
func NewNeuron(id string, potentialPoolPercent float64, inputNeurons []*Neuron, miniColumnNeurons []*Neuron) *Neuron {
	if id == "" {
		panic("must not create a neuron with empty ID")
	}
	connectionPoolSize := int(float64(len(inputNeurons)) * potentialPoolPercent)
	n := &Neuron{
		ID:                  id,
		columnFamily:        id,
		proximalInputLookup: map[string]*Dendrite{},
		ProximalInputs:      make([]*Dendrite, connectionPoolSize),
		MiniColumnNeurons:   miniColumnNeurons,
	}
	for _, miniColumnNeuron := range miniColumnNeurons {
		miniColumnNeuron.columnFamily = id
	}
	if inputNeurons != nil && len(inputNeurons) > 1 {
		inputNeuronsRandom := NewUniqueRand(len(inputNeurons))
		for j := range n.ProximalInputs {
			inputNeuron := inputNeurons[inputNeuronsRandom.Int()]
			permanence := rand.Int() % 10
			n.ProximalInputs[j] = NewDendrite(inputNeuron, permanence)
			if inputNeuron.ID == "" {
				panic("must not create a neuron with no ID")
			}
			n.proximalInputLookup[inputNeuron.ID] = n.ProximalInputs[j]
		}
	}

	return n
}

// ConnectDistal creates all the context connections
func (n *Neuron) ConnectDistal(allNeurons []*Neuron, potentialPoolPercent float64, temporalPoolingSize int) {
	connectionPoolSize := int(float64(((len(allNeurons)/(temporalPoolingSize+1))-1)*(temporalPoolingSize+1)) * potentialPoolPercent)
	n.DistalInputs = make([]*Dendrite, connectionPoolSize)
	allNeuronsRandom := NewUniqueRand(len(allNeurons))
	for i := 0; i < len(n.DistalInputs); i++ {
		contextNeuron := allNeurons[allNeuronsRandom.Int()]
		for contextNeuron.columnFamily == n.columnFamily {
			contextNeuron = allNeurons[allNeuronsRandom.Int()]
		}
		permanence := rand.Int() % 10
		n.DistalInputs[i] = NewDendrite(contextNeuron, permanence)
	}
}

// IsConnected is this neuron connected to the input
func (n *Neuron) IsConnected(inputNeuron *Neuron) bool {
	_, ok := n.proximalInputLookup[inputNeuron.ID]
	return ok
}

// GetDendrite get a dendrite connected to this coordinate
func (n *Neuron) GetDendrite(inputNeuron *Neuron) *Dendrite {
	dendrite, _ := n.proximalInputLookup[inputNeuron.ID]
	return dendrite
}

// GetActive get all the active neurons in this column
func (n *Neuron) GetActive() []*Neuron {
	active := []*Neuron{}
	if n.Active {
		active = append(active, n)
	}
	for _, miniNeuron := range n.MiniColumnNeurons {
		if miniNeuron.Active {
			active = append(active, miniNeuron)
		}
	}
	return active
}

// AllActive is this neuron and all the mini column neurons active
func (n *Neuron) AllActive() bool {
	return len(n.MiniColumnNeurons)+1 == len(n.GetActive())
}

// SomeActive at least one but not all the neurons are active in this column
func (n *Neuron) SomeActive() bool {
	active := len(n.GetActive())
	return active > 0 && active < len(n.MiniColumnNeurons)+1
}

// GetPredictive returns a list of all the predictive neurons.
func (n *Neuron) GetPredictive() []*Neuron {
	predictive := []*Neuron{}
	if n.Predictive {
		predictive = append(predictive, n)
	}
	for _, miniNeuron := range n.MiniColumnNeurons {
		if miniNeuron.Predictive {
			predictive = append(predictive, miniNeuron)
		}
	}
	return predictive
}
