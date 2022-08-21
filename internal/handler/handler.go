package handler

import (
	"net/url"
)

// Handler is a generic interface to which the Vault Token Helper must conform to
type Handler interface {
	// Name of the handler
	Name() string
	// Description of the handler
	Description() string
	// Path is the path where the data is stored
	Path() string
	// Get retrieves a token associated with the vaultAddr, and returns error if not being able to
	// retrieve a token
	Get(vaultAddr *url.URL) (string, error)
	// Store a token associated with the vaultAddr, and returns error if it fails to store the
	// vault token in the handler, otherwise nil
	Store(token string, vaultAddr *url.URL) error
	// Erase removes a token associated with the vaultAddr, and returns error if it fails to erase the
	// vault token in the handler, otherwise nil
	Erase(vaultAddr *url.URL) error
}
