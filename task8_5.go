package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpointExample(t *testing.T) {
	// пример “как тестить”: создаём запрос, поднимаем handler, проверяем ответ
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	req := httptest.NewRequest("GET", "http://example.com/hello", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Fatalf("expected 200 got %d", rr.Code)
	}
	if rr.Body.String() != "ok" {
		t.Fatalf("expected ok got %q", rr.Body.String())
	}
}

func TestPOSTJSONExample(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}
		w.WriteHeader(201)
	})

	req := httptest.NewRequest("POST", "http://example.com/users", bytes.NewBufferString(`{"name":"Ann","age":20}`))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != 201 {
		t.Fatalf("expected 201 got %d", rr.Code)
	}
}
