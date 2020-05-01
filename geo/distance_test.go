package geo

import "testing"

func TestDistance(t *testing.T) {

	taliLon := 24.846823667701187
	taliLat := 60.21315136048605
	hermanniLat := 60.195099
	hermanniLon := 24.966789

	dist := Distance(taliLat, taliLon, hermanniLat, hermanniLon)
	if dist >= 7000 || dist <= 6900 {
		t.Errorf("got: %f, wanted: %s", dist, "6900-7000m")
	}
}
