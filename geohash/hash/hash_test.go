package hash

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tiborv/hilbert-geohash/geohash/point"
)

func TestHash(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(time.Now().UTC().UnixNano())
	n := 100000

	hashes := make([]Hash, n)

	for i := 0; i < n; i++ {
		lat := rand.Float64() * point.MAX_LATITUDE
		lng := rand.Float64() * point.MAX_LONGITUDE
		p := point.NewPoint(lat, lng, 0, 0)
		hash, err := NewHashPoint(p)
		hashes[i] = hash
		assert.Nil(err, "Err should be nil")
		assert.True(hash.GenPoint().WithinErr(p), "Hash result-point should be within point error")
	}

}

func SumArr(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func BenchmarkHash(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())
	latVals := make([]float64, b.N)
	lngVals := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		latVals[i] = rand.Float64() * 90
		lngVals[i] = rand.Float64() * 180
	}

	b.ResetTimer() // Start benchmark after setup
	for i := 0; i < b.N; i++ {
		p := point.NewPoint(latVals[i], lngVals[i], 0, 0)
		NewHashPoint(p)
	}

}
func BenchmarkPoint(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())
	hashes := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		hash := make([]byte, 13)
		for j := 0; j < 12; j++ {
			hash[j] = []byte(strconv.FormatInt(rand.Int63n(31), 32))[0] //Random hash
		}
		hashes[i] = hash
	}

	b.ResetTimer() // Start benchmark after setup
	for i := 0; i < b.N; i++ {
		h, _ := NewHashString(string(hashes[i]))
		h.GenPoint()
	}

}
