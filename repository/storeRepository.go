package repository

import "sync"

// Store holds a map of key-value pairs
// key for transaction is transactionId + accountId
// key for account is accountId
// although only manager routine accesses the cache
// but cache uses sync to make manager scale easily
type Store struct {
	sync.RWMutex
	transactions map[string]bool
	accounts     map[string]*Account
}

// IStoreRepository defines interface for cache operations
type IStoreRepository interface {
	GetAccount(ID string) *Account
	AddTransaction(transactionID string, accountID string)
	IsDuplicatedTransaction(transactionID string, accountID string) bool
}

// NewStore initializes the cache
func NewStore() *Store {
	return &Store{
		transactions: make(map[string]bool),
		accounts:     make(map[string]*Account),
	}
}

// GetAccount returns an existing account from cache (if available)
// returns a newly created account if not exist
func (s *Store) GetAccount(ID string) *Account {
	s.RLock()
	defer s.RUnlock()

	if acc, ok := s.accounts[ID]; ok {
		return acc
	}

	return s.addAccount(ID)
}

// AddAccount adds a new account to cache
func (s *Store) addAccount(ID string) *Account {
	acc := NewAccount(ID)
	s.accounts[ID] = acc

	return acc
}

// AddTransaction records an entry for an account
func (s *Store) AddTransaction(transactionID string, accountID string) {
	s.Lock()
	defer s.Unlock()

	s.transactions[transactionID+accountID] = true
}

// IsDuplicatedTransaction checks if a transaction is duplicated a an account
func (s *Store) IsDuplicatedTransaction(transactionID string, accountID string) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.transactions[transactionID+accountID]

	return ok
}
