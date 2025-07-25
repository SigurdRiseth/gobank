package main

import (
	"testing"
	"time"
)

// MockStore implements the Storage interface for testing
type MockStore struct {
	MockAccounts []*Account
	MockError error
}

func (m *MockStore) GetAccounts() ([]*Account, error) {
	return m.MockAccounts, m.MockError
}

func (m *MockStore) CreateAccount(*Account) error { return nil }
func (m *MockStore) DeleteAccount(int) error { return nil }
func (m *MockStore) UpdateAccount(*Account) error { return nil }
func (m *MockStore) GetAccountByID(int) (*Account, error) { return nil, nil }

func TestGetAccounts(t *testing.T) {
	mock := &MockStore{
		MockAccounts: []*Account{
			{
				ID: 1,
				FirstName: "Alice",
				LastName: "Smith",
				Email: "alice@example.com",
				Number: 123,
				Balance: 100,
				CreatedAt: time.Now(),
			},
		},
		MockError: nil,
	}

	accounts, err := mock.GetAccounts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(accounts) != 1 {
		t.Errorf("expected 1 account, got %d", len(accounts))
	}
	if accounts[0].FirstName != "Alice" {
		t.Errorf("expected account first name to be Alice, got %s", accounts[0].FirstName)
	}
}
