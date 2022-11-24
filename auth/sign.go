package auth

type SignInRes struct {
	Error ErrorRes `json:"error"`
	User
}

func (res SignInRes) Status() bool {
	return res.Error.Status()
}

func (res SignInRes) Result() User {
	return res.User
}

func (res SignInRes) ErrorMessage() string {
	return res.Error.Message
}

type SignUpRes struct {
	Error ErrorRes `json:"error"`
	User
}

func (res SignUpRes) Status() bool {
	return res.Error.Status()
}

func (res SignUpRes) Result() User {
	return res.User
}

func (res SignUpRes) ErrorMessage() string {
	return res.Error.Message
}
