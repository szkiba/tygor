package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	xdoc "github.com/szkiba/tygor/internal/doc"

	"github.com/spf13/cobra/doc"
	"github.com/stretchr/testify/assert"
)

func TestDocs(t *testing.T) {
	t.Parallel()

	if !*update {
		return
	}

	var buff bytes.Buffer

	root := rootCmd()

	linkHandler := func(name string) string {
		link := strings.ReplaceAll(strings.TrimSuffix(name, ".md"), "_", "-")
		return "#" + link
	}

	assert.NoError(t, doc.GenMarkdownCustom(root, &buff, linkHandler))

	for _, cmd := range root.Commands() {
		_, err := (&buff).WriteString("---\n")
		assert.NoError(t, err)
		assert.NoError(t, doc.GenMarkdownCustom(cmd, &buff, linkHandler))
	}

	readme := filepath.Clean(filepath.Join("..", "..", "README.md"))

	src, err := os.ReadFile(readme) //nolint:forbidigo
	assert.NoError(t, err)

	res, err := xdoc.Inject(src, "cli", buff.Bytes())
	assert.NoError(t, err)

	assert.NoError(t, os.WriteFile(readme, res, 0o600)) //nolint:forbidigo
}
