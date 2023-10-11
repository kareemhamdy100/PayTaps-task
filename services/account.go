package services

import (
	"errors"
	"paytabs-task/types"
)

func (s *Services) GetAllAccounts() ([]*types.Account, error) {
	allAccounts := s.db.GetAllAccounts()
	return allAccounts, nil
}

func (s *Services) MakeTransAction(trasactionDTO *types.TransAction) (*types.Account, error) {
	if trasactionDTO.FromId == trasactionDTO.ToId {
		return nil, errors.New("you cant send to same account")
	}

	fromAccountMutex, errFrom := s.db.GetAccountLock(trasactionDTO.FromId)
	toAccountMutex, errTo := s.db.GetAccountLock(trasactionDTO.ToId)

	if errFrom != nil || errTo != nil {
		return nil, errors.New("one of the accounts ID not found")
	}

	if isAllowLockAccountFrom := fromAccountMutex.TryLock(); !isAllowLockAccountFrom {
		return nil, errors.New("from account is currently busy; another transaction is in progress")
	}
	defer fromAccountMutex.Unlock()

	if isAllowLockAccountTo := toAccountMutex.TryLock(); !isAllowLockAccountTo {
		return nil, errors.New("to account is currently busy; another transaction is in progress")
	}
	defer toAccountMutex.Unlock()

	fromAccount, err := s.db.GetAccountById(trasactionDTO.FromId)
	if err != nil {
		return nil, err
	}

	toAccount, err := s.db.GetAccountById(trasactionDTO.ToId)

	if err != nil {
		return nil, err
	}

	if fromAccount.Balance < trasactionDTO.Amount {
		return nil, errors.New("amount greater than balance")
	}

	fromAccountAfterUpdate, err := s.db.UpdateBalance(fromAccount.ID, fromAccount.Balance-trasactionDTO.Amount)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.UpdateBalance(toAccount.ID, toAccount.Balance+trasactionDTO.Amount); err != nil {
		return nil, err
	}
	return fromAccountAfterUpdate, nil
}
