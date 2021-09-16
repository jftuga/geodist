/*

vincenty.go
-John Taylor

Compute the distance between two geographic points when given a pair of latitude-longitude coordinates

Vincenty formula:
https://en.wikipedia.org/wiki/Vincenty%27s_formulae

The code below was ported from Chris Veness's JavaScript version:
https://web.archive.org/web/20181109001358/http://www.5thandpenn.com/GeoMaps/GMapsExamples/distanceComplete2.html

*/

package geodist

import (
	"errors"
	"math"
)

// these constants are used for vincentyDistance()
const a = 6378137
const b = 6356752.3142
const f = 1 / 298.257223563 // WGS-84 ellipsiod

/*
VincentyDistance computes the distances between two georgaphic coordinates

Args:
	p1: the 'starting' point, given in latitude, longitude as a Coord struct

	p2: the 'ending' point

Returns:
	A 3 element tuple: distance between the 2 points given in (1) miles and (2) kilometers
	The 3rd element will return true upon a successful computation or
	false if the algorithm fails to converge. -1, -1, false is returned upon failure
*/
func VincentyDistance(p1, p2 Coord) (float64, float64, error) {
	// convert from degrees to radians
	piRad := math.Pi / 180
	p1.Lat = p1.Lat * piRad
	p1.Lon = p1.Lon * piRad
	p2.Lat = p2.Lat * piRad
	p2.Lon = p2.Lon * piRad

	L := p2.Lon - p1.Lon

	U1 := math.Atan((1 - f) * math.Tan(p1.Lat))
	U2 := math.Atan((1 - f) * math.Tan(p2.Lat))

	sinU1 := math.Sin(U1)
	cosU1 := math.Cos(U1)
	sinU2 := math.Sin(U2)
	cosU2 := math.Cos(U2)

	lambda := L
	lambdaP := 2 * math.Pi
	iterLimit := 20

	var sinLambda, cosLambda, sinSigma float64
	var cosSigma, sigma, sinAlpha, cosSqAlpha, cos2SigmaM, C float64

	for {
		if math.Abs(lambda-lambdaP) > 1e-12 && (iterLimit > 0) {
			iterLimit -= 1
		} else {
			break
		}
		sinLambda = math.Sin(lambda)
		cosLambda = math.Cos(lambda)

		sinSigma = math.Sqrt((cosU2*sinLambda)*(cosU2*sinLambda) + (cosU1*sinU2-sinU1*cosU2*cosLambda)*(cosU1*sinU2-sinU1*cosU2*cosLambda))
		if sinSigma == 0 {
			return 0, 0, nil // co-incident points
		}

		cosSigma = sinU1*sinU2 + cosU1*cosU2*cosLambda
		sigma = math.Atan2(sinSigma, cosSigma)
		sinAlpha = cosU1 * cosU2 * sinLambda / sinSigma
		cosSqAlpha = 1 - sinAlpha*sinAlpha
		cos2SigmaM = cosSigma - 2*sinU1*sinU2/cosSqAlpha
		if math.IsNaN(cos2SigmaM) {
			cos2SigmaM = 0 // equatorial line: cosSqAlpha=0
		}

		C = f / 16 * cosSqAlpha * (4 + f*(4-3*cosSqAlpha))
		lambdaP = lambda
		lambda = L + (1-C)*f*sinAlpha*(sigma+C*sinSigma*(cos2SigmaM+C*cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)))
	}
	if iterLimit == 0 {
		return -1, -1, errors.New("vincenty algorithm failed to converge") // formula failed to converge
	}

	uSq := cosSqAlpha * (a*a - b*b) / (b * b)
	A := 1 + uSq/16384*(4096+uSq*(-768+uSq*(320-175*uSq)))
	B := uSq / 1024 * (256 + uSq*(-128+uSq*(74-47*uSq)))
	deltaSigma := B * sinSigma * (cos2SigmaM + B/4*(cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)-B/6*cos2SigmaM*(-3+4*sinSigma*sinSigma)*(-3+4*cos2SigmaM*cos2SigmaM)))
	meters := b * A * (sigma - deltaSigma)
	kilometers := meters / 1000
	miles := kilometers * 0.621371
	return miles, kilometers, nil
}
