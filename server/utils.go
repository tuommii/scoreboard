package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"miikka.xyz/scoreboard/game"
)

func gameFromJSON(body io.ReadCloser) (*game.Course, []byte, error) {
	var c *game.Course
	bytes, err := ioutil.ReadAll(io.LimitReader(body, maxBodySize))
	if err != nil {
		return c, bytes, err
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return c, bytes, err
	}
	return c, bytes, nil
}

func parseBasis(body io.ReadCloser) (game.Basis, error) {
	var basis game.Basis

	bytes, err := ioutil.ReadAll(io.LimitReader(body, maxBodySize))
	if err != nil {
		return basis, err
	}
	err = json.Unmarshal(bytes, &basis)
	if err != nil {
		return basis, err
	}
	return basis, nil
}

func jsonErr(msg string) string {
	return fmt.Sprintf(`{"err":"%s"}`, msg)
}

func text(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

// Worker for future use
func (s *Server) worker(lat float64, lon float64) {
	log.Println("simulating api request with:", lat, lon)
	time.Sleep(time.Second * 10)
	log.Println("api simulation done!")
}

func (s *Server) updateCounter() {
	s.mu.Lock()
	if s.counter > maxGames {
		s.counter = 1
	}
	s.counter++
	s.mu.Unlock()
}

func (s *Server) clean(editedAgo time.Duration, createdAgo time.Duration) {
	// s.rw.Lock()
	// defer s.rw.Unlock()

	// easier id
	if s.games2.Count() == 0 {
		s.counter = 1
		return
	}

	for temp := range s.games2.IterBuffered() {
		game := temp.Val.(*game.Course)
		if time.Since(game.EditedAt) > editedAgo || time.Since(game.CreatedAt) > createdAgo {
			s.games2.Remove(temp.Key)
			// delete(s.games, id)
			log.Println("deleted", temp.Key, game.Name)
		}
	}
	// for id, temp := range s.games2 {
	// 	game := temp.(*game.Course)
	// 	if time.Since(game.EditedAt) > editedAgo || time.Since(game.CreatedAt) > createdAgo {
	// 		delete(s.games, id)
	// 		log.Println("deleted", id, game.Name)
	// 	}
	// }
}
