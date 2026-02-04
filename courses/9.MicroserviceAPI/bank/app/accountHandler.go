package app

import (
	"bank/dto"
	"bank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	account, err := ah.service.NewAccount(request)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, http.StatusCreated, account)
}
