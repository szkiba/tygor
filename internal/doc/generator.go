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

type templateData struct {
	*idl.Module

	Namespace  *idl.Declaration
	Classes    idl.Declarations
	Interfaces idl.Declarations
	Variables  idl.Declarations
	Functions  idl.Declarations

	GitHub *github
}

type github struct {
	Repo      string
	Packages  bool
	Releases  bool
	Examples  bool
	RepoName  string
	RepoOwner string
}

func newTemplateData(mod *idl.Module, opts *options) *templateData {
	data := new(templateData)

	data.Module = mod

	data.GitHub = &github{
		Repo:     opts.githubRepo,
		Releases: opts.linkReleases,
		Packages: opts.linkPackages,
		Examples: opts.linkExamples,
	}

	flds := strings.Split(opts.githubRepo, "/")
	if len(flds) > 1 {
		data.GitHub.RepoOwner = flds[0]
		data.GitHub.RepoName = flds[1]
	}

	for _, dec := range mod.Declarations {
		switch dec.Kind {
		case idl.KindNamespace:
			data.Namespace = dec
		case idl.KindClass:
			data.Classes = append(data.Classes, dec)
		case idl.KindInterface:
			data.Interfaces = append(data.Interfaces, dec)
		case idl.KindFunction:
			data.Functions = append(data.Functions, dec)
		case idl.KindVariable:
			data.Variables = append(data.Variables, dec)
		default:
		}
	}

	return data
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

	if e := t.Execute(&buff, newTemplateData(mod, opts)); e != nil {
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
