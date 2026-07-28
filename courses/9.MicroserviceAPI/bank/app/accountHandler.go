package app

import (
	"bank/dto"
	"bank/errs"
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

func (h AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := h.service.MakeTransaction(request)
		if appError != nil {
			appErr := appError.(*errs.AppError)
			writeResponse(w, appErr.Code, appErr)
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
