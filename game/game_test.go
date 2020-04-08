package game

import "testing"

func TestGame(t *testing.T) {
	wanted := "1"
	got := "1"
	if got != wanted {
		t.Errorf("[%s] got, [%s] wanted\n", got, wanted)
	}
}
