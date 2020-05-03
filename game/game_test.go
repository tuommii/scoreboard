package game

import (
	"testing"
)

func TestCreateID(t *testing.T) {
	id := createID([]string{"TigerKing", "Pesukarhu", "Jian Yang", "Bubba"}, 1)
	wanted := "bpty1"
	if id != wanted {
		t.Errorf("got: %s, wanted: %s\n", id, wanted)
	}

	id = createID([]string{"Aapo", "Ari", "Jiang Yang"}, 24)
	wanted = "aay24"
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
