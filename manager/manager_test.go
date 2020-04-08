package manager

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateID(t *testing.T) {
	id := createID([]string{"TigerKing", "Pesukarhu", "Ying Jang", "Bubba"}, 1)
	if id != "1bpty" {
		t.Errorf("%s", id)
	}
}

func TestCreateCourse(t *testing.T) {
	course := CreateCourse([]string{"TigerKing", "Pesukarhu", "Ying Jang"}, 18)
	if len(course.Baskets) != 18 {
		t.Errorf("%+v", course)
	}
	if len(course.Baskets[1].Scores) != 3 {
		t.Errorf("%+v", len(course.Baskets[1].Scores))
	}
}
func TestJSONToCourse(t *testing.T) {
	jsonFile, err := os.Open("../example.json")
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Error(err)
	}
	g := JSONToCourse(string(bytes))
	wanted := 3
	if len(g.Baskets) != wanted {
		t.Errorf("%d, %d\n", len(g.Baskets), wanted)
	}

	wanted = 4
	if g.Baskets[2].Par != wanted {
		t.Errorf("%d, %d\n", g.Baskets[2].Par, wanted)
	}
}
