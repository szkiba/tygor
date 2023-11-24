package toml

import (
	"github.com/pelletier/go-toml"
	"go.k6.io/k6/js/modules"
)

type goModuleImpl struct{}

func (self *goModuleImpl) parseMethod(textArg string) (interface{}, error) {
	obj := map[string]interface{}{}

	err := toml.Unmarshal([]byte(textArg), &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (self *goModuleImpl) stringifyMethod(valueArg interface{}) (string, error) {
	b, err := toml.Marshal(valueArg)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func init() {
	register(newModule)
}

func newModule(_ modules.VU) goModule {
	return new(goModuleImpl)
}
