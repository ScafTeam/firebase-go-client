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
  "log"

  "github.com/scaf-team/firebase-go-client"
)
```

### 初始化

```go
auth.Auth("[YOUR FIREBASE API KEY]")
```

### 使用信箱與密碼登入

```go
res := auth.SignInWithEmail("[EMAIL]", "[PASSWORD]")
if res.Status() {
  fmt.Println(res.Email + " is signed in")
} else {
  // EMAIl_NOT_FOUND 沒有此用戶
  // INVALID_PASSWORD 密碼錯誤
  // USER_DISABLED 帳號被停用
  fmt.Println(res.Error.Message)
}
```