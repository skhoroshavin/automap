package internal

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"path"
	"testing"
)

func Test(t *testing.T) {
	testsPath := "_tests"

	tests, err := os.ReadDir(testsPath)
	require.NoError(t, err)

	for _, item := range tests {
		if !item.IsDir() {
			continue
		}

		t.Run(item.Name(), func(t *testing.T) {
			testDir := path.Join(testsPath, item.Name(), "my")

			buf := bytes.Buffer{}
			err := AutoMap(&buf, testDir)
			require.NoError(t, err)

			expectedFile, err := os.Open(path.Join(testDir, "automap_gen.go"))
			require.NoError(t, err)

			expectedData, err := io.ReadAll(expectedFile)
			require.NoError(t, err)

			assert.Equal(t, string(expectedData), buf.String())
		})
	}
}
