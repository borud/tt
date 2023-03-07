package sqlitestore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqliteStore(t *testing.T) {
	store, err := New(":memory:")
	require.NoError(t, err)
	require.NoError(t, store.Close())
}
