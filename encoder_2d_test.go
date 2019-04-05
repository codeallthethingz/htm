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
	encoded := Encode(testCup, 19, 0.04)
	onBits, offBits := CountBits(testCup)
	onBitsEncoded, _ := CountBits(encoded)
	onBitsEncodedTarget := int(float32(onBits+offBits) * 0.04)
	require.True(t, (onBitsEncodedTarget*2) > onBitsEncoded)
}
func TestEncodeZero(t *testing.T) {
	width := 29
	encoded := Encode(testZero, width, 0.04)
	onBits, offBits := CountBits(testZero)
	onBitsEncoded, _ := CountBits(encoded)
	onBitsEncodedTarget := int(float32(onBits+offBits) * 0.04)
	require.True(t, (onBitsEncodedTarget*2) > onBitsEncoded)
}
