package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lorem(t *testing.T) goLorem {
	t.Helper()

	return &goLoremImpl{fakeit: fakeit(t)}
}

func Test_goLorem(t *testing.T) {
	t.Parallel()
	must := noError(t)

	assert.Equal(
		t,
		"It regularly hourly stairs. Stack poorly twist troop.",
		must(lorem(t).paragraphMethod(1, 2, 4, "\n")),
	)
	assert.Equal(t, "It regularly hourly stairs.", must(lorem(t).sentenceMethod(4)))
	assert.Equal(t, "it", must(lorem(t).wordMethod()))
	assert.Equal(
		t,
		"Forage pinterest direct trade pug skateboard food truck flannel cold-pressed?",
		must(lorem(t).questionMethod()),
	)
	assert.Equal(
		t,
		"\"Forage pinterest direct trade pug skateboard food truck flannel cold-pressed.\" - Lukas Ledner",
		must(lorem(t).quoteMethod()),
	)
}
