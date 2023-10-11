package types

import "sync"

type AccountLock struct {
	Mutex sync.Mutex
}
type Account struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance,string"`
}

type TransAction struct {
	FromId string  `json:"fromId"`
	ToId   string  `json:"toId"`
	Amount float64 `json:"amount"`
}
