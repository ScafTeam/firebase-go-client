package auth

type User struct {
	IdToken      string
	Email        string
	RefreshToken string
	ExpiresIn    string
	LocalId      string
}
