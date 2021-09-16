package main

import (
	"fmt"
	"github.com/jftuga/geodist"
)

func main() {
	var newYork = geodist.Coord{Lat: 40.7128, Lon: 74.0060}
	var sanDiego = geodist.Coord{Lat: 32.7157, Lon: 117.1611}
	miles, km, err := geodist.VincentyDistance(newYork, sanDiego)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" [Vincenty] New York to San Diego: %.3f m, %.3f km\n", miles, km)
	miles, km = geodist.HaversineDistance(newYork, sanDiego)
	fmt.Printf("[Haversine] New York to San Diego: %.3f m, %.3f km\n", miles, km)

	fmt.Println()

	var elPaso = geodist.Coord{Lat: 31.7619, Lon: 106.4850}
	var stLouis = geodist.Coord{Lat: 38.6270, Lon: 90.1994}
	miles, km, err = geodist.VincentyDistance(elPaso, stLouis)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" [Vincenty] El Paso to St. Louis: %.3f m, %.3f km\n", miles, km)
	miles, km = geodist.HaversineDistance(elPaso, stLouis)
	fmt.Printf("[Haversine] El Paso to St. Louis: %.3f m, %.3f km\n", miles, km)
}
