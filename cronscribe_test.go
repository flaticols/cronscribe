package cronscribe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCronscribe(t *testing.T) {
	sc, err := New("./pkg/core/rules")
	require.NoError(t, err)

	err = sc.SetLanguage("en")
	require.NoError(t, err)

	exp, err := sc.Convert("every first monday of month")
	require.NoError(t, err)

	require.Equal(t, "0 0 * * 1#1", exp)
}
