package common

/*
Common functions and types for the following -

* bit manipulation
* address
* register types and grouped register types
*/

const (
	ClkFrequency = 4120000 // ~4.12 Mhz
)
// Cycles represents the number of "machine cycles"
type Cycles uint8

/* Bit operations */
func SetBitAtIndex(r, i byte) byte {
	var mask byte = 1 << i
	return r | mask
}
func ResetBitAtIndex(r, i byte) byte {
	var mask byte = 1 << i
	return r & ^mask
}