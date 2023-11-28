package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hipster(t *testing.T) goHipster {
	t.Helper()

	return &goHipsterImpl{fakeit: fakeit(t)}
}

func Test_goHipster(t *testing.T) {
	t.Parallel()
	must := noError(t)

	assert.Equal(t, "offal", must(hipster(t).wordMethod()))
	assert.Equal(t, "Offal forage pinterest direct trade.", must(hipster(t).sentenceMethod(4)))
	assert.Equal(
		t,
		"Offal forage pinterest direct trade. Pug skateboard food truck flannel.",
		must(hipster(t).paragraphMethod(1, 2, 4, "\n")),
	)
}
