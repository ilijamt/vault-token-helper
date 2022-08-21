package resources_test

import (
	"github.com/ilijamt/vault-token-helper/internal/handler/localpath"
	"github.com/ilijamt/vault-token-helper/internal/resources"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandlers(t *testing.T) {
	require.NotEmpty(t, resources.Available())

	lp, err := resources.Get(localpath.NameLocalPath)
	require.NotNil(t, lp)
	require.NoError(t, err)

	_, err = resources.Get("unknown")
	require.Error(t, err)
}
