package manager

import (
	"io/ioutil"
	"os"
	"testing"
)

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
	if len(g.Baskets) != 4 {
		t.Errorf("%d, %d\n", len(g.Baskets), 4)
	}
	if g.Baskets[2].Par != 4 {
		t.Errorf("%d, %d\n", g.Baskets[2].Par, 4)
	}
}
