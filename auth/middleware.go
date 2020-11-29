package auth

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const UserContextKey = contextKey("user")

type AnonUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (user *AnonUser) GetId() string {
	return user.Id
}

func (user *AnonUser) GetName() string {
	return user.Name
}

func AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, tok := r.URL.Query()["bearer"]
		name, nok := r.URL.Query()["name"]

		if tok && len(token) == 1 {

			user, err := ValidateToken(token[0])
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)

			} else {
				ctx := context.WithValue(r.Context(), UserContextKey, user)
				f(w, r.WithContext(ctx))
			}

		} else if nok && len(name) == 1 {
			// Continue with new Anon. user
			user := AnonUser{Id: uuid.New().String(), Name: name[0]}
			ctx := context.WithValue(r.Context(), UserContextKey, &user)
			f(w, r.WithContext(ctx))

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please login or provide name"))
		}
	})
}
