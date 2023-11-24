package hitchhiker

import "go.k6.io/k6/js/modules"

func init() { register(newModule) }

func newModule(_ modules.VU) goModule {
	return &goModuleImpl{defaultGuide: &goGuideImpl{question: "What's up?"}}
}

type goModuleImpl struct{ defaultGuide goGuide }

func (self *goModuleImpl) newGuide(questionArg string) (goGuide, error) {
	return &goGuideImpl{question: questionArg}, nil
}

func (self *goModuleImpl) checkMethod(valueArg int) (bool, error) {
	return self.defaultGuide.checkMethod(valueArg)
}

func (self *goModuleImpl) defaultGuideGetter() (goGuide, error) { return self.defaultGuide, nil }

type goGuideImpl struct{ question string }

func (self *goGuideImpl) checkMethod(valueArg int) (bool, error) {
	return valueArg == 42, nil
}

func (self *goGuideImpl) questionGetter() (string, error) { return self.question, nil }

func (self *goGuideImpl) questionSetter(questionArg string) error {
	self.question = questionArg
	return nil
}

func (self *goGuideImpl) answerGetter() (int, error) {
	return 42, nil
}
