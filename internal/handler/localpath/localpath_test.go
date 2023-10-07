package localpath_test

import (
	"github.com/ilijamt/vault-token-helper/internal/handler"
	"github.com/ilijamt/vault-token-helper/internal/handler/localpath"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"testing"
)

func TestLocalPathHomeDir(t *testing.T) {
	h := localpath.New()
	require.Implements(t, (*handler.Handler)(nil), h)
}

func TestLocalPath(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, os.RemoveAll(tmpDir))
	}()

	h := localpath.NewWithPath(tmpDir)

	require.NotEmpty(t, h.Name())
	require.NotEmpty(t, h.Description())
	require.NotEmpty(t, h.Path())

	t.Run("invalid vault url", func(t *testing.T) {
		require.ErrorIs(t, h.Erase(nil), handler.ErrNoVaultAddress)
		require.ErrorIs(t, h.Store("token", nil), handler.ErrNoVaultAddress)
		_, err := h.Get(nil)
		require.ErrorIs(t, err, handler.ErrNoVaultAddress)
	})

	t.Run("full rundown", func(t *testing.T) {
		u, err := url.Parse("http://localhost:1234")
		require.NoError(t, err)

		// No token
		token, err := h.Get(u)
		require.Error(t, err)
		require.Empty(t, token)

		// Store a token
		require.NoError(t, h.Store("test", u))

		// Get token
		token, err = h.Get(u)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		var items []handler.Data
		items, err = h.List()
		require.Len(t, items, 1)
		require.NoError(t, err)

		items, err = h.Purge(true)
		require.Len(t, items, 1)
		require.NoError(t, err)

		// Remove the token
		require.NoError(t, h.Erase(u))

		// Store a token
		require.NoError(t, h.Store("test", u))

		items, err = h.Purge(false)
		require.Len(t, items, 1)
		require.NoError(t, err)

		// Remove the token
		require.Error(t, h.Erase(u))

		// No token
		token, err = h.Get(u)
		require.Error(t, err)
		require.Empty(t, token)
	})

}
