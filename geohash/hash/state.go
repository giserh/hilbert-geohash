package hash

type HilbertMapper struct {
	v uint64 //Value
	s int    //Next HilbertMapper
}

var map0 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 0, s: 0},
	1: HilbertMapper{v: 1, s: 1},
	2: HilbertMapper{v: 3, s: 2},
	3: HilbertMapper{v: 2, s: 1},
}

var map1 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 0, s: 1},
	1: HilbertMapper{v: 3, s: 3},
	2: HilbertMapper{v: 1, s: 0},
	3: HilbertMapper{v: 2, s: 0},
}

var map2 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 2, s: 2},
	1: HilbertMapper{v: 1, s: 2},
	2: HilbertMapper{v: 3, s: 1},
	3: HilbertMapper{v: 0, s: 3},
}

var map3 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 2, s: 3},
	1: HilbertMapper{v: 3, s: 0},
	2: HilbertMapper{v: 1, s: 3},
	3: HilbertMapper{v: 0, s: 2},
}

func getHilbertMap(zvalue uint64, currentState int) (uint64, int) {
	switch currentState {
	case 0:
		return map0[zvalue].v, map0[zvalue].s
	case 1:
		return map1[zvalue].v, map1[zvalue].s
	case 2:
		return map2[zvalue].v, map2[zvalue].s
	case 3:
		return map3[zvalue].v, map3[zvalue].s
	default:
		panic("Unkown HilbertMapper")
	}
}

func getHilbertInverseMap(hilbertvalue uint64, currentState int) (uint64, int) {
	switch currentState {
	case 0:
		return inverse0[hilbertvalue].v, inverse0[hilbertvalue].s
	case 1:
		return inverse1[hilbertvalue].v, inverse1[hilbertvalue].s
	case 2:
		return inverse2[hilbertvalue].v, inverse2[hilbertvalue].s
	case 3:
		return inverse3[hilbertvalue].v, inverse3[hilbertvalue].s
	default:
		panic("Unkown HilbertMapper")
	}
}

var inverse0 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 0, s: 0},
	1: HilbertMapper{v: 1, s: 1},
	2: HilbertMapper{v: 3, s: 1},
	3: HilbertMapper{v: 2, s: 2},
}

var inverse1 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 0, s: 1},
	1: HilbertMapper{v: 2, s: 0},
	2: HilbertMapper{v: 3, s: 0},
	3: HilbertMapper{v: 1, s: 3},
}

var inverse2 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 3, s: 3},
	1: HilbertMapper{v: 1, s: 2},
	2: HilbertMapper{v: 0, s: 2},
	3: HilbertMapper{v: 2, s: 1},
}

var inverse3 = map[uint64]HilbertMapper{
	0: HilbertMapper{v: 3, s: 2},
	1: HilbertMapper{v: 2, s: 3},
	2: HilbertMapper{v: 0, s: 3},
	3: HilbertMapper{v: 1, s: 0},
}
