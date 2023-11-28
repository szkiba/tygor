package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

type goHipsterImpl struct {
	fakeit *gofakeit.Faker
}

var _ goHipster = (*goHipsterImpl)(nil)

func (self *goHipsterImpl) wordMethod() (string, error) {
	return self.fakeit.HipsterWord(), nil
}

func (self *goHipsterImpl) sentenceMethod(wordCountArg int) (string, error) {
	return self.fakeit.HipsterSentence(wordCountArg), nil
}

func (self *goHipsterImpl) paragraphMethod(
	paragraphCountArg int,
	sentenceCountArg int,
	wordCountArg int,
	separatorArg string,
) (string, error) {
	return self.fakeit.HipsterParagraph(
		paragraphCountArg,
		sentenceCountArg,
		wordCountArg,
		separatorArg,
	), nil
}
