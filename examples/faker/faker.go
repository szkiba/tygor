package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

type goFakerImpl struct {
	fakeit  *gofakeit.Faker
	person  goPerson
	company goCompany
	lorem   goLorem
	hacker  goHacker
	hipster goHipster
}

var _ goFaker = (*goFakerImpl)(nil)

func (self *goFakerImpl) personGetter() (goPerson, error) {
	return self.person, nil
}

func (self *goFakerImpl) companyGetter() (goCompany, error) {
	return self.company, nil
}

func (self *goFakerImpl) loremGetter() (goLorem, error) {
	return self.lorem, nil
}

func (self *goFakerImpl) hackerGetter() (goHacker, error) {
	return self.hacker, nil
}

func (self *goFakerImpl) hipsterGetter() (goHipster, error) {
	return self.hipster, nil
}
