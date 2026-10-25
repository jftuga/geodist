/*

haversine.go
-John Taylor

Compute the distance between two geographic points when given a pair of latitude-longitude coordinates

Haversine formula:
https://en.wikipedia.org/wiki/Haversine_formula

The code below was adapted from Conner DiPaolo:
https://gist.github.com/cdipaolo/d3f8db3848278b49db68

*/

package geodist

import (
	"math"
)

// adapted from: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// HaversineDistance returns the great-circle distance in miles and kilometers,
// using a spherical approximation of the Earth with a radius of 6,378,100 meters.
// Coordinates are supplied in degrees. Invalid coordinates return NaN, NaN.
//
// http://en.wikipedia.org/wiki/Haversine_formula
func HaversineDistance(p1, p2 Coord) (float64, float64) {
	if !validCoord(p1) || !validCoord(p2) {
		return math.NaN(), math.NaN()
	}
	dLon := longitudeDifference(p1, p2)
	if coincident(p1, p2, dLon) {
		return 0, 0
	}
	// convert to radians
	// must cast radius as float to multiply later
	var la1, la2, r float64

	piRad := math.Pi / 180
	la1 = p1.Lat * piRad
	la2 = p2.Lat * piRad

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(dLon*piRad)
	// Roundoff near antipodes can put h slightly outside its mathematical range.
	h = math.Max(0, math.Min(1, h))

	meters := 2 * r * math.Asin(math.Sqrt(h))
	kilometers := meters / 1000
	miles := kilometers / kilometersPerMile
	return miles, kilometers
}
