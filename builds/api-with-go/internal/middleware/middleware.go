package middleware

import "errors"

var unauthorizedError = errors.New("Invalid username token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username") //get the username from the query parameters
		var tokenn string = r.URL.Query().Get("Authorization") 
		var err error

		if username == "" || tokenn == "" { //check if the username or token is empty

	}

}