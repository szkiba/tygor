package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

type goHackerImpl struct {
	fakeit *gofakeit.Faker
}

var _ goHacker = (*goHackerImpl)(nil)

func (self *goHackerImpl) abbreviationMethod() (string, error) {
	return self.fakeit.HackerAbbreviation(), nil
}

func (self *goHackerImpl) adjectiveMethod() (string, error) {
	return self.fakeit.HackerAdjective(), nil
}

func (self *goHackerImpl) ingverbMethod() (string, error) {
	return self.fakeit.HackeringVerb(), nil
}

func (self *goHackerImpl) nounMethod() (string, error) {
	return self.fakeit.HackerNoun(), nil
}

func (self *goHackerImpl) phraseMethod() (string, error) {
	return self.fakeit.HackerPhrase(), nil
}

func (self *goHackerImpl) verbMethod() (string, error) {
	return self.fakeit.HackerVerb(), nil
}
