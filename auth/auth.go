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

func (res ErrorRes) Status() bool {
	if res.Code != 0 {
		return false
	} else {
		return true
	}
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

func SignUpWithEmailAndPassword(email, password string) SignUpRes {
	request_str := `{"email":"` + email + `", "password":"` + password + `", "returnSecureToken":true}`
	res, err := http.Post("https://identitytoolkit.googleapis.com"+
		"/v1/accounts:signUp?key="+
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

	ret := SignUpRes{}
	sitemap_s := []byte(string(sitemap))
	json.Unmarshal(sitemap_s, &ret)

	return ret
}

func GetUserDetail(user User) UserDetailRes {
	res, err := http.Post("https://identitytoolkit.googleapis.com"+
		"/v1/accounts:lookup?key="+apiKey, "application/json",
		bytes.NewBuffer([]byte(`{"idToken":"`+user.IdToken+`"}`)))

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	sitemap, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ret := UserDetailRes{}
	sitemap_s := []byte(string(sitemap))
	json.Unmarshal(sitemap_s, &ret)

	return ret
}

func updateUser(idToken string, data map[string]string) UserDetailRes {
	str := `{"idToken":"` + idToken + `", "returnSecureToken":true`
	for key, value := range data {
		str += `,"` + key + `":"` + value + `"`
	}
	str += `}`

	res, err := http.Post("https://identitytoolkit.googleapis.com"+
		"/v1/accounts:update?key="+apiKey, "application/json",
		bytes.NewBuffer([]byte(str)))

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	sitemap, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ret := UserDetailRes{}
	sitemap_s := []byte(string(sitemap))
	json.Unmarshal(sitemap_s, &ret)

	return ret
}

func UpdatePassword(user User, password string) UserDetailRes {
	data := map[string]string{
		"password": password,
	}
	return updateUser(user.IdToken, data)
}

func UpdateProfile(user User, data map[string]string) UserDetailRes {
	return updateUser(user.IdToken, data)
}
