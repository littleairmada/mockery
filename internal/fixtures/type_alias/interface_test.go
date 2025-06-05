package type_alias_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeAlias(t *testing.T) {
	for _, tt := range []struct {
		name          string
		filepath      string
		expectedRegex string
	}{
		{
			name:          "With alias unresolved",
			filepath:      "./mocks_testify_type_alias_test.go",
			expectedRegex: `func \(_mock \*MockInterface1\) Foo\(\) Type {`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			regex, err := regexp.Compile(tt.expectedRegex)
			require.NoError(t, err)
			bytes, err := os.ReadFile(tt.filepath)
			require.NoError(t, err)

			assert.True(t, regex.Match(bytes), "expected regex was not found in file")
		})
	}
}
