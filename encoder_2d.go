package main

// Encode turn a 2d image into a a set of turned on bits at a specific sparsity
func Encode(obj string, inputNeurons []*Neuron, sparsity float32, width int) {
	onBits, offBits := CountBits(obj)
	totalBits := offBits + onBits
	target := int(float32(totalBits) * sparsity)
	encoded := turnOffBitsLinear(obj, width, onBits, target)

	for c := range encoded {
		inputNeurons[c].Active = encoded[c] == "X"[0]
	}
}

func turnOffBitsLinear(obj string, width int, currentlyOn int, targetOn int) string {
	newObj := ""
	for c := range obj {
		if currentlyOn > targetOn && c > 0 && obj[c-1] == "X"[0] && obj[c] == "X"[0] && c < len(obj) && obj[c+1] == "X"[0] {
			newObj += " "
			currentlyOn--
		} else {
			newObj += string(obj[c])
		}
	}
	superNewObj := ""
	for c := range newObj {
		if currentlyOn > targetOn && newObj[c] == "X"[0] && c < len(newObj) && newObj[c+1] == "X"[0] {
			superNewObj += " "
			currentlyOn--
		} else {
			superNewObj += string(newObj[c])
		}

	}
	newObj = ""
	for c := range superNewObj {
		if currentlyOn > targetOn && c > width && newObj[c-width] == "X"[0] {
			newObj += " "
			currentlyOn--
		} else {
			newObj += string(superNewObj[c])
		}
	}
	return newObj
}

// CountBits count the number of on and off bits
func CountBits(obj string) (int, int) {
	onBits := 0
	offBits := 0
	for c := range obj {
		if obj[c] == "X"[0] {
			onBits++
		} else if obj[c] == " "[0] {
			offBits++
		}
	}
	return onBits, offBits
}

// MakeInputNeurons create an empty list of neurons for use as an input field.
func MakeInputNeurons(len int) []*Neuron {
	neurons := make([]*Neuron, len)
	for i := 0; i < len; i++ {
		neurons[i] = &Neuron{}
	}
	return neurons
}

// InputNeuronsToString convert to a string
func InputNeuronsToString(neurons []*Neuron) string {
	result := ""
	for _, n := range neurons {
		if n.Active {
			result += "X"
		} else {
			result += " "
		}
	}
	return result
}
