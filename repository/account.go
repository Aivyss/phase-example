package repository

import (
	"errors"
	"fmt"
	"phase_example/dto/account"
	"sync"
)

type AccountRepository interface {
	Insert(userID string, password string) error
}

type accountRepository struct {
	sequence int
	mutex    sync.Mutex
	memory   map[int]account.Account
	idSet    map[string]bool
}

func (a *accountRepository) Insert(userID string, password string) error {
	a.mutex.Lock()
	if a.idSet[userID] {
		a.mutex.Unlock()
		return errors.New("duplicated user id")
	}

	accnt := account.Account{
		ID:         a.sequence,
		UserID:     userID,
		Password:   password,
		IsVerified: false,
	}
	a.memory[accnt.ID] = accnt
	a.idSet[accnt.UserID] = true
	a.sequence++
	a.mutex.Unlock()

	return nil
}

func NewAccountRepository() AccountRepository {
	fmt.Println("autowired: AccountRepository")

	return &accountRepository{
		sequence: 0,
		memory:   map[int]account.Account{},
		idSet:    map[string]bool{},
	}
}
