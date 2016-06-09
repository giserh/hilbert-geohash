package hash

import (
	"strconv"
	"unicode/utf8"

	"github.com/tiborv/hilbert-geohash/geohash/bitarray"
	"github.com/tiborv/hilbert-geohash/geohash/error"
	"github.com/tiborv/hilbert-geohash/geohash/point"
)

type Hash struct {
	String  string `form:"hash" json:"hash" validate:"exists"`
	zorder  ba.BitArray
	hilbert ba.BitArray
}

func NewHashPoint(p point.Point) (Hash, error) {
	if !p.IsValid() {
		return Hash{}, errors.New("Invalid Point")
	}
	zorder := genZorder(p)
	hilbert := genHilbert(zorder)
	return Hash{hilbert: hilbert,
			zorder: zorder,
			String: hilbert.HashString(32)},
		nil
}

func NewHashString(hash string) (Hash, error) {
	h := Hash{String: hash}
	err := h.InitFromString()
	return h, err
}

func (h *Hash) InitFromString() error {
	val, err := strconv.ParseUint(h.String, 32, 64)
	bitArraylen := uint64(0)
	if h.String != "" {
		if err != nil {
			return err
		}
		//Each base32 value adds 5 bits
		bitArraylen = uint64(utf8.RuneCountInString(h.String)*5 - 1)
		if bitArraylen > 64 {
			return errors.New("Hash too long")
		}
	}
	h.hilbert = ba.NewBitArray(val, bitArraylen)
	h.zorder = genHilbertInverse(h.hilbert)
	return nil
}

func (h Hash) GetZorder() ba.BitArray {
	return h.zorder
}

func (h Hash) GetHilbert() ba.BitArray {
	return h.hilbert
}

func (h Hash) Equal(other Hash) bool {
	return h.hilbert.Equal(other.hilbert) && h.zorder.Equal(other.zorder)
}

func genZorder(p point.Point) ba.BitArray {
	zorder := ba.NewBitArray(0, 0)
	latdevide, lngdevide := float64(0), float64(0)
	lat, lng := float64(point.MAX_LATITUDE), float64(point.MAX_LONGITUDE)
	for i := 0; i < 32; i++ {
		lat /= 2
		lng /= 2
		switch setZorderQuadrant(&zorder, p, latdevide, lngdevide) {
		case 3:
			latdevide += lat
			lngdevide += lng
		case 2:
			latdevide -= lat
			lngdevide += lng
		case 1:
			latdevide += lat
			lngdevide -= lng
		case 0:
			latdevide -= lat
			lngdevide -= lng
		}
	}
	return zorder
}

func setZorderQuadrant(zorder *ba.BitArray,
	p point.Point, latdevide, lngdevide float64) uint64 {
	if p.Lng > lngdevide {
		if p.Lat > latdevide {
			return zorder.AppendPair(3)
		}
		return zorder.AppendPair(2)
	}
	if p.Lat > latdevide {
		return zorder.AppendPair(1)
	}
	return zorder.AppendPair(0)

}

func genHilbert(zorder ba.BitArray) ba.BitArray {
	hilbertValue := uint64(0)
	nextState := 0
	ba := ba.NewBitArray(0, 0)
	for i := uint64(0); i < zorder.Len(); i += 2 {
		hilbertValue, nextState = getHilbertMap(zorder.GetPair(i), nextState)
		ba.AppendPair(hilbertValue)
	}
	return ba
}

func genHilbertInverse(hilbert ba.BitArray) ba.BitArray {
	zorder := ba.NewBitArray(0, 0)
	nextState := 0
	zorderValue := uint64(0)
	for i := uint64(0); i < hilbert.Len(); i += 2 {
		zorderValue, nextState = getHilbertInverseMap(hilbert.GetPair(i), nextState)
		zorder.AppendPair(zorderValue)
	}
	return zorder
}

func (h Hash) GenPoint() point.Point {
	reslng, reslat := float64(0), float64(0)
	lat, lng := float64(point.MAX_LATITUDE), float64(point.MAX_LONGITUDE)
	errLat, errLng := float64(point.MAX_LATITUDE), float64(point.MAX_LONGITUDE)
	for i := uint64(0); i < h.zorder.Len(); i += 2 {
		lng /= 2
		lat /= 2
		errLat /= 2
		errLng /= 2
		switch h.zorder.GetPair(i) {
		case 3:
			reslat += lat
			reslng += lng
		case 2:
			reslat -= lat
			reslng += lng
		case 1:
			reslat += lat
			reslng -= lng
		case 0:
			reslat -= lat
			reslng -= lng
		}
	}
	return point.NewPoint(reslat, reslng, errLat, errLng)
}

func (h Hash) DistanceZorder(to Hash) uint64 {
	return h.zorder.Diff(to.zorder)
}

func (h Hash) DistanceHilbert(to Hash) uint64 {
	return h.hilbert.Diff(to.hilbert)
}
