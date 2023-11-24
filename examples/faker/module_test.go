package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func noError(t *testing.T) func(interface{}, error) interface{} {
	t.Helper()

	return func(i interface{}, err error) interface{} {
		assert.NoError(t, err)

		return i
	}
}

func Test_newModule(t *testing.T) {
	t.Parallel()

	must := noError(t)
	mod := newModule(nil)

	assert.NotNil(t, mod)

	faker, err := mod.fakerGetter()

	assert.NoError(t, err)
	assert.NotNil(t, faker)
	assert.NotNil(t, must(faker.personGetter()))

	faker, err = mod.newFaker(0)

	assert.NoError(t, err)
	assert.NotNil(t, faker)
	assert.NotNil(t, must(faker.personGetter()))
}
