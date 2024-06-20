package model

import (
	"errors"
)
// define error types based on business logic
var (
	ERROR_USER_NOTEXISTS = errors.New("User Doesn't Exist")
	ERROR_USER_EXISTS = errors.New("User Already Exists")
	ERROR_USER_PWD = errors.New("Invalid Password")
)