package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ErrorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Message string `json:"message"`
		Domain  string `json:"domain"`
		Reason  string `json:"reason"`
	} `json:"errors"`
}

var apiKey string

func Auth(_apiKey string) {
	apiKey = _apiKey
}

func SignInWithEmailAndPassword(email, password string) SignInRes {
	request_str := `{"email":"` + email + `", "password":"` + password + `", "returnSecureToken":true}`
	res, err := http.Post("https://identitytoolkit.googleapis.com"+
		"/v1/accounts:signInWithPassword?key="+
		apiKey, // API key
		"application/json",
		bytes.NewBuffer([]byte(request_str)))

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	sitemap, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ret := SignInRes{}
	sitemap_s := []byte(string(sitemap))
	json.Unmarshal(sitemap_s, &ret)

	return ret
}
