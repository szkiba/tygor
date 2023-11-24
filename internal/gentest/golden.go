// Package gentest contains generator test helpers.
package gentest

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szkiba/tygor/internal/cmd"
	"github.com/szkiba/tygor/internal/doc"
	"github.com/szkiba/tygor/internal/gen"
	"github.com/szkiba/tygor/internal/idl"
)

var update = flag.Bool("update", false, "update golden files") //nolint:gochecknoglobals

func readTextFile(t *testing.T, filename string) string {
	t.Helper()

	source, err := os.ReadFile(filename) //nolint:forbidigo,gosec

	assert.NoError(t, err)

	// git + windows workaround
	return string(bytes.ReplaceAll(source, []byte{'\r', '\n'}, []byte{'\n'}))
}

func golden(t *testing.T, filename string, content string) {
	t.Helper()

	if *update {
		assert.NoError(t, os.WriteFile(filename, []byte(content), 0o600)) //nolint:forbidigo
	}

	expected := readTextFile(t, filename)

	assert.Equal(t, expected, content)
}

// Golden test code generator using golden files. Golden files can be updated using -update CLI flag.
func Golden(t *testing.T, filename string, opts ...gen.Option) {
	t.Helper()

	opts = append(opts, gen.WithPackage(filepath.Dir(filename)))

	mod, err := idl.Parse(filename, readTextFile(t, filename))

	assert.NoError(t, err)

	basename := strings.TrimSuffix(filename, ".d.ts")

	golden(t, basename+cmd.CodeSuffix+".go", gen.Code(mod, opts...).GoString())

	opts = append(opts, gen.WithTag(cmd.SketchTag))
	golden(t, basename+cmd.SketchSuffix+".go", gen.Sketch(mod, opts...).GoString())
}

// GoldenGlob calls Golden on files matching to the glob pattern.
func GoldenGlob(t *testing.T, pattern string, opts ...gen.Option) {
	t.Helper()

	files, err := filepath.Glob(pattern)
	assert.NoError(t, err)

	for _, filename := range files {
		filename := filename

		t.Run(filename, func(t *testing.T) {
			t.Parallel()
			Golden(t, filename, opts...)
		})
	}
}

// GoldenDoc test doc generator using golden files. Golden files can be updated using -update CLI flag.
func GoldenDoc(t *testing.T, filename string) {
	t.Helper()

	dir := filepath.Dir(filename)

	mod, err := idl.Parse(filename, readTextFile(t, filename))

	assert.NoError(t, err)

	must := func(data []byte, err error) string {
		if err != nil {
			t.Error(err)
		}

		// windows workaround
		return string(bytes.ReplaceAll(data, []byte{'\r', '\n'}, []byte{'\n'}))
	}

	golden(t, filepath.Join(dir, "README.md"), must(doc.Generate(mod)))
	golden(t, filepath.Join(dir, "index.html"),
		must(doc.Generate(mod, doc.WithFormat(doc.FormatHTML))),
	)
}

// GoldenGlobDoc calls GoldenDoc on files matching to the glob pattern.
func GoldenGlobDoc(t *testing.T, pattern string) {
	t.Helper()

	files, err := filepath.Glob(pattern)
	assert.NoError(t, err)

	for _, filename := range files {
		filename := filename

		t.Run(filename, func(t *testing.T) {
			t.Parallel()
			GoldenDoc(t, filename)
		})
	}
}
