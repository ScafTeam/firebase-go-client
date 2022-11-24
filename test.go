package main

import (
	"fmt"
	"github.com/ScafTeam/firebase-go-client/auth"
)

func signIn() auth.User {
	var user auth.User
	res := auth.SignInWithEmailAndPassword("test@test.com", "testtest1")
	if res.Status() {
		user = res.Result()
		fmt.Println(user.Email + " is signed in")
	} else {
		// EMAIl_NOT_FOUND 沒有此用戶
		// INVALID_PASSWORD 密碼錯誤
		// USER_DISABLED 帳號被停用
		fmt.Println(res.ErrorMessage())
	}
	return user
}

func signUp() auth.User {
	var user auth.User
	res := auth.SignUpWithEmailAndPassword("test1@test.com", "testtest")
	if res.Status() {
		user = res.Result()
		fmt.Println(user.Email + " is signed up")
	} else {
		// EMAIL_EXISTS 用戶已存在
		// OPERATION_NOT_ALLOWED 無法使用此功能
		// WEAK_PASSWORD 密碼太簡單
		fmt.Println(res.ErrorMessage())
	}
	return user
}

func getUserDetail(user auth.User) auth.UserDetail {
	var userDetail auth.UserDetail
	res := auth.GetUserDetail(user)
	if res.Status() {
		UserDetail := res.Result()
		fmt.Println(UserDetail.Email)
	} else {
		// INVALID_ID_TOKEN 用戶身份驗證失敗
		// USER_NOT_FOUND 沒有此用戶
		fmt.Println(res.ErrorMessage())
	}
	return userDetail
}

func updatePassword(user auth.User) {
	res := auth.UpdatePassword(user, "testtest")
	if res.Status() {
		fmt.Println("Password is updated")
	} else {
		// INVALID_ID_TOKEN 用戶身份驗證失敗
		// USER_NOT_FOUND 沒有此用戶
		// WEAK_PASSWORD 密碼太簡單
		fmt.Println(res.ErrorMessage())
	}
}

func UpdateProfile(user auth.User) {
	res := auth.UpdateProfile(user, map[string]string{
		"displayName": "test",
	})
	if res.Status() {
		fmt.Println("Profile is updated")
	} else {
		// INVALID_ID_TOKEN 用戶身份驗證失敗
		// USER_NOT_FOUND 沒有此用戶
		fmt.Println(res.ErrorMessage())
	}
}

func main() {
	auth.Auth("AIzaSyBUTf0DULItJDOl1t6tvTZ8_sP8_wL-cPg")
	// var res auth.Res
	var user auth.User

	user = signIn()
	// user = signUp()
	getUserDetail(user)
	UpdateProfile(user)
	updatePassword(user)
}
