package localpath

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

const (
	localStoragePerm = os.FileMode(0750)
)

func localStorage(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, localStoragePerm)
	}
	return err
}

func getPath(path string, vaultAddr *url.URL) (p string, err error) {
	return filepath.Join(path, id(vaultAddr)), nil
}

func id(vaultAddr *url.URL) string {
	var id = fmt.Sprintf("%s%s", vaultAddr.Hostname(), vaultAddr.RequestURI())
	return fmt.Sprintf("%x", md5.Sum([]byte(id)))
}
