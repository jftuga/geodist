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

// HaversineDistance returns the distance (in miles) between two points of
//	 a given longitude and latitude relatively accurately (using a spherical
//	 approximation of the Earth) through the Haversin Distance Formula for
//	 great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// http://en.wikipedia.org/wiki/Haversine_formula
func HaversineDistance(p1, p2 Coord) (float64, float64) {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64

	piRad := math.Pi / 180
	la1 = p1.Lat * piRad
	lo1 = p1.Lon * piRad
	la2 = p2.Lat * piRad
	lo2 = p2.Lon * piRad

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	meters := 2 * r * math.Asin(math.Sqrt(h))
	kilometers := meters / 1000
	miles := kilometers * 0.621371
	return miles, kilometers
}
