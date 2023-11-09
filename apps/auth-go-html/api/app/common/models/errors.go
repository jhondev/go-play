package models

import "errors"

var (
	ErrRecordAlreadyExists = errors.New("record already exists")
	ErrRecordNotFound      = errors.New("record not found")
	ErrEditConflict        = errors.New("edit conflict")
)
