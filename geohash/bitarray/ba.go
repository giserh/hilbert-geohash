package ba

import (
	"strconv"
	"strings"
)

type BitArray struct {
	arr uint64
	len uint64
}

func NewBitArray(val uint64, len uint64) BitArray {
	return BitArray{arr: val, len: len}
}

func (ba *BitArray) SetB(b, pos uint64) uint64 {
	ba.arr = ba.arr | (b << pos)
	return b
}

func (ba *BitArray) AppendPair(b uint64) uint64 {
	ba.SetB(b, ba.len)
	ba.len += 2
	return b
}

func (ba BitArray) HashString(base int) string {
	return strconv.FormatUint(ba.arr, base)
}

func (ba BitArray) GetPair(pos uint64) uint64 {
	return (ba.arr >> pos) & 3
}

func (ba BitArray) Len() uint64 {
	return ba.len
}

func (ba BitArray) GetArray() uint64 {
	return ba.arr
}

func (ba BitArray) String() string {
	s := strconv.FormatUint(ba.arr, 2)
	return strings.Repeat("0", int(ba.len)-len(s)) + s
}

func (ba BitArray) Equal(other BitArray) bool {
	return ba.arr == other.arr
}

func (ba BitArray) Diff(other BitArray) uint64 {
	if ba.arr > other.arr {
		return ba.arr - other.arr
	}
	return other.arr - ba.arr
}
