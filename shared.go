/*
shared.go
-John Taylor

shared components between the distance algorithms

*/

package geodist

import "math"

const version string = "1.1.0"

const kilometersPerMile = 1.609344

// Coord represents a geographic coordinate in degrees. Latitude must be in
// [-90, 90]; longitude may be any finite value and is normalized modulo 360.
type Coord struct {
	Lat float64
	Lon float64
}

func validCoord(p Coord) bool {
	return p.Lat >= -90 && p.Lat <= 90 && !math.IsNaN(p.Lon) && !math.IsInf(p.Lon, 0)
}

func longitudeDifference(p1, p2 Coord) float64 {
	// Reduce before subtracting to avoid overflow for large finite longitudes.
	return math.Remainder(math.Remainder(p2.Lon, 360)-math.Remainder(p1.Lon, 360), 360)
}

func coincident(p1, p2 Coord, longitudeDifference float64) bool {
	return p1.Lat == p2.Lat && (longitudeDifference == 0 || math.Abs(p1.Lat) == 90)
}
