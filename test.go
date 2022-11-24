package main

import (
	"firebase-go-client/auth"
	"fmt"
)

func main() {
	auth.Auth("AIzaSyBUTf0DULItJDOl1t6tvTZ8_sP8_wL-cPg")
	res := auth.SignInWithEmailAndPassword("test@test.com", "testtest")
	if res.Status() {
		fmt.Println(res.Email + " is signed in")
	} else {
		// EMAIl_NOT_FOUND 沒有此用戶
		// INVALID_PASSWORD 密碼錯誤
		// USER_DISABLED 帳號被停用
		fmt.Println(res.Error.Message)
	}
}
