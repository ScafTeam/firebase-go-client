# Firebase Auth For Go

基於 Go http 實現 Firebase Auth

## 安裝

```bash
go get github.com/ScafTeam/firebase-go-client
```

## 使用

```go
package main

import (
  "fmt"

  "github.com/scaf-team/firebase-go-client"
)
```

### 初始化

```go
auth.Auth("[YOUR FIREBASE API KEY]")
```

### 使用信箱與密碼登入

```go
res := auth.SignInWithEmailAndPassword("[EMAIL]", "[PASSWORD]")
if res.Status() {
  user = res.Result()
  fmt.Println(user.Email + " is signed in")
} else {
  // EMAIl_NOT_FOUND 沒有此用戶
  // INVALID_PASSWORD 密碼錯誤
  // USER_DISABLED 帳號被停用
  fmt.Println(res.Error.Message)
}
```

### 使用信箱與密碼註冊

```go
res := auth.CreateUserWithEmailAndPassword("[EMAIL]", "[PASSWORD]")
if res.Status() {
  user = res.Result()
  fmt.Println(user.Email + " is signed up")
} else {
  // EMAIL_EXISTS 用戶已存在
  // OPERATION_NOT_ALLOWED 無法使用此功能
  // WEAK_PASSWORD 密碼太簡單
  fmt.Println(res.ErrorMessage())
}
```

### 獲取使用者資料

```go
res := auth.GetUserDetail(user)
if res.Status() {
  UserDetail := res.Result()
  fmt.Println(UserDetail.Email)
} else {
  // INVALID_ID_TOKEN 用戶身份驗證失敗
  // USER_NOT_FOUND 沒有此用戶
  fmt.Println(res.ErrorMessage())
}
```

資料包含

```json
{
  "LocalId": "[The uid of the current User]",
  "Email": "[The email of the current User]",
  "EmailVerified ": "[Whether the email of the current User has been verified]",
  "DisplayName ": "[The display name of the current User]",
  "PhotoUrl": "[The photo URL of the current User]",
  "PasswordHash": "[The password hash of the current User]",
  "PasswordUpdatedAt": "[The timestamp when the password of the current User was last updated]",
  "ValidSince": "[The timestamp when the ID token of the current User was issued]",
  "Disabled": "[Whether the current User is disabled]",
  "LastLoginAt": "[The timestamp when the current User last signed in]",
  "CreatedAt": "[The timestamp when the current User was created]",
}
```

### 更新使用者密碼

```go
res := auth.UpdatePassword(user, "[new password]")
if res.Status() {
  fmt.Println("Password is updated")
} else {
  // INVALID_ID_TOKEN 用戶身份驗證失敗
  // USER_NOT_FOUND 沒有此用戶
  // WEAK_PASSWORD 密碼太簡單
  fmt.Println(res.ErrorMessage())
}
```

Remember: 修改完密碼後，應需要重新登入，重新獲取新的 token

### 更新使用者資料

```go
res := auth.UpdateProfile(user, map[string]string{
  "[user profile]": "[new value]",
})
if res.Status() {
  fmt.Println("Profile is updated")
} else {
  // INVALID_ID_TOKEN 用戶身份驗證失敗
  // USER_NOT_FOUND 沒有此用戶
  fmt.Println(res.ErrorMessage())
}
```
