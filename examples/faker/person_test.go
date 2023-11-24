package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func person(t *testing.T) goPerson {
	t.Helper()

	return &goPersonImpl{fakeit: fakeit(t)}
}

func Test_goPerson(t *testing.T) {
	t.Parallel()
	must := noError(t)

	assert.Equal(t, "Josiah", must(person(t).firstNameMethod()))
	assert.Equal(t, "Abshire", must(person(t).lastNameMethod()))
	assert.Equal(t, "Mr.", must(person(t).prefixMethod()))
	assert.Equal(t, "Sr.", must(person(t).suffixMethod()))
	assert.Equal(t, "Representative", must(person(t).jobTitleMethod()))
	assert.Equal(t, "Internal", must(person(t).jobDescriptorMethod()))
	assert.Equal(t, "Identity", must(person(t).jobLevelMethod()))
	assert.Equal(t, "male", must(person(t).sexTypeMethod()))
}
