package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testCup = "" +
	"                   " +
	"                   " +
	"                   " +
	"   XXXXXXXXXX XX   " +
	"   XXXXXXXXXXX  X  " +
	"   XXXXXXXXXXX  X  " +
	"   XXXXXXXXXX XX   " +
	"    XXXXXXXX       " +
	"                   " +
	"                   " +
	"                   "

var testZero = "" +
	"         XXXXXX              " +
	"         XXXXXXX             " +
	"        XXXXXXXXX            " +
	"        XXXXXXXXXX           " +
	"        XXXXXXXXXX           " +
	"       XXXXX  XXXX           " +
	"       XXXXXX   XXX          " +
	"       XXXXXX    XXX         " +
	"       XXXXX     XXXX        " +
	"       XXXX      XXXX        " +
	"       XXXX       XXX        " +
	"       XXXX       XXX        " +
	"       XXXX       XXX        " +
	"       XXXX       XXX        " +
	"       XXXX      XXXX        " +
	"       XXXXX   XXXXX         " +
	"        XXXX XXXXXXX         " +
	"         XXXXXXXXXX          " +
	"         XXXXXXXXX           " +
	"           XXXXXXX           "

func TestEncodeCup(t *testing.T) {
	neurons := MakeInputNeurons(19, 11)
	Encode(testCup, neurons, 0.04, 19)
	onBits, offBits := CountBits(testCup)
	onBitsEncoded := count(neurons)
	onBitsEncodedTarget := int(float64(onBits+offBits) * 0.04)
	require.True(t, (onBitsEncodedTarget*2) > onBitsEncoded)
}

func TestEncodeZero(t *testing.T) {
	width := 29
	height := 20

	neurons := MakeInputNeurons(width, height)
	Encode(testZero, neurons, 0.04, width)
	onBits, offBits := CountBits(testZero)
	onBitsEncoded := count(neurons)
	onBitsEncodedTarget := int(float64(onBits+offBits) * 0.04)
	require.True(t, (onBitsEncodedTarget*2) > onBitsEncoded)
}

func count(neurons []*Neuron) int {
	onBitsEncoded := 0
	for _, e := range neurons {
		if e.Active {
			onBitsEncoded++
		}
	}
	return onBitsEncoded
}
