package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

type goFakerImpl struct {
	fakeit *gofakeit.Faker
	person goPerson
}

var _ goFaker = (*goFakerImpl)(nil)

func (self *goFakerImpl) personGetter() (goPerson, error) {
	return self.person, nil
}
