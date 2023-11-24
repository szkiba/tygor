package idl

import (
	_ "embed"

	"github.com/dop251/goja"
)

type extractFunc func(filename, source string) (string, error)

func extract(filename, source string) (string, error) {
	fn, err := newExtractFunc(extractorScriptName, extractorScript, newRuntime())
	if err != nil {
		return "", err
	}

	return fn(filename, source)
}

func newExtractFunc(scriptName, script string, runtime *goja.Runtime) (extractFunc, error) {
	var err error

	var prog *goja.Program

	prog, err = goja.Compile(scriptName, script, true)
	if err != nil {
		return nil, err
	}

	module := runtime.NewObject()
	exports := runtime.NewObject()

	if err = module.Set(propExports, exports); err != nil {
		return nil, err
	}

	if err = runtime.Set(propModule, module); err != nil {
		return nil, err
	}

	_, err = runtime.RunProgram(prog)
	if err != nil {
		return nil, err
	}

	obj := module.Get(propExports).ToObject(runtime)

	value := obj.Get(propDefault)

	var fn extractFunc

	if err = runtime.ExportTo(value, &fn); err != nil {
		return nil, err
	}

	return fn, nil
}

//nolint:errcheck,gosec
func newRuntime() *goja.Runtime {
	runtime := goja.New()

	console := runtime.NewObject()

	noop := func(...goja.Value) {}

	console.Set("debug", noop)
	console.Set("log", noop)
	console.Set("info", noop)
	console.Set("warn", noop)
	console.Set("error", noop)

	runtime.Set("console", console)

	return runtime
}

const (
	propModule  = "module"
	propExports = "exports"
	propDefault = "default"

	extractorScriptName = "extractor.cjs"
)

//go:embed extractor/dist/extractor.cjs
var extractorScript string
