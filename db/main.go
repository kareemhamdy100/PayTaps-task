package db

import (
	"encoding/json"
	"fmt"
	"os"
	"paytabs-task/types"
	"sync"
)

type Db struct {
	accounts          []*types.Account
	accountsMapWithId map[string]*types.Account
	accountsLock      map[string]*sync.Mutex
}

func (db *Db) LoadDataFromJSONFile() (*Db, error) {

	file, err := os.Open("accounts-mock.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	decoder.UseNumber()

	if err := decoder.Decode(&db.accounts); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	db.accountsMapWithId = map[string]*types.Account{}
	db.accountsLock = map[string]*sync.Mutex{}
	for _, account := range db.accounts {
		db.accountsMapWithId[account.ID] = account
		db.accountsLock[account.ID] = &sync.Mutex{}
	}

	return db, nil

}
