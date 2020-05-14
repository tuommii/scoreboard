package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"miikka.xyz/scoreboard/game"
)

func TestServer(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	s := New("../")
	server := httptest.NewUnstartedServer(s.HTTP.Handler)

	server.Listener.Close()
	server.Listener = l
	server.Start()
	defer server.Close()

	resp, err := http.Get(server.URL + "/_status")
	if err != nil {
		t.Fatal(err)
	}

	wanted := http.StatusOK
	if resp.StatusCode != wanted {
		t.Errorf("WRONG STATUS! GOT: %d, WANTED: %d", resp.StatusCode, wanted)
	}

	resp, err = http.Get(server.URL + "/games/1")
	if err != nil {
		t.Errorf(err.Error())
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("WRONG STATUS! GOT: %d, WANTED: %d", resp.StatusCode, wanted)
	}

	basis := game.Basis{BasketCount: 2, Lat: 0, Lon: 0, Players: []string{"Miikka", "Pasi"}}
	jsonStr, err := json.Marshal(basis)
	if err != nil {
		t.Fatalf(err.Error())
	}
	resp, err = http.Post(server.URL+"/games/create", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("WRONG STATUS! GOT: %d, WANTED: %d", resp.StatusCode, wanted)
	}
}
