package issue_test

import (
	"testing"

	"github.com/heyinggang/expr"
	"github.com/heyinggang/expr/internal/testify/require"
)

func TestIssue723(t *testing.T) {
	emptyMap := make(map[string]string)
	env := map[string]interface{}{
		"empty_map": emptyMap,
	}

	code := `get(empty_map, "non_existing_key")`

	program, err := expr.Compile(code, expr.Env(env))
	require.NoError(t, err)

	output, err := expr.Run(program, env)
	require.NoError(t, err)
	require.Equal(t, nil, output)
}
