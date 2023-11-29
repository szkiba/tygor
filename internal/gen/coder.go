package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/szkiba/tygor/internal/idl"
)

type coder struct {
	src    *idl.Declaration
	naming naming
}

func newCoder(src *idl.Declaration, naming naming) *coder {
	return &coder{src: src, naming: naming}
}

func (c *coder) code(out *jen.Group) {
	switch c.src.Kind {
	case idl.KindInterface:
		c.interfaceCode(out)
	case idl.KindClass:
		c.classCode(out)
	case idl.KindNamespace:
		c.namespaceCode(out)
	default:
	}
}

func (c *coder) sketch(out *jen.Group) {
	switch c.src.Kind {
	case idl.KindInterface:
		c.interfaceSketch(out)
	case idl.KindClass:
		c.classSketch(out)
	case idl.KindNamespace:
		c.namespaceSketch(out)
	default:
	}
}

func (c *coder) interfaceCode(out *jen.Group) {
	c.jsInterfaceCode(out)
	c.goInterfaceCode(out)
	c.jsAdapterCode(out)
	c.goAdapterCode(out)
	c.jsToCode(out)
	c.jsFromCode(out)
	c.goFromCode(out)
	c.goToObjectCode(out)
}

func (c *coder) interfaceSketch(out *jen.Group) {
	c.goInterfaceSketch(out)
}

func (c *coder) classCode(out *jen.Group) {
	c.jsInterfaceCode(out)
	c.goInterfaceCode(out)
	c.jsAdapterCode(out)
	c.goAdapterCode(out)
	c.jsToCode(out)
	c.jsFromCode(out)
	c.goFromCode(out)
	c.goToObjectCode(out)
	c.jsConstructorCode(out)
}

func (c *coder) classSketch(out *jen.Group) {
	c.goInterfaceSketch(out)
}

func (c *coder) foreach(decls idl.Declarations, visitor func(*coder)) {
	for _, decl := range decls {
		visitor(newCoder(decl, c.naming))
	}
}

func (c *coder) as(what nameable) string {
	return c.naming(what, c.src.Name)
}

func (c *coder) goName() string {
	switch c.src.Kind {
	case idl.KindMethod, idl.KindFunction:
		return c.as(goMethod)
	case idl.KindProperty, idl.KindVariable:
		return c.as(goGetter)
	case idl.KindClass, idl.KindInterface:
		return c.as(goInterface)
	case idl.KindConstructor:
		return c.as(goFactory)
	default:
		return c.src.Name
	}
}

func (c *coder) asName() *jen.Statement {
	return jen.Id(c.goName())
}

func (c *coder) asType() *jen.Statement {
	switch c.src.Type {
	case "void":
		return jen.Null()
	case "number":
		return jen.Id("float64")
	case "boolean":
		return jen.Id("bool")
	case "string":
		return jen.Id("string")

	case "int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64",
		"byte",
		"rune",
		"float32",
		"float64",
		"int", "uint":
		return jen.Id(c.src.Type)

	case "any", "object":
		return jen.Id("interface{}")

	case "ArrayBuffer":
		return jen.Id("[]byte")

	case "Date":
		return jen.Qual(timePath, "Time")

	default:
		return jen.Id("go" + c.src.Type)
	}
}

func (c *coder) asValueToCall(value *jen.Statement, vm *jen.Statement) *jen.Statement {
	switch c.src.Type {
	case "number", "float64":
		return value.Dot("ToFloat").Call()
	case "boolean":
		return value.Dot("ToBoolean").Call()
	case "string":
		return value.Dot("String").Call()

	case "int64":
		return value.Dot("ToInteger").Call()

	case "int8", "int16", "int32", "int",
		"uint8", "uint16", "uint32", "uint64", "uint",
		"byte", "rune":
		return jen.Id(c.src.Type).Call(value.Dot("ToInteger").Call())

	case "float32":
		return jen.Id(c.src.Type).Call(value.Dot("ToFloat").Call())

	case "any", "object":
		return value.Dot("Export").Call()

	case "void":
		return nil

	case "ArrayBuffer":
		return jen.Parens(value.Dot("Export").
			Call().
			Assert(jen.Qual(gojaPath, "ArrayBuffer"))).
			Dot("Bytes").Call()

	case "Date":
		return jen.Id("goTimeFromDate").Call(value, vm)

	default:
		return value.Dot("Export").Call().Assert(c.asType())
	}
}

func (c *coder) asToValueCall(goval *jen.Statement) *jen.Statement {
	switch c.src.Type {
	case "ArrayBuffer":
		return jen.Id("vm").Dot("ToValue").Call(jen.Id("vm").Dot("NewArrayBuffer").Call(goval))
	case "Date":
		return jen.Id("jsDateFromTime").Call(goval, jen.Id("vm"))

	case "number", "boolean", "string",
		"any", "object",
		"int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64",
		"byte",
		"rune",
		"float32",
		"float64",
		"int", "uint":
		return jen.Id("vm").Dot("ToValue").Call(goval)

	default:
		return jen.Id(c.naming(goInterface, c.src.Type)+"ToObject").Call(goval, jen.Id("vm"))
	}
}

func (c *coder) asArg() *jen.Statement {
	return jen.Id(c.as(goParam))
}

func (c *coder) asParameter() *jen.Statement {
	return c.asArg().Add(c.asType())
}

func (c *coder) parameters() *jen.Statement {
	return jen.ListFunc(func(g *jen.Group) {
		c.foreach(c.src.Parameters, func(p *coder) {
			g.Add(p.asParameter())
		})
	})
}

func (c *coder) asZero() *jen.Statement {
	var value string

	switch c.src.Type {
	case "number":
		value = "0"
	case "boolean":
		value = "false"
	case "string":
		value = `""`
	case "void":
		return nil

	case "int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64",
		"byte",
		"rune",
		"float32",
		"float64",
		"int", "uint":

		value = "0"

	case "Date":
		value = timePath + ".Time{}"

	default:
		value = "nil"
	}

	return jen.Id(value)
}

func (c *coder) constructor() *idl.Declaration {
	if len(c.src.Constructors) == 0 {
		ctor := new(idl.Declaration)
		ctor.Kind = idl.KindConstructor

		return ctor
	}

	return c.src.Constructors[0]
}
