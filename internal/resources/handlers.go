package resources

import (
	"fmt"
	"github.com/ilijamt/vault-token-helper/internal/handler"
	"github.com/ilijamt/vault-token-helper/internal/handler/localpath"
)

var handlers = map[string]handler.Handler{
	localpath.NameLocalPath: localpath.New(),
}

// Available returns all available handler
func Available() (hndlrs []handler.Handler) {
	for _, h := range handlers {
		hndlrs = append(hndlrs, h)
	}
	return hndlrs
}

// Get retrieves a handler if it's defined
func Get(name string) (handler.Handler, error) {
	if h, ok := handlers[name]; ok {
		return h, nil
	}
	return nil, fmt.Errorf("missing handler: %s", name)
}
