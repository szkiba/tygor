package state

import (
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/lib"
)

type goModuleImpl struct{ vu modules.VU }

func (self *goModuleImpl) activeVUsGetter() (int64, error) {
	return lib.GetExecutionState(self.vu.Context()).GetCurrentlyActiveVUsCount(), nil
}

func (self *goModuleImpl) iterationGetter() (int64, error) {
	return self.vu.State().Iteration, nil
}

func (self *goModuleImpl) vuIDGetter() (uint64, error) {
	return self.vu.State().VUID, nil
}

func (self *goModuleImpl) vuIDFromRuntimeGetter() (int64, error) {
	return self.vu.Runtime().Get("__VU").ToInteger(), nil
}

func init() {
	register(newModule)
}

func newModule(vu modules.VU) goModule {
	return &goModuleImpl{vu}
}
