package geo

import "testing"

func TestDistance(t *testing.T) {

	taliLon := 24.846823667701187
	taliLat := 60.21315136048605
	hermanniLat := 60.195099
	hermanniLon := 24.966789

	hiveLat := 60.180771
	hiveLon := 24.9560864

	dist := Distance(taliLat, taliLon, hermanniLat, hermanniLon)
	wanted := 1.0
	if dist != wanted {
		t.Errorf("got: %f, wanted: %f", dist, wanted)
	}
}
