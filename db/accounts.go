package db

import (
	"errors"
	"paytabs-task/types"
	"sync"
)

func (db *Db) GetAllAccounts() []*types.Account {
	return db.accounts
}

func (db *Db) GetAccountLock(id string) (*sync.Mutex, error) {
	mutex, exists := db.accountsLock[id]
	if !exists {
		return nil, errors.New("this account not found")
	}
	return mutex, nil
}

func (db *Db) GetAccountById(id string) (*types.Account, error) {

	account, exsists := db.accountsMapWithId[id]

	if !exsists {
		return nil, errors.New("account not found")
	}
	accountCopy := *account
	return &accountCopy, nil

}

func (db *Db) UpdateBalance(id string, amount float64) (*types.Account, error) {
	account, exsists := db.accountsMapWithId[id]

	if !exsists {
		return nil, errors.New("account not found")
	}
	account.Balance = amount

	accountCopy := *account
	return &accountCopy, nil

}
