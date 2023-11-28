package faker

import "github.com/brianvoe/gofakeit/v6"

type goCompanyImpl struct {
	fakeit *gofakeit.Faker
}

var _ goCompany = (*goCompanyImpl)(nil)

func (self *goCompanyImpl) nameMethod() (string, error) {
	return self.fakeit.Company(), nil
}

func (self *goCompanyImpl) suffixMethod() (string, error) {
	return self.fakeit.CompanySuffix(), nil
}

func (self *goCompanyImpl) buzzWordMethod() (string, error) {
	return self.fakeit.BuzzWord(), nil
}

func (self *goCompanyImpl) bsMethod() (string, error) {
	return self.fakeit.BS(), nil
}
