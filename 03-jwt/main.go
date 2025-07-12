package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

var (
	jwtAccessTokenSecret  = "keyboard cat"
	jwtRefreshTokenSecret = "keyboard dog"
	jwtAccessTokenExpire  = 2
	jwtRefreshTokenExpire = 168
)

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Name     string `json:"name" form:"name" binding:"required`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.Replace(authorization, "Bearer ", "", 1)
		toke, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return []byte(jwtAccessTokenSecret), nil
		})

		if err != nil || !toke.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := toke.Claims.(jwt.MapClaims)
		if !ok || !toke.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", id)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	userRouter := api.PathPrefix("/users").Subrouter()
	userRouter.Use(Middleware)

	userRouter.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id")
		if userID == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Protected route accessed by user ID: " + userID.(string)))
	}).Methods("GET")

	r.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Public route accessed"))
	}).Methods("GET")

	// r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Signup successful"))
	// }).Methods("POST")

	// r.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Signin successful"))
	// }).Methods("POST")

	http.ListenAndServe(":3050", r)
}
