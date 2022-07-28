package handler

import (
	"net/url"
)

type Handler interface {
	Name() string
	Description() string
	Get(vaultAddr *url.URL) (string, error)
	Store(token string, vaultAddr *url.URL) error
	Erase(vaultAddr *url.URL) error
}
