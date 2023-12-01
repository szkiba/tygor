package doc

import (
	_ "embed"

	"github.com/szkiba/tygor/internal/idl"
)

type options struct {
	format       Format
	template     string
	templateName string
	outer        []byte
	heading      uint

	githubRepo   string
	linkReleases bool
	linkPackages bool
	linkExamples bool
}

// Option defines a documentation generator option function.
type Option func(*options)

// WithTemplate can be used to specify go template source.
func WithTemplate(tmpl string) Option {
	return func(o *options) {
		o.template = tmpl
	}
}

// WithTemplateName can be used to specify go template name.
func WithTemplateName(name string) Option {
	return func(o *options) {
		o.templateName = name
	}
}

// WithFormat can be used to specify the output format.
func WithFormat(format Format) Option {
	return func(o *options) {
		o.format = format
	}
}

// WithOuter can be used to specify the outer document to update.
func WithOuter(outer []byte) Option {
	return func(o *options) {
		o.outer = outer
	}
}

// WithHeading can be used to specify the initial heading level.
func WithHeading(heading uint) Option {
	return func(o *options) {
		o.heading = heading
	}
}

// WithGitHubRepo can be used to specify the GitHub repository in owner/name form.
func WithGitHubRepo(repo string) Option {
	return func(o *options) {
		o.githubRepo = repo
	}
}

// WithLinkReleases can be used to enable/disable GitHub releases link.
func WithLinkReleases(flag bool) Option {
	return func(o *options) {
		o.linkReleases = flag
	}
}

// WithLinkPackages can be used to enable/disable GitHub container packages link.
func WithLinkPackages(flag bool) Option {
	return func(o *options) {
		o.linkPackages = flag
	}
}

// WithLinkExamples can be used to enable/disable examples folder link.
func WithLinkExamples(flag bool) Option {
	return func(o *options) {
		o.linkExamples = flag
	}
}

func getopts(_ *idl.Module, args ...Option) *options {
	opts := new(options)

	opts.template = defaultTemplate
	opts.templateName = defaultTemplateName
	opts.format = defaultFormat
	opts.heading = defaultHeading

	for _, fn := range args {
		fn(opts)
	}

	if opts.format == FormatHTML && opts.outer == nil {
		opts.outer = defaultOuterHTML
	}

	return opts
}

//go:embed doc.gtpl
var defaultTemplate string

//go:embed outer.html
var defaultOuterHTML []byte

const (
	defaultFormat            = FormatMarkdown
	defaultTemplateName      = "doc.gtpl"
	defaultHeading      uint = 1
)
