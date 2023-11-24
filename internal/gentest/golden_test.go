package gentest

import (
	"testing"

	"github.com/szkiba/tygor/internal/gen"
)

func TestGolden(t *testing.T) {
	t.Parallel()

	GoldenGlob(t, "testdata/*.d.ts", gen.WithGenerator("tygor"))
}
