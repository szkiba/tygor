package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func company(t *testing.T) goCompany {
	t.Helper()

	return &goCompanyImpl{fakeit: fakeit(t)}
}

func Test_goCompany(t *testing.T) {
	t.Parallel()
	must := noError(t)

	assert.Equal(t, "Xatori", must(company(t).nameMethod()))
	assert.Equal(t, "LLC", must(company(t).suffixMethod()))
	assert.Equal(t, "Reverse-engineered", must(company(t).buzzWordMethod()))
	assert.Equal(t, "24-7", must(company(t).bsMethod()))
}
