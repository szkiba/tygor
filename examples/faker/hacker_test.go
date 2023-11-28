package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hacker(t *testing.T) goHacker {
	t.Helper()

	return &goHackerImpl{fakeit: fakeit(t)}
}

func Test_goHacker(t *testing.T) {
	t.Parallel()
	must := noError(t)

	assert.Equal(t, "GB", must(hacker(t).abbreviationMethod()))
	assert.Equal(t, "auxiliary", must(hacker(t).adjectiveMethod()))
	assert.Equal(t, "quantifying", must(hacker(t).ingverbMethod()))
	assert.Equal(t, "application", must(hacker(t).nounMethod()))
	assert.Equal(
		t,
		"Try to transpile the EXE sensor, maybe it will deconstruct the wireless interface!",
		must(hacker(t).phraseMethod()),
	)
	assert.Equal(t, "read", must(hacker(t).verbMethod()))
}
