package controllers

import (
	"encoding/json"
	"net/http"
	"paytabs-task/types"
)

func (c *Controllers) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	accounts, err := c.services.GetAllAccounts()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(accounts)

}

func (c *Controllers) MakeTransAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var trasactionBody types.TransAction

	if err := json.NewDecoder(r.Body).Decode(&trasactionBody); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if trasactionBody.Amount <= 0 {
		http.Error(w, "Amount should Have Positive value", http.StatusBadRequest)
		return
	}

	account, err := c.services.MakeTransAction(&trasactionBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)

}
