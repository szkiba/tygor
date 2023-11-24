package gen

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/szkiba/tygor/internal/idl"
)

func rem(g *jen.Group, format string, args ...any) {
	g.Add(jen.Comment(fmt.Sprintf(format, args...)))
}

func tsdoc(g *jen.Group, src *idl.Declaration) {
	if len(src.Doc) == 0 {
		return
	}

	g.Add(jen.Comment(""))
	g.Add(jen.Comment("TSDoc:"))

	for _, line := range strings.Split(src.Doc, "\n") {
		g.Add(jen.Comment(line))
	}
}

func vmParam() *jen.Statement {
	return jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")
}

func valueType() *jen.Statement {
	return jen.Qual(gojaPath, "Value")
}

func errorType() *jen.Statement {
	return jen.Id("error")
}

func callParam() *jen.Statement {
	return jen.Id("call").Qual(gojaPath, "FunctionCall")
}

func jsFunctionParams(g *jen.Group) {
	g.Add(callParam())
	g.Add(vmParam())
}

func jsFunctionReturns(g *jen.Group) {
	g.Add(valueType())
}

func gojaUndefined() *jen.Statement {
	return jen.Qual(gojaPath, "Undefined").Call()
}

func gojaFlagTrue() *jen.Statement {
	return jen.Qual(gojaPath, "FLAG_TRUE")
}

func gojaFlagFalse() *jen.Statement {
	return jen.Qual(gojaPath, "FLAG_FALSE")
}

func cloneDeclaration(src *idl.Declaration) *idl.Declaration {
	return &idl.Declaration{
		Name:         src.Name,
		Doc:          src.Doc,
		Kind:         src.Kind,
		Type:         src.Type,
		Modifiers:    src.Modifiers,
		Methods:      src.Methods,
		Properties:   src.Properties,
		Constructors: src.Constructors,
		Tags:         src.Tags,
		Parameters:   src.Parameters,
	}
}
