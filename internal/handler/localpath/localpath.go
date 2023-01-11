package localpath

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/url"
	"os"
	"os/user"

	"github.com/ilijamt/vault-token-helper/internal/handler"
)

const (
	// NameLocalPath is the name of the handler
	NameLocalPath = "LocalPath"
	// VaultTokenPath is the default token path where we will store the data regarding the vault tokens
	VaultTokenPath = "%s/.vault-tokens"
)

// LocalPath is a vault token handler implementing handler.Handler, this handler is used to
// handle vault tokens based on the vault address.
type LocalPath struct {
	TokenDirectory string
}

// New create a new vault token handler with the default vault path specified as VaultTokenPath
func New() handler.Handler {
	usr, _ := user.Current()
	homedir := usr.HomeDir
	return NewWithPath(fmt.Sprintf(VaultTokenPath, homedir))
}

// NewWithPath create a new vault token handler with a specified vault path
func NewWithPath(path string) handler.Handler {
	return &LocalPath{TokenDirectory: path}
}

// Description return the description of the handler
func (l LocalPath) Description() string {
	return fmt.Sprintf(`Stores all the vault tokens under %s/ (%s) with vault hostname md5 encoded`,
		l.TokenDirectory, localStoragePerm.String())
}

// Name return the name of the handler
func (l LocalPath) Name() string {
	return NameLocalPath
}

// Get retrieves a token associated with the vaultAddr from LocalPath.TokenDirectory, returns an error
// if there is an issue with creating, reading the data from the directory
func (l LocalPath) Get(vaultAddr *url.URL) (token string, err error) {
	if vaultAddr == nil {
		return token, fmt.Errorf("get failed: %w", handler.ErrNoVaultAddress)
	}
	var path string
	if path, err = preparePath(vaultAddr, l.TokenDirectory); err != nil {
		return token, err
	}
	if _, err = os.Stat(path); err == nil {
		var payload []byte
		if payload, err = os.ReadFile(path); err == nil {
			var data handler.Data
			if err = gob.NewDecoder(bytes.NewReader(payload)).Decode(&data); err == nil {
				token = data.Token
			}
		}
		return token, err
	}
	return "", err
}

// Store a token associated with the vaultAddr, and returns error if it fails to store the
// vault token in the handler, otherwise nil
func (l LocalPath) Store(token string, vaultAddr *url.URL) (err error) {
	if vaultAddr == nil {
		return fmt.Errorf("store failed: %w", handler.ErrNoVaultAddress)
	}
	var path string
	if path, err = preparePath(vaultAddr, l.TokenDirectory); err != nil {
		return err
	}
	var data = handler.Data{
		Address: vaultAddr.String(),
		Token:   token,
	}
	var payload bytes.Buffer
	_ = gob.NewEncoder(&payload).Encode(data)
	return os.WriteFile(path, payload.Bytes(), 0600)
}

// Erase removes a token associated with the vaultAddr, and returns error if it fails to erase the
// vault token in the handler, otherwise nil
func (l LocalPath) Erase(vaultAddr *url.URL) (err error) {
	if vaultAddr == nil {
		return fmt.Errorf("erase failed: %w", handler.ErrNoVaultAddress)
	}
	var path string
	if path, err = preparePath(vaultAddr, l.TokenDirectory); err != nil {
		return err
	}

	return os.Remove(path)
}

// Path is the directory where the data is stored for the vault tokens
func (l LocalPath) Path() string {
	return l.TokenDirectory
}

var _ handler.Handler = (*LocalPath)(nil)
