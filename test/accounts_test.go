package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"paytabs-task/types"
	"testing"
)

func TestTransactionWithWrongFromAccount(t *testing.T) {
	body := types.TransAction{
		FromId: "wrong_Fron_id",
		ToId:   "ccd1e5cc-c798-4407-883f-f2c62e0d7106",
		Amount: 30,
	}

	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	App.controllers.MakeTransAction(response, req)
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	expected := "one of the accounts ID not found\n"
	if expected != response.Body.String() {
		t.Errorf("'%s', but got '%s'", expected, response.Body.String())
	}

}

func TestTransactionWithWrongToAccount(t *testing.T) {
	body := types.TransAction{
		FromId: "ccd1e5cc-c798-4407-883f-f2c62e0d7106",
		ToId:   "wrong_Fron_id",
		Amount: 30,
	}

	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	App.controllers.MakeTransAction(response, req)
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	expected := "one of the accounts ID not found\n"
	if expected != response.Body.String() {
		t.Errorf("'%s', but got '%s'", expected, response.Body.String())
	}

}

func TestMakeTransactionWithGreaterAmount(t *testing.T) {
	body := types.TransAction{
		FromId: "3d253e29-8785-464f-8fa0-9e4b57699db9",
		ToId:   "ccd1e5cc-c798-4407-883f-f2c62e0d7106",
		Amount: 400,
	}
	fromAccoutBefore, _ := App.db.GetAccountById(body.FromId)
	toAccountBefore, _ := App.db.GetAccountById(body.ToId)
	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	App.controllers.MakeTransAction(response, req)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	expected := "amount greater than balance\n"
	if expected != response.Body.String() {
		t.Errorf("'%s', but got '%s'", expected, response.Body.String())
	}

	fromAccoutAfter, _ := App.db.GetAccountById(body.FromId)
	toAccountAfter, _ := App.db.GetAccountById(body.ToId)
	if fromAccoutAfter.Balance != (fromAccoutBefore.Balance) {
		t.Errorf("Expected fromAccount Balance %v, but got %v", fromAccoutAfter.Balance, fromAccoutBefore.Balance)
	}

	if toAccountAfter.Balance != (toAccountBefore.Balance) {
		t.Errorf("Expected toAccount Balance  %v, but got %v", toAccountAfter.Balance, toAccountBefore.Balance)
	}

}

func TestMakeTransactionSendToSameAccount(t *testing.T) {
	body := types.TransAction{
		FromId: "3d253e29-8785-464f-8fa0-9e4b57699db9",
		ToId:   "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Amount: 50,
	}
	fromAccoutBefore, _ := App.db.GetAccountById(body.FromId)
	toAccountBefore, _ := App.db.GetAccountById(body.ToId)
	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	App.controllers.MakeTransAction(response, req)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	expected := "you cant send to same account\n"
	if expected != response.Body.String() {
		t.Errorf("'%s', but got '%s'", expected, response.Body.String())
	}

	fromAccoutAfter, _ := App.db.GetAccountById(body.FromId)
	toAccountAfter, _ := App.db.GetAccountById(body.ToId)
	if fromAccoutAfter.Balance != (fromAccoutBefore.Balance) {
		t.Errorf("Expected fromAccount Balance %v, but got %v", fromAccoutAfter.Balance, fromAccoutBefore.Balance)
	}

	if toAccountAfter.Balance != (toAccountBefore.Balance) {
		t.Errorf("Expected toAccount Balance  %v, but got %v", toAccountAfter.Balance, toAccountBefore.Balance)
	}

}

func TestMakeTransactionWithNegativeAmount(t *testing.T) {
	body := types.TransAction{
		FromId: "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		ToId:   "ccd1e5cc-c798-4407-883f-f2c62e0d7106",
		Amount: -30,
	}
	fromAccoutBefore, _ := App.db.GetAccountById(body.FromId)
	toAccountBefore, _ := App.db.GetAccountById(body.ToId)
	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	App.controllers.MakeTransAction(response, req)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	expected := "Amount should Have Positive value\n"
	if expected != response.Body.String() {
		t.Errorf("Expected Error %s, but got %s", expected, response.Body.String())
	}

	fromAccoutAfter, _ := App.db.GetAccountById(body.FromId)
	toAccountAfter, _ := App.db.GetAccountById(body.ToId)
	if fromAccoutAfter.Balance != (fromAccoutBefore.Balance) {
		t.Errorf("Expected fromAccount Balance %v, but got %v", fromAccoutAfter.Balance, fromAccoutBefore.Balance)
	}

	if toAccountAfter.Balance != (toAccountBefore.Balance) {
		t.Errorf("Expected toAccount Balance  %v, but got %v", toAccountAfter.Balance, toAccountBefore.Balance)
	}

}

func TestValidMakeTransaction(t *testing.T) {
	body := types.TransAction{
		FromId: "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		ToId:   "ccd1e5cc-c798-4407-883f-f2c62e0d7106",
		Amount: 30,
	}
	fromAccoutBefore, _ := App.db.GetAccountById(body.FromId)
	toAccountBefore, _ := App.db.GetAccountById(body.ToId)
	requestBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	App.controllers.MakeTransAction(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
	}

	fromAccoutAfter, _ := App.db.GetAccountById(body.FromId)
	toAccountAfter, _ := App.db.GetAccountById(body.ToId)
	if fromAccoutAfter.Balance != (fromAccoutBefore.Balance - body.Amount) {
		t.Errorf("Expected fromAccount Balance %v, but got %v", fromAccoutAfter.Balance, fromAccoutBefore.Balance-body.Amount)
	}

	if toAccountAfter.Balance != (toAccountBefore.Balance + body.Amount) {
		t.Errorf("Expected toAccount Balance  %v, but got %v", toAccountAfter.Balance, toAccountBefore.Balance+body.Amount)
	}

}
