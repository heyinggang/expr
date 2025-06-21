package issue_test

import (
	"testing"

	"github.com/heyinggang/expr"
	"github.com/heyinggang/expr/internal/testify/require"
)

func TestIssue739(t *testing.T) {
	jsonString := `{"Num": 1}`
	env := map[string]any{
		"aJSONString": &jsonString,
	}

	result, err := expr.Eval("fromJSON(aJSONString)", env)
	require.NoError(t, err)
	require.Contains(t, result, "Num")
}
