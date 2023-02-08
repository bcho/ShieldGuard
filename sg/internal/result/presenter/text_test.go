package presenter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Text(t *testing.T) {
	presenter := Text(testQueryResults())
	output := new(bytes.Buffer)
	err := presenter.WriteQueryResultTo(output)
	assert.NoError(t, err)
	t.Log("\n" + output.String())
	assert.Equal(
		t,
		`FAIL - file name - fail message1
Document: https://github.com/Azure/ShieldGuard/docs/001-rego.md
FAIL - file name - fail message2
WARN - file name - warn message1
WARN - file name - warn message2
Document: https://github.com/Azure/ShieldGuard/docs/002-rego.md
EXCEPTION - file name - exception message1
7 test(s), 2 passed, 2 failure(s) 2 warning(s), 1 exception(s)
`,
		output.String(),
	)
}