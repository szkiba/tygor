package gen

import "github.com/szkiba/tygor/internal/idl"

type options struct {
	pkg       string
	generator string
	tag       string
}

// Option defines a code generator option function.
type Option func(*options)

// WithPackage can be used to specify go package name for generated code.
func WithPackage(name string) Option {
	return func(o *options) {
		o.pkg = name
	}
}

// WithGenerator can be used to specify the generator name for generated code.
// Generator name will appear in the file header comment.
func WithGenerator(name string) Option {
	return func(o *options) {
		o.generator = name
	}
}

// WithTag can be used to specify go build tag for generated code.
func WithTag(name string) Option {
	return func(o *options) {
		o.tag = name
	}
}

func getopts(mod *idl.Module, args ...Option) *options {
	opts := new(options)

	opts.pkg = mod.Name

	for _, fn := range args {
		fn(opts)
	}

	return opts
}
