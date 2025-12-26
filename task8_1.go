package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	mu    sync.Mutex
	next  int
	users map[int]User
}

func NewStore() *Store {
	return &Store{next: 1, users: make(map[int]User)}
}

func (s *Store) list() []User {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]User, 0, len(s.users))
	for _, u := range s.users {
		out = append(out, u)
	}
	return out
}

func (s *Store) get(id int) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *Store) create(u User) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	u.ID = s.next
	s.next++
	s.users[u.ID] = u
	return u
}

func (s *Store) update(id int, u User) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return User{}, false
	}
	u.ID = id
	s.users[id] = u
	return u, true
}

func (s *Store) delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	store := NewStore()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, 200, store.list())
		case http.MethodPost:
			var u User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				http.Error(w, "bad json", 400)
				return
			}
			created := store.create(u)
			writeJSON(w, 201, created)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "bad id", 400)
			return
		}

		switch r.Method {
		case http.MethodGet:
			u, ok := store.get(id)
			if !ok {
				http.NotFound(w, r)
				return
			}
			writeJSON(w, 200, u)

		case http.MethodPut:
			var u User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				http.Error(w, "bad json", 400)
				return
			}
			updated, ok := store.update(id, u)
			if !ok {
				http.NotFound(w, r)
				return
			}
			writeJSON(w, 200, updated)

		case http.MethodDelete:
			if !store.delete(id) {
				http.NotFound(w, r)
				return
			}
			w.WriteHeader(204)

		default:
			http.Error(w, "method not allowed", 405)
		}
	})

	log.Println("REST :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
