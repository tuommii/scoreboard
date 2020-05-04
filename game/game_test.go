package game

import (
	"testing"
)

func TestCreateID(t *testing.T) {
	id := createID([]string{"TigerKing", "Pesukarhu", "Jian Yang", "Bubba"}, 1)
	wanted := "bjpt1"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}

	id = createID([]string{"Aapo", "Ari", "Jiang Yang"}, 24)
	wanted = "aaj24"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}

	id = createID([]string{"123xxx123"}, 20)
	wanted = "120"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}
}

func TestCreateCourse(t *testing.T) {
	course := createCourse([]string{"TigerKing", "Pesukarhu", "Jian Yang"}, 18, 1)

	wanted := 1
	if course.Active != wanted {
		t.Errorf("got: %d, wanted: %d", course.Active, wanted)
	}

	wanted = 18
	if len(course.Baskets) != wanted {
		t.Errorf("got: %d, wanted: %d", len(course.Baskets), wanted)
	}

	wanted = 3
	if len(course.Baskets[1].Scores) != wanted {
		t.Errorf("got: %d, wanted: %d", len(course.Baskets[1].Scores), wanted)
	}

	wanted = 3
	if course.Baskets[10].Par != wanted {
		t.Errorf("got: %d, wanted: %d", course.Baskets[10].Par, wanted)
	}
}

func TestCreate(t *testing.T) {
	// var templates []CourseInfo
	templates := LoadCourseTemplates("../")

	wanted := 3
	if len(templates) != wanted {
		t.Errorf("got: %d, wanted: %d", len(templates), wanted)
	}

	players := []string{"Miikka", "Pasi"}
	course := create(templates, 0, 0, players, 18, 1)

	wanted = 18
	if course.BasketCount != wanted {
		t.Errorf("got: %d, wanted: %d", course.BasketCount, wanted)
	}

	wanted = 3
	if course.Baskets[18].Par != wanted {
		t.Errorf("got: %d, wanted: %d", course.Baskets[18].Par, wanted)
	}

	tLat := 60.2124424
	tLon := 24.8446608
	// course = create(templates, 60.21315136048604, 24.846823667701181, players, 10, 1)
	course = create(templates, tLat, tLon, players, 10, 1)

	wantedName := "Tali"
	if course.Name != wantedName {
		t.Errorf("got: %s, wanted: %s", course.Name, wantedName)
	}

	wanted = 5
	if course.Baskets[1].Par != wanted {
		t.Errorf("got: %d, wanted: %d", course.Baskets[1].Par, wanted)
	}
}
