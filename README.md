# geodist
GoLang package to compute the distance between two geographic latitude, longitude coordinates

## Algorithm Comparison
* `Vincenty` is more accurate than `Haversine` because is considers the Earth's [ellipticity](https://www.dictionary.com/browse/ellipticity) when performing the calculation, but takes a longer time to compute.
* [Wikipedia: Haversine](https://en.wikipedia.org/wiki/Haversine_formula)
* [Wikipedia: Vincenty](https://en.wikipedia.org/wiki/Vincenty%27s_formulae)
* [Is the Haversine Formula or the Vincenty's Formula better for calculating distance?](https://stackoverflow.com/q/38248046/452281)

## Example
* [example.go](example/example.go)

```go
	var newYork = geodist.Coord{Lat: 40.7128, Lon: 74.0060}
	var sanDiego = geodist.Coord{Lat: 32.7157, Lon: 117.1611}
	miles, km, ok := geodist.VincentyDistance(newYork, sanDiego)
	if !ok {
		fmt.Println("Unable to compute Vincenty Distance.")
		return
	}
	fmt.Printf(" [Vincenty] New York to San Diego: %.3f m, %.3f km\n", miles, km)

	var elPaso = geodist.Coord{Lat: 31.7619, Lon: 106.4850}
	var stLouis = geodist.Coord{Lat: 38.6270, Lon: 90.1994}
	miles, km = geodist.HaversineDistance(elPaso, stLouis)
	fmt.Printf("[Haversine] El Paso to St. Louis:  %.3f m, %.3f km\n", miles, km)
```

## Online Calculators
* **Great Circle**: [NOAA: Latitude/Longitude Distance Calculator](https://www.nhc.noaa.gov/gccalc.shtml)
* **Haversine**: [Moveable Type: Calculate distance, bearing and more between Latitude/Longitude points](https://www.movable-type.co.uk/scripts/latlong.html)
* **Vincenty**: [CQSRG: WGS-84 World Geodetic System Distance Calculator](https://www.cqsrg.org/tools/GCDistance/)

## Acknowledgements
* [Haversine Algorithm](https://gist.github.com/cdipaolo/d3f8db3848278b49db68) in `GoLang`
* [Vincenty Algorithm](https://web.archive.org/web/20181109001358/http://www.5thandpenn.com/GeoMaps/GMapsExamples/distanceComplete2.html) in `JavaScript`
