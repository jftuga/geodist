package geodist

import (
	"math"
	"testing"
)

func checkDistance(t *testing.T, miles, km, wantKM, tolerance float64) {
	t.Helper()
	if math.IsNaN(km) || math.IsInf(km, 0) || math.Abs(km-wantKM) > tolerance {
		t.Errorf("kilometers = %.12f, want %.12f (tolerance %g)", km, wantKM, tolerance)
	}
	if math.IsNaN(miles) || math.IsInf(miles, 0) || math.Abs(miles*1.609344-km) > 1e-9 {
		t.Errorf("inconsistent units: %g miles, %g kilometers", miles, km)
	}
}

func TestHaversineAntipodes(t *testing.T) {
	for _, pair := range [][2]Coord{
		{{0, 0}, {0, 180}},
		{{-49.17091865018985, 0.9778610880352412}, {49.17091865018985, -179.02213891196476}},
		{{90, 0}, {-90, 0}},
	} {
		miles, km := HaversineDistance(pair[0], pair[1])
		checkDistance(t, miles, km, math.Pi*6378.1, 0.001)
	}
}

func TestVincentyReferenceDistances(t *testing.T) {
	// Reference distances from GeographicLib's WGS-84 test data:
	// https://github.com/geographiclib/geographiclib-python/blob/main/geographiclib/test/test_geodesic.py
	// A 1 mm tolerance accommodates Vincenty's series approximation.
	for _, tt := range []struct {
		name   string
		p1, p2 Coord
		meters float64
	}{
		{"intercontinental", Coord{35.60777, -139.44815}, Coord{-11.17491, -69.95921}, 8935244.5604818305},
		{"unrolled longitude", Coord{55.52454, 106.05087}, Coord{77.03196, 197.18234}, 4105086.1713924406},
		{"short distance", Coord{54.1589, 15.3872}, Coord{54.1591, 15.3877}, 39.527686385},
		// Computed with GeographicLib's Geodesic.WGS84.Inverse(0, 0, 0.5, 179.5).
		{"slow convergence", Coord{0, 0}, Coord{0.5, 179.5}, 19936288.578965314},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for _, pair := range [][2]Coord{{tt.p1, tt.p2}, {tt.p2, tt.p1}} {
				miles, km, err := VincentyDistance(pair[0], pair[1])
				if err != nil {
					t.Fatal(err)
				}
				checkDistance(t, miles, km, tt.meters/1000, 1e-6)
			}
		})
	}
}

func TestDateLine(t *testing.T) {
	// Changing the representation of the same meridian must not change distance.
	_, want, err := VincentyDistance(Coord{0, 0}, Coord{10, 0})
	if err != nil {
		t.Fatal(err)
	}
	for _, pair := range [][2]Coord{
		{{0, -180}, {10, 180}},
		{{10, 180}, {0, -180}},
		{{0, 180}, {10, -180}},
		{{10, -180}, {0, 180}},
	} {
		miles, km, err := VincentyDistance(pair[0], pair[1])
		if err != nil {
			t.Fatal(err)
		}
		checkDistance(t, miles, km, want, 1e-9)
	}
	// Two degrees along the equator across the date line.
	miles, km, err := VincentyDistance(Coord{0, 179}, Coord{0, -179})
	if err != nil {
		t.Fatal(err)
	}
	checkDistance(t, miles, km, 6378.137*2*math.Pi/180, 1e-6)
}

func TestEquivalentCoordinates(t *testing.T) {
	for _, pair := range [][2]Coord{
		{{10, -180}, {10, 180}},
		{{10, 0}, {10, 360}},
		{{90, 0}, {90, 120}},
		{{-90, -70}, {-90, 40}},
		{{40, math.MaxFloat64}, {40, math.MaxFloat64}},
	} {
		miles, km := HaversineDistance(pair[0], pair[1])
		checkDistance(t, miles, km, 0, 0)
		miles, km, err := VincentyDistance(pair[0], pair[1])
		if err != nil {
			t.Fatal(err)
		}
		checkDistance(t, miles, km, 0, 0)
	}
}

func TestInvalidCoordinates(t *testing.T) {
	for _, bad := range []Coord{
		{math.NaN(), 0}, {0, math.NaN()},
		{math.Inf(1), 0}, {math.Inf(-1), 0},
		{0, math.Inf(1)}, {0, math.Inf(-1)},
		{90.001, 0}, {-90.001, 0},
	} {
		for _, pair := range [][2]Coord{{bad, {}}, {{}, bad}, {bad, bad}} {
			miles, km := HaversineDistance(pair[0], pair[1])
			if !math.IsNaN(miles) || !math.IsNaN(km) {
				t.Errorf("HaversineDistance(%v) = %g, %g; want NaN, NaN", pair, miles, km)
			}
			miles, km, err := VincentyDistance(pair[0], pair[1])
			if err == nil || miles != -1 || km != -1 {
				t.Errorf("VincentyDistance(%v) = %g, %g, %v; want -1, -1, error", pair, miles, km, err)
			}
		}
	}
}

func TestVincentyNonConvergence(t *testing.T) {
	for _, pair := range [][2]Coord{
		{{0, 0}, {0, 180}},
		{{0, 0}, {0, 179.9}},
		{{27.2, 0}, {-27.1, 179.5}},
	} {
		miles, km, err := VincentyDistance(pair[0], pair[1])
		if err == nil || miles != -1 || km != -1 {
			t.Errorf("VincentyDistance(%v) = %g, %g, %v; want -1, -1, error", pair, miles, km, err)
		}
	}
}
