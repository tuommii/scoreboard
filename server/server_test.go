package server

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandle(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	s := New("../")
	server := httptest.NewUnstartedServer(s.HTTP.Handler)

	// Change listener
	server.Listener.Close()
	server.Listener = l
	server.Start()
	defer server.Close()

	resp, err := http.Get("http://localhost:8080/_status")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("WRONG STATUS: %d", resp.StatusCode)
	}

	resp, err = http.Get("http://localhost:8080/games/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 500 {
		t.Errorf("WRONG STATUS: %d", resp.StatusCode)
	}
}
