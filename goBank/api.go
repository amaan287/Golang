package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	store      Storage
}

func NewApiServer(listenAddr string, store Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}
func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccounts))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountById))
	fmt.Println("Server is running on http://localhost", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		fmt.Println("Error", err)
	}
}
func (s *ApiServer) handleAccounts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccunts(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccunts(w, r)
	}
	return fmt.Errorf("method not allowed, your method: %s", r.Method)

}
func (s *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}
	account, err := s.store.GetAccountbById(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, account)

}
func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err

	}
	return writeJSON(w, http.StatusOK, accounts)

}
func (s *ApiServer) handleCreateAccunts(w http.ResponseWriter, r *http.Request) error {

	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return err
	}
	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, account)

}
func (s *ApiServer) handleDeleteAccunts(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
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
