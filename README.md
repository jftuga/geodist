# geodist

A Go package for computing the distance between two geographic coordinates using
the Haversine or Vincenty formula.

## Installation

```sh
go get github.com/jftuga/geodist
```

## Algorithm Comparison

* `Vincenty` is more accurate than `Haversine` because it considers the Earth's [ellipticity](https://www.dictionary.com/browse/ellipticity) when performing the calculation, but takes a longer time to compute.
* [Wikipedia: Haversine](https://en.wikipedia.org/wiki/Haversine_formula)
* [Wikipedia: Vincenty](https://en.wikipedia.org/wiki/Vincenty%27s_formulae)
* [Is the Haversine Formula or the Vincenty's Formula better for calculating distance?](https://stackoverflow.com/q/38248046/452281)

Both functions return distances in miles and kilometers, in that order. Coordinates
are in degrees, with latitude in `[-90, 90]` and any finite longitude (normalized
modulo 360). West longitudes and south latitudes are negative.

`VincentyDistance` returns `-1, -1, error` for invalid coordinates or failure to
converge. Some nearly antipodal points do not converge. `HaversineDistance` returns
`NaN, NaN` for invalid coordinates; it uses a sphere with a radius of 6,378,100 meters.

## Example

The following complete program is also available in [example/example.go](example/example.go).

```go
package main

import (
	"fmt"

	"github.com/jftuga/geodist"
)

func main() {
	var newYork = geodist.Coord{Lat: 40.7128, Lon: -74.0060}
	var sanDiego = geodist.Coord{Lat: 32.7157, Lon: -117.1611}
	miles, km, err := geodist.VincentyDistance(newYork, sanDiego)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" [Vincenty] New York to San Diego: %.3f mi, %.3f km\n", miles, km)
	miles, km = geodist.HaversineDistance(newYork, sanDiego)
	fmt.Printf("[Haversine] New York to San Diego: %.3f mi, %.3f km\n", miles, km)

	fmt.Println()

	var elPaso = geodist.Coord{Lat: 31.7619, Lon: -106.4850}
	var stLouis = geodist.Coord{Lat: 38.6270, Lon: -90.1994}
	miles, km, err = geodist.VincentyDistance(elPaso, stLouis)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" [Vincenty] El Paso to St. Louis: %.3f mi, %.3f km\n", miles, km)
	miles, km = geodist.HaversineDistance(elPaso, stLouis)
	fmt.Printf("[Haversine] El Paso to St. Louis: %.3f mi, %.3f km\n", miles, km)
}
```

Run it from the repository root:

```sh
go run ./example
```

Output:

```text
 [Vincenty] New York to San Diego: 2432.880 mi, 3915.341 km
[Haversine] New York to San Diego: 2430.504 mi, 3911.517 km

 [Vincenty] El Paso to St. Louis: 1033.858 mi, 1663.833 km
[Haversine] El Paso to St. Louis: 1033.678 mi, 1663.543 km
```

## Online Calculators

* **Great Circle**: [NOAA: Latitude/Longitude Distance Calculator](https://www.nhc.noaa.gov/gccalc.shtml)
* **Haversine**: [Movable Type: Calculate distance, bearing and more between Latitude/Longitude points](https://www.movable-type.co.uk/scripts/latlong.html)
* **Vincenty**: [Mike Turnbull (formerly CQSRG): WGS-84 World Geodetic System Distance Calculator](https://www.cqsrg.org/tools/GCDistance/)

When comparing results, check the calculator's units and Earth model. Spherical
calculators may use a different Earth radius from this package.

## Acknowledgements

* [Haversine implementation in Go](https://gist.github.com/cdipaolo/d3f8db3848278b49db68) by Conner DiPaolo.
* [Vincenty implementation in JavaScript](https://www.movable-type.co.uk/scripts/latlong-vincenty.html) by Chris Veness; [archived version used for the original port](https://web.archive.org/web/20181109001358/http://www.5thandpenn.com/GeoMaps/GMapsExamples/distanceComplete2.html).
