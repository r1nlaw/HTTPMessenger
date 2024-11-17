package handlers

import (
	"HTTPMessenger/user"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var (
	message []user.Message
	mut     sync.Mutex
)

func GetMessege(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	mut.Lock()

	if len(message) == 0 {
		mut.Unlock()
		time.Sleep(10 * time.Second)
		mut.Lock()
	} else {
		json.NewEncoder(w).Encode(message)
	}
	mut.Unlock()

}

func PostMessege(w http.ResponseWriter, r *http.Request) {
	var msg user.Message

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mut.Lock()
	message = append(message, msg)
	mut.Unlock()

	w.WriteHeader(http.StatusNoContent)

}
