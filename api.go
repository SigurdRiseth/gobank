package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunction func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("unsupported method: %s", r.Method)
}

// GET /account
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

// GET /account{id}
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println("Fetching account with ID:", id)

	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}

	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName, createAccountRequest.Email)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
