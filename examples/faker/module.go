package faker

import (
	"github.com/brianvoe/gofakeit/v6"
	"go.k6.io/k6/js/modules"
	"lukechampine.com/frand"
)

func init() {
	register(newModule)
}

type goModuleImpl struct {
	faker goFaker
}

var _ goModule = (*goModuleImpl)(nil)

func newModule(_ modules.VU) goModule {
	mod := new(goModuleImpl)
	faker, _ := mod.newFaker(0)
	mod.faker = faker

	return mod
}

func (self *goModuleImpl) newFaker(seedArg int64) (goFaker, error) {
	src := frand.NewSource()

	if seedArg != 0 {
		src.Seed(seedArg)
	}

	fakeit := gofakeit.NewCustom(src)

	return &goFakerImpl{
		fakeit: fakeit,
		person: &goPersonImpl{fakeit: fakeit},
	}, nil
}

func (self *goModuleImpl) fakerGetter() (goFaker, error) {
	return self.faker, nil
}
