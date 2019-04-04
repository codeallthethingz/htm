package main

import (
	"fmt"
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

func TestEncode(t *testing.T) {
	encoded := Encode(testCup, 19, 0.04)
	onBits, offBits := CountBits(testCup)
	onBitsEncoded, _ := CountBits(encoded)

	onBitsEncodedTarget := int(float32(onBits+offBits) * 0.04)
	printEncoding(testCup, 19)
	fmt.Println(onBitsEncodedTarget, onBitsEncoded)
	printEncoding(encoded, 19)
	require.Equal(t, onBitsEncodedTarget, onBitsEncoded)
}
