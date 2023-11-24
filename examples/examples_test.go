package examples

import (
	"testing"

	"github.com/szkiba/tygor/internal/gen"
	"github.com/szkiba/tygor/internal/gentest"
	"github.com/szkiba/tygor/internal/xk6it"
)

func TestGolden(t *testing.T) {
	t.Parallel()

	gentest.GoldenGlob(t, "*/*.d.ts", gen.WithGenerator("tygor"))
	gentest.GoldenGlobDoc(t, "*/*.d.ts")
}

func TestScript(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	xk6it.Run(t, "*/test.js")
}
