package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("supersecret")

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func makeToken(role string) (string, error) {
	claims := Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}

func auth(requiredRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "no token", 401)
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")

		tok, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
			return secret, nil
		})
		if err != nil || !tok.Valid {
			http.Error(w, "bad token", 401)
			return
		}

		claims := tok.Claims.(*Claims)
		if requiredRole != "" && claims.Role != requiredRole {
			http.Error(w, "forbidden", 403)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// для лабы: роль выбираем по query ?role=admin или user
		role := r.URL.Query().Get("role")
		if role == "" {
			role = "user"
		}
		t, _ := makeToken(role)
		json.NewEncoder(w).Encode(map[string]string{"token": t})
	})

	http.Handle("/admin", auth("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello admin\n"))
	})))

	http.Handle("/profile", auth("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello authorized user\n"))
	})))

	log.Println("JWT server :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
