package localpath

import (
	"fmt"
	"github.com/ilijamt/vault-token-helper/internal/handler"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
)

const (
	NameLocalPath = "LocalPath"
)

type LocalPath struct {
	TokenDirectory string
}

func New() handler.Handler {
	usr, _ := user.Current()
	homedir := usr.HomeDir
	return &LocalPath{TokenDirectory: fmt.Sprintf("%s/.vault-tokens", homedir)}
}

func (l LocalPath) Description() string {
	return fmt.Sprintf(`Stores all the vault tokens under %s/ (%s) with vault hostname md5 encoded`,
		l.TokenDirectory, localStoragePerm.String())
}

func (l LocalPath) Name() string {
	return NameLocalPath
}

func (l LocalPath) Get(vaultAddr *url.URL) (token string, err error) {
	var path string
	if path, err = getPath(l.TokenDirectory, vaultAddr); err != nil {
		return "", err
	}
	if err = localStorage(l.TokenDirectory); err != nil {
		return "", err
	}
	if _, err = os.Stat(path); err == nil {
		var payload []byte
		if payload, err = ioutil.ReadFile(path); err == nil {
			token = string(payload)
		}
		return token, err
	}
	return "", err
}

func (l LocalPath) Store(token string, vaultAddr *url.URL) (err error) {
	var path string
	if path, err = getPath(l.TokenDirectory, vaultAddr); err != nil {
		return err
	}
	if err = localStorage(l.TokenDirectory); err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(token), 0600)
}

func (l LocalPath) Erase(vaultAddr *url.URL) (err error) {
	var path string
	if path, err = getPath(l.TokenDirectory, vaultAddr); err != nil {
		return err
	}
	if err = localStorage(l.TokenDirectory); err != nil {
		return err
	}
	return os.Remove(path)
}

var _ handler.Handler = (*LocalPath)(nil)
