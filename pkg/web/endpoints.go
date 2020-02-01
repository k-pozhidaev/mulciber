package web

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func CreateHandlers() {
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := ServerHandleTokenRequest(w, r)
		if err != nil {
			log.Println(fmt.Sprintf("Tocken handler error: %v",err))
		}
	})

	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
		clientId := uuid.New().String()[:8]
		clientSecret := uuid.New().String()[:8]
		err := ClientStoreSet(clientId, clientSecret)
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
		if err != nil {
			log.Println(err.Error())
		}
	})

	http.HandleFunc("/safe", validateToken(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, I'm safe"))
		if err != nil {
			log.Printf("Write error: %s \n", err.Error())
		}
	}))

	http.HandleFunc("/protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
		write, err := w.Write([]byte("Hello, I'm protected"))
		if write == 0 || err != nil {
			log.Println(fmt.Sprintf("Protected handler error: %v",err))
		}

	}))
}

func validateToken(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := ServerValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		f.ServeHTTP(w, r)
	}
}
