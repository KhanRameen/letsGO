package api

import (
	"encoding/json"
	"net/http"
	"fmt"
)

//Params Struct
type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	Code int //status code
	Message string //response message
	Balance float64 //response data
}

type ErrorResponse struct{
	Code int //status code
	Message string //response message
}

func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	res := Error{
		Code : code,
		Message : message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
	
var ( RequestErrorHandler = func(w http.ResponseWriter, r *http.Request) {