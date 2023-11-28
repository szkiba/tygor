package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

type goLoremImpl struct {
	fakeit *gofakeit.Faker
}

var _ goLorem = (*goLoremImpl)(nil)

func (self *goLoremImpl) paragraphMethod(
	paragraphCountArg int,
	sentenceCountArg int,
	wordCountArg int,
	separatorArg string,
) (string, error) {
	return self.fakeit.Paragraph(
		paragraphCountArg,
		sentenceCountArg,
		wordCountArg,
		separatorArg,
	), nil
}

func (self *goLoremImpl) sentenceMethod(wordCountArg int) (string, error) {
	return self.fakeit.Sentence(wordCountArg), nil
}

func (self *goLoremImpl) wordMethod() (string, error) {
	return self.fakeit.Word(), nil
}

func (self *goLoremImpl) questionMethod() (string, error) {
	return self.fakeit.Question(), nil
}

func (self *goLoremImpl) quoteMethod() (string, error) {
	return self.fakeit.Quote(), nil
}
