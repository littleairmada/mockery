package recursivegenerationwithsubpkgexclude_test

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vektra/mockery/v3/internal/file"
)

func TestSubpkg2NotExist(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	subpkg2MockFile := path.Join(wd, "subpkg2", "mocks.go")
	exists, err := file.Exists(subpkg2MockFile)
	require.NoError(t, err)
	assert.False(t, exists, "subpkg2 mocks.go file exists when it shouldn't")
}

func TestSubpkg1Exists(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	subpkg2MockFile := path.Join(wd, "subpkg1", "mocks.go")
	exists, err := file.Exists(subpkg2MockFile)
	require.NoError(t, err)
	assert.True(t, exists, "subpkg1 mocks.go file doesn't exist when it should")
}
