package faker

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"lukechampine.com/frand"
)

const sampleSeed = 11

func fakeit(t *testing.T) *gofakeit.Faker {
	t.Helper()

	src := frand.NewSource()

	src.Seed(sampleSeed)

	return gofakeit.NewCustom(src)
}

func Test_goFaker(t *testing.T) {
	t.Parallel()

	must := noError(t)

	faker, err := newModule(nil).newFaker(sampleSeed)

	assert.NoError(t, err)
	assert.NotNil(t, faker)

	assert.NotNil(t, must(faker.personGetter()))
}
