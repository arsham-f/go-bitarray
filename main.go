package bitarray

import (
	"fmt"
	"math"
	"runtime"
)

var (
	offset = map[uint8]uint8{
		0: 0,
		1: 1,
		2: 1 << 1,
		3: 1 << 2,
		4: 1 << 3,
		5: 1 << 4,
		6: 1 << 5,
		7: 1 << 6,
	}
)

type BitArray struct {
	bits []uint8
}

func New(size uint64) *BitArray {
	ba := &BitArray{}
	s := math.Ceil(float64(size) / float64(8))
	ba.bits = make([]uint8, int(s))
	return ba
}

/*
	Find offsets for n
		s: slice offset
		b: bit offset
*/
func FindOffsets(n uint64) (s_offset uint64, b_offset uint8) {
	s_offset = uint64(math.Floor(float64(n) / float64(8)))
	b_offset = uint8(n - (s_offset * 8))
	return
}

func (ba *BitArray) Add(n uint64) {
	s, b := FindOffsets(n)
	ba.bits[s] = ba.bits[s] | offset[b]

}

func (ba *BitArray) Remove(n uint64) {
	s, b := FindOffsets(n)
	ba.bits[s] = ba.bits[s] & ^offset[b]
}

func (ba BitArray) Contains(n uint64) bool {
	s, b := FindOffsets(n)
	bitfield := ba.bits[s]
	return (bitfield&offset[b] > 0)
}
