package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

type goPersonImpl struct {
	fakeit *gofakeit.Faker
}

var _ goPerson = (*goPersonImpl)(nil)

func (self *goPersonImpl) firstNameMethod() (string, error) {
	return self.fakeit.FirstName(), nil
}

func (self *goPersonImpl) lastNameMethod() (string, error) {
	return self.fakeit.LastName(), nil
}

func (self *goPersonImpl) prefixMethod() (string, error) {
	return self.fakeit.NamePrefix(), nil
}

func (self *goPersonImpl) suffixMethod() (string, error) {
	return self.fakeit.NameSuffix(), nil
}

func (self *goPersonImpl) sexTypeMethod() (string, error) {
	return self.fakeit.Gender(), nil
}

func (self *goPersonImpl) jobTitleMethod() (string, error) {
	return self.fakeit.JobTitle(), nil
}

func (self *goPersonImpl) jobLevelMethod() (string, error) {
	return self.fakeit.JobLevel(), nil
}

func (self *goPersonImpl) jobDescriptorMethod() (string, error) {
	return self.fakeit.JobDescriptor(), nil
}
