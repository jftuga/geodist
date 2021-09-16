package geodist

import (
	"testing"
)

func TestHaversineDistance(t *testing.T) {
	// distance NY-SD: 3915340.577 m
	// distance SD-EP: 1011300.217 m
	// distance EP-SL: 1663833.491 m
	// distance SL-NY: 1406519.972 m
	var newYork = Coord{40.7128, 74.0060}
	var sanDiego = Coord{32.7157, 117.1611}
	var elPaso = Coord{31.7619, 106.4850}
	var stLouis = Coord{38.6270, 90.1994}

	var miles, km float64

	miles, km = HaversineDistance(newYork, sanDiego)
	if int(miles) != 2430 || int64(km) != 3911 {
		t.Errorf("Computed values: %v %10f\n", miles, km)
		t.Errorf("Incorrect computation between New York and San Diego: %v %v\n", int(miles), int64(km))
	}

	miles, km = HaversineDistance(sanDiego, elPaso)
	if int(miles) != 627 || int64(km) != 1010 {
		t.Errorf("Computed values: %v %10f\n", miles, km)
		t.Errorf("Incorrect computation between San Diego and El Paso: %v %v\n", int(miles), int64(km))
	}

	miles, km = HaversineDistance(elPaso, stLouis)
	if int(miles) != 1033 || int(km) != 1663 {
		t.Errorf("Computed values: %v %10f\n", miles, km)
		t.Errorf("Incorrect computation between El Paso and St. Louis: %v %v\n", int(miles), int64(km))
	}

	miles, km = HaversineDistance(stLouis, newYork)
	if int(miles) != 872 || int(km) != 1404 {
		t.Errorf("Computed values: %v %10f\n", miles, km)
		t.Errorf("Incorrect computation between St. Louis and New York: %v %v\n", int(miles), int64(km))
	}

	miles, km = HaversineDistance(newYork, newYork)
	if int(miles) != 0 || int(km) != 0 {
		t.Errorf("Computed values: %v %10f\n", miles, km)
		t.Errorf("Incorrect computation between New York and New York: %v %v\n", int(miles), int64(km))
	}
}
