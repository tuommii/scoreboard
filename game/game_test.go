package game

import (
	"testing"
)

func TestCreateID(t *testing.T) {
	id := createID([]string{"TigerKing", "Pesukarhu", "Ying Jang", "Bubba"}, 1)
	wanted := "1bpty"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}

	id = createID([]string{"Aapo", "Ari", "Ying Jang"}, 24)
	wanted = "24aay"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}

	id = createID([]string{"123xxx123"}, 20)
	wanted = "201"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}
}

func TestCreateCourse(t *testing.T) {
	course := CreateCourse([]string{"TigerKing", "Pesukarhu", "Ying Jang"}, 18, 1)

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