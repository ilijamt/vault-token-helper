package handler

import "errors"

var (
	// ErrNoVaultAddress means that we don't have a valid vault address
	ErrNoVaultAddress = errors.New("no vault address defined")
)
