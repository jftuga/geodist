/*
shared.go
-John Taylor

shared components between the distance algorithms

*/

package geodist

const version string = "1.0.0"

// Coord represents a geographic coordinate
type Coord struct {
	Lat float64
	Lon float64
}
