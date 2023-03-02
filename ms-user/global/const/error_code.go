package _const

const (
	Unauthorized      = "unauthorized"
	StatusBadRequest  = "invalid request"
	InternalServerErr = "internal server error"

	UserNotFound           = "wrong email"
	EmailAlreadyExist      = "email already taken"
	PhoneAlreadyExist      = "phone already taken"
	UserWrongPassword      = "wrong password"
	InvalidFieldData       = "invalid user data (missing field or invalid format)"
	MissingEmailOrPassword = "missing email or password"
)
