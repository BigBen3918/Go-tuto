package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/registrations", registrationHandler)
	http.HandleFunc("/authentications", authentications)
	http.HandleFunc("/test", testResourceHandler)

	http.ListenAndServe(":8081", nil)
}

func registrationHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	if req.FormValue("username") == "" || req.FormValue("password") == "" {
		fmt.Fprint(w, "Please enter valid username and password.\r\n")
	} else {
		response, err := registerUser(req.FormValue("username"), req.FormValue("password"))

		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, response)
		}
	}
}

func authentications(w http.ResponseWriter, req *http.Request) {

	username, password, ok := req.BasicAuth()

	if ok {
		tokenDetails, err := generateToken(username, password)

		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			enc := json.NewEncoder(w)
			enc.SetIndent("", "  ")
			enc.Encode(tokenDetails)
		}
	} else {
		fmt.Fprintf(w, "You require a username/password to get a token\r\n")
	}
}

func testResourceHandler(w http.ResponseWriter, req *http.Request) {
	authToken := strings.Split(req.Header.Get("Authentication"), "Bearer")[1]

	userDetails, err := validateToken(authToken)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		username := fmt.Sprint(userDetails["username"])
		fmt.Fprintf(w, "Welcome, "+username+"\r\n")
	}
}
