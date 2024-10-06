package myError

import "errors"

var (
	ErrMissingRequestID = errors.New("отсутствует Request Id")
	ErrServer           = errors.New("произошла ошибка на сервере, попробуйте позже")
)
