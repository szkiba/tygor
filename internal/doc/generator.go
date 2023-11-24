// Package doc contains documentation generator.
package doc

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/russross/blackfriday"
	"github.com/shurcooL/markdownfmt/markdown"
	"github.com/szkiba/tygor/internal/idl"
)

type module struct {
	*idl.Module

	Namespace  *idl.Declaration
	Classes    idl.Declarations
	Interfaces idl.Declarations
	Variables  idl.Declarations
	Functions  idl.Declarations
}

func newModule(mod *idl.Module) *module {
	self := new(module)

	self.Module = mod

	for _, dec := range mod.Declarations {
		switch dec.Kind {
		case idl.KindNamespace:
			self.Namespace = dec
		case idl.KindClass:
			self.Classes = append(self.Classes, dec)
		case idl.KindInterface:
			self.Interfaces = append(self.Interfaces, dec)
		case idl.KindFunction:
			self.Functions = append(self.Functions, dec)
		case idl.KindVariable:
			self.Variables = append(self.Variables, dec)
		default:
		}
	}

	return self
}

// Generate generates documentation for the module.
func Generate(mod *idl.Module, options ...Option) ([]byte, error) {
	opts := getopts(mod, options...)

	t, err := template.New(opts.templateName).
		Funcs(funcMap(opts)).
		Funcs(sprig.TxtFuncMap()).
		Parse(opts.template)
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer

	if e := t.Execute(&buff, newModule(mod)); e != nil {
		return nil, e
	}

	switch opts.format {
	case FormatHTML:
		return renderHTML(buff.Bytes(), addTitle(opts.outer, mod.Name))
	case FormatMarkdown:
		fallthrough
	default:
		return renderMarkdown(buff.Bytes(), opts.outer)
	}
}

const mdFlags = blackfriday.EXTENSION_TABLES |
	blackfriday.EXTENSION_FENCED_CODE |
	blackfriday.EXTENSION_AUTOLINK |
	blackfriday.EXTENSION_STRIKETHROUGH |
	blackfriday.EXTENSION_FOOTNOTES |
	blackfriday.EXTENSION_DEFINITION_LISTS |
	blackfriday.EXTENSION_AUTO_HEADER_IDS

func addTitle(outer []byte, title string) []byte {
	return bytes.ReplaceAll(
		outer,
		[]byte("<title><!-- --></title>"),
		[]byte(fmt.Sprintf("<title>k6/x/%s</title>", title)),
	)
}

func renderHTML(input []byte, outer []byte) ([]byte, error) {
	renderer := blackfriday.HtmlRenderer(0, "", "")

	inter := blackfriday.Markdown(input, renderer, mdFlags)

	return Inject(outer, injectName, inter)
}

func renderMarkdown(input []byte, outer []byte) ([]byte, error) {
	raw := blackfriday.Markdown(input, markdown.NewRenderer(nil), mdFlags)

	if outer == nil {
		return raw, nil
	}

	return Inject(outer, injectName, raw)
}

func funcMap(opts *options) template.FuncMap {
	format := func() string {
		return opts.format.String()
	}

	return template.FuncMap{
		"example": exampleFunc,
		"h":       headingFunc(opts.heading),
		"format":  format,
		"doc":     docFunc,
	}
}

func exampleFunc(dec *idl.Declaration) string {
	if dec == nil {
		return ""
	}

	all, found := dec.Tags["example"]
	if !found {
		return ""
	}

	return strings.Join(all, "\n")
}

func docFunc(dec *idl.Declaration) string {
	if dec == nil || len(dec.Doc) == 0 {
		return ""
	}

	return reJoinLines.ReplaceAllString(dec.Doc, "$1 $2")
}

func headingFunc(startLevel uint) func(uint) string {
	return func(level uint) string {
		var buff strings.Builder

		for i := 1; i < int(startLevel+level); i++ {
			buff.WriteRune('#')
		}

		return buff.String()
	}
}

var reJoinLines = regexp.MustCompile("(?sm)(\\S)\n(\\S)")

const injectName = "api"
