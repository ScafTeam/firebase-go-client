package auth

type User struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
	Registered   bool   `json:"registered"`
}

type UserDetail struct {
	LocalId           string `json:"localId"`
	Email             string `json:"email"`
	EmailVerified     bool   `json:"emailVerified"`
	DisplayName       string `json:"displayName"`
	PhotoUrl          string `json:"photoUrl"`
	PasswordHash      string `json:"passwordHash"`
	PasswordUpdatedAt int64  `json:"passwordUpdatedAt"`
	ValidSince        string `json:"validSince"`
	Disabled          bool   `json:"disabled"`
	LastLoginAt       string `json:"lastLoginAt"`
	CreatedAt         string `json:"createdAt"`
}

type UserDetailRes struct {
	Error      ErrorRes     `json:"error"`
	UserDetail []UserDetail `json:"users"`
}

func (res UserDetailRes) Status() bool {
	return res.Error.Status()
}

func (res UserDetailRes) Result() UserDetail {
	return res.UserDetail[0]
}

func (res UserDetailRes) ErrorMessage() string {
	return res.Error.Message
}
