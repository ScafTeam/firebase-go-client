package auth

type SignInRes struct {
	Error        ErrorRes `json:"error"`
	Email        string   `json:"email"`
	LocalId      string   `json:"localId"`
	DisplayName  string   `json:"displayName"`
	IdToken      string   `json:"idToken"`
	Registered   bool     `json:"registered"`
	RefreshToken string   `json:"refresh"`
	ExpiresIn    string   `json:"expiresIn"`
}

func (res SignInRes) Status() bool {
	if res.Error.Code != 0 {
		return false
	} else {
		return true
	}
}

func (res SignInRes) Message() string {
	return res.Error.Message
}

func (res SignInRes) User() string {
	return res.Email
}
