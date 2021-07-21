package workaround

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	slice := []string{"a", "b", "c"}

	formatted := FormatContentTypes(slice)

	require.Equal(t, "a,b,c", formatted)
}
