package cmd

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestUsage(t *testing.T) {
	t.Parallel()

	dir := "testdata"

	root := rootCmd()

	golden(t, root, filepath.Join(dir, "root.txt"))

	for _, cmd := range root.Commands() {
		golden(t, cmd, filepath.Join(dir, cmd.Name()+".txt"))
	}
}

var update = flag.Bool("update", false, "update golden files") //nolint:gochecknoglobals

func readTextFile(t *testing.T, filename string) []byte {
	t.Helper()

	source, err := os.ReadFile(filename) //nolint:forbidigo,gosec

	assert.NoError(t, err)

	// git + windows workaround
	return bytes.ReplaceAll(source, []byte{'\r', '\n'}, []byte{'\n'})
}

func golden(t *testing.T, cmd *cobra.Command, filename string) {
	t.Helper()

	if *update {
		var buff bytes.Buffer

		cmd.SetOut(&buff)
		assert.NoError(t, cmd.Usage())

		assert.NoError(t, os.WriteFile(filename, buff.Bytes(), 0o600)) //nolint:forbidigo
	}

	expected := readTextFile(t, filename)

	var buff bytes.Buffer

	cmd.SetOut(&buff)
	assert.NoError(t, cmd.Usage())

	assert.Equal(t, string(expected), buff.String())
}
