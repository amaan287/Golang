package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiServer struct {
	listenAddr string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//Handle error here
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}
func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccunts))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccunts))

	fmt.Println("Server is running on http://localhost", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		fmt.Println("Error", err)
	}
}
func (s *ApiServer) handleAccunts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccunts(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccunts(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccunts(w, r)
	}
	return fmt.Errorf("method not allowed, your method: %s", r.Method)

}
func (s *ApiServer) handleGetAccunts(w http.ResponseWriter, r *http.Request) error {
	//id := mux.Vars(r)["id"]
	//account := NewAccount("Amaan", "Mirza")
	return writeJSON(w, http.StatusOK, &Account{})

}
func (s *ApiServer) handleCreateAccunts(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (s *ApiServer) handleDeleteAccunts(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil

}
