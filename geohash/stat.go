package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/tiborv/hilbert-geohash/geohash/hash"
	"github.com/tiborv/hilbert-geohash/geohash/point"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		n, err := strconv.Atoi(args[1])
		p := 1000 //Default amount of points if none provided
		if len(args) > 2 {
			p, err = strconv.Atoi(args[2])
		}
		if err != nil {
			fmt.Println("Invalid input")
			PrintHelp()
			return
		}
		switch args[0] {
		case "random":
			RunRandomPoints(n, p)
		case "maxdistancelat":
			RunMaxDistancepLat(n, p)
		case "maxdistancelng":
			RunMaxDistancepLng(n, p)
		case "all":
			RunRandomPoints(n, p)
			RunMaxDistancepLat(n, p)
			RunMaxDistancepLng(n, p)
		default:
			PrintHelp()
		}
		return
	}
	PrintHelp()
}

func PrintHelp() {
	fmt.Println("Invalid input.\nUse: stats.go |random|maxdistancelat|maxdistancelng|all <numberofiterations> <numberofpoints>")
}

func RunRandomPoints(n, p int) {
	fmt.Printf("\nRunning %v iterations, using %v RANDOM points...", n, p)
	resultsHilb := make([]float64, n)
	resultsZ := make([]float64, n)
	levenshteinDiff := make([]float64, n)
	for i := 0; i < n; i++ { //Repeat n times
		hashes := make([]hash.Hash, p)
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < p; i++ { // Use p amount of points
			randLat := rand.Float64() * point.MAX_LATITUDE
			randLng := rand.Float64() * point.MAX_LONGITUDE
			p := point.NewPoint(randLat, randLng, 0, 0)
			hash, _ := hash.NewHashPoint(p)
			hashes[i] = hash
		}

		resultsHilb[i], resultsZ[i] = compareHashes(hashes, hashes, true)
		levenshteinDiff[i] = compareLevenshtein(hashes, hashes, true)

	}
	PrintResults(resultsHilb, resultsZ, levenshteinDiff)
}

func RunMaxDistancepLng(n, p int) {
	fmt.Printf("\nRunning %v iterations, using %v MAX DISTANCE LNG points...", n, p)
	resultsHilb := make([]float64, n)
	resultsZ := make([]float64, n)
	levenshteinDiff := make([]float64, n)
	for i := 0; i < n; i++ { //Repeat n times
		h1 := make([]hash.Hash, p/2)
		h2 := make([]hash.Hash, p/2)

		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < p/2; i++ { // Create p/2 amount of points on one side
			randLat := rand.Float64() * point.MAX_LATITUDE
			randLng := 1.0 * point.MAX_LONGITUDE
			p := point.NewPoint(randLat, randLng, 0, 0)
			hash, _ := hash.NewHashPoint(p)
			h1[i] = hash
		}
		for i := 0; i < p/2; i++ { // Create p/2 amount of points on other side
			randLat := rand.Float64() * point.MAX_LATITUDE
			randLng := -1.0 * point.MAX_LONGITUDE
			p := point.NewPoint(randLat, randLng, 0, 0)
			hash, _ := hash.NewHashPoint(p)
			h2[i] = hash
		}

		resultsHilb[i], resultsZ[i] = compareHashes(h1, h2, false)
		levenshteinDiff[i] = compareLevenshtein(h1, h2, false)
	}
	PrintResults(resultsHilb, resultsZ, levenshteinDiff)

}

func RunMaxDistancepLat(n, p int) {
	fmt.Printf("\nRunning %v iterations, using %v MAX DISTANCE LAT points...", n, p)
	resultsHilb := make([]float64, n)
	resultsZ := make([]float64, n)
	levenshteinDiff := make([]float64, n)
	for i := 0; i < n; i++ { //Repeat n times
		h1 := make([]hash.Hash, p/2)
		h2 := make([]hash.Hash, p/2)

		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < p/2; i++ { // Create p/2 amount of points on one side
			randLat := 1.0 * point.MAX_LATITUDE // Max latitude
			randLng := rand.Float64() * point.MAX_LONGITUDE
			p := point.NewPoint(randLat, randLng, 0, 0)
			hash, _ := hash.NewHashPoint(p)
			h1[i] = hash
		}
		for i := 0; i < p/2; i++ { // Create p/2 amount of points on other side
			randLat := -1.0 * point.MAX_LATITUDE //Opposite side maximum
			randLng := rand.Float64() * point.MAX_LONGITUDE
			p := point.NewPoint(randLat, randLng, 0, 0)
			hash, _ := hash.NewHashPoint(p)
			h2[i] = hash
		}
		resultsHilb[i], resultsZ[i] = compareHashes(h1, h2, false)
		levenshteinDiff[i] = compareLevenshtein(h1, h2, false)
	}
	PrintResults(resultsHilb, resultsZ, levenshteinDiff)
}

func compareHashes(h1, h2 []hash.Hash, equalInput bool) (float64, float64) {
	hilbertIsShorter, zordertIsShorter, checked := 0, 0, float64(0)
	for i := 0; i < len(h1); i++ {
		for j := 0; j < len(h2); j++ {
			if equalInput && i == j {
				continue //Skip comparison of same point when h1 == h2
			}
			hilbertDistance := h1[i].DistanceHilbert(h2[j])
			zorderDistance := h1[i].DistanceZorder(h2[j])
			if zorderDistance > hilbertDistance {
				hilbertIsShorter++
			}
			if zorderDistance < hilbertDistance {
				zordertIsShorter++
			}
			checked++
		}
	}
	return float64(hilbertIsShorter) / checked,
		float64(zordertIsShorter) / checked
}

func PrintResults(hilbRes, zorderRes, levenshteinRes []float64) {
	hilbAvg := 100 * Sum(hilbRes) / float64(len(hilbRes))
	zorderAvg := 100 * Sum(zorderRes) / float64(len(zorderRes))
	equal := 100 - (hilbAvg + zorderAvg)
	levenshteinAvg := Sum(levenshteinRes) / float64(len(levenshteinRes))
	fmt.Printf("\nThe Hilbert curve had a shorter distance %.2f%% of the time ", hilbAvg)
	fmt.Printf("\nThe Z-order curve had a shorter distance %.2f%% of the time ", zorderAvg)
	fmt.Printf("\nThe curves had eual distance %.2f%% of the time ", equal)
	fmt.Printf("\nThe Hilbert curve was on average %.2f shorter in Levenshtein distance\n",
		levenshteinAvg)
}

func Sum(arr []float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	return
}

func compareLevenshtein(h1, h2 []hash.Hash, equalInput bool) float64 {
	levenshtein, checked := 0, 0
	for i := 0; i < len(h1); i++ {
		for j := 0; j < len(h2); j++ {
			if equalInput && i == j {
				continue //Skip comparison of same point when h1 == h2
			}
			levenshteinHilbert := DistanceLevenshtein(h1[i].String, h2[j].String)
			levenshteinZorder := DistanceLevenshtein(h1[i].GetZorder().HashString(32),
				h2[j].GetZorder().HashString(32))
			levenshtein += levenshteinZorder - levenshteinHilbert //Hilbert is closer if positive
			checked++
		}
	}
	return float64(levenshtein) / float64(checked)
}

//http://rosettacode.org/wiki/Levenshtein_distance#Go
func DistanceLevenshtein(s, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}
