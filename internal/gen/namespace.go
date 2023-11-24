package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/szkiba/tygor/internal/idl"
)

func (c *coder) namespaceCode(out *jen.Group) {
	out.Comment("k6Module represents k6 JavaScript extension module.")
	out.Type().Id("k6Module").Struct(jen.Id("goModuleConstructor").Id("goModuleConstructor"))

	c.newModuleInstanceCode(out)

	out.Comment("register registers k6 JavaScript extension module.")
	out.Func().
		Id("register").
		Params(jen.Id("ctor").Id("goModuleConstructor")).
		Block(jen.Id("m").Op(":=").Id("new").
			Call(jen.Id("k6Module")),
			jen.Id("m").Dot("goModuleConstructor").Op("=").Id("ctor"),
			jen.Id("modules").Dot("Register").Call(jen.Lit("k6/x/"+c.src.Name), jen.Id("m")))

	out.Comment("k6ModuleInstance represents per VU module instance.")
	out.Type().
		Id("k6ModuleInstance").
		Struct(jen.Id("exports").Id("modules").Dot("Exports"))

	out.Comment("Exports returns exported symbols.")
	out.Func().
		Params(jen.Id("mi").Op("*").Id("k6ModuleInstance")).
		Id("Exports").
		Params().
		Params(jen.Id("modules").Dot("Exports")).
		Block(jen.Return().Id("mi").Dot("exports"))

	c.moduleInterfaceCode(out)
}

func (c *coder) toModuleInterface() *coder {
	src := cloneDeclaration(c.src)
	src.Name = "Module"
	src.Kind = idl.KindInterface

	for _, method := range src.Methods {
		if method.Kind == idl.KindConstructor {
			src.Constructors = append(src.Constructors, method)
		}
	}

	return newCoder(src, c.naming)
}

func (c *coder) moduleInterfaceCode(out *jen.Group) {
	src := c.toModuleInterface()

	src.jsInterfaceCode(out)
	src.jsAdapterCode(out)
	src.jsFromCode(out)

	src.goInterfaceCode(out)

	ctor := src.as(goConstructor)

	rem(out, "%s creates new %s instance.", ctor, src.as(goInterface))
	out.Type().
		Id(ctor).Func().
		Params(jen.Id("vu").Qual(modulesPath, "VU")).
		Params(src.asName())
}

func (c *coder) newModuleInstanceCode(out *jen.Group) { //nolint:funlen
	out.Comment("NewModuleInstance creates new per VU module instance.")
	out.Func().
		Params(jen.Id("m").Op("*").Id("k6Module")).
		Id("NewModuleInstance").
		Params(jen.Id("vu").Qual(modulesPath, "VU")).
		Params(jen.Qual(modulesPath, "Instance")).
		BlockFunc(func(g *jen.Group) {
			g.Id("mi").Op(":=").Id("new").Call(jen.Id("k6ModuleInstance"))
			g.Id("adaptee").Op(":=").Id("m").Dot("goModuleConstructor").Call(jen.Id("vu"))

			needAdapter := false

			for _, m := range c.src.Methods {
				if m.Kind != idl.KindConstructor {
					needAdapter = true

					break
				}
			}

			needAdapter = needAdapter || (len(c.src.Properties) != 0)

			var defaultExport *jen.Statement

			c.foreach(c.src.Properties, func(prop *coder) {
				if prop.src.Modifiers.Contains(idl.ModifierDefault) {
					needAdapter = true

					defaultExport = jen.
						Id("adapter").
						Dot(prop.as(jsGetter)).
						Call(
							jen.Qual(gojaPath, "FunctionCall").
								Values(jen.Id("This").Op(":").Qual(gojaPath, "Undefined").Call()),
							jen.Id("vm")).
						Dot("ToObject").
						Call(jen.Id("vm"))
				}
			})

			if needAdapter {
				g.Id("adapter").Op(":=").Id("jsModuleFrom").Call(jen.Id("adaptee"))
			}

			g.Empty()
			g.Id("vm").Op(":=").Id("vu").Dot("Runtime").Call()
			g.Empty()

			g.Id("dict").Op(":=").Id("make").Call(jen.Map(jen.Id("string")).Interface())
			g.Empty()

			c.foreach(c.src.Methods, func(m *coder) {
				var value *jen.Statement

				if m.src.Kind == idl.KindConstructor {
					value = jen.Id("vm").Dot("ToValue").Call(
						jen.Id(m.as(jsConstructorFactory)).
							Call(jen.Id("adaptee").Dot(m.as(goFactory))),
					)
				} else {
					value = jen.Id("vm").Dot("ToValue").Call(
						jen.Id("adapter").Dot(m.as(jsMethod)),
					)
				}

				g.Id("dict").Index(jen.Lit(m.src.Name)).Op("=").Add(value)

				if m.src.Modifiers.Contains(idl.ModifierDefault) {
					defaultExport = value.Dot("ToObject").Call(jen.Id("vm"))
				}
			})
			g.Empty()

			g.Id("mi").Dot("exports").Dot("Named").Op("=").Id("dict")
			g.Empty()

			if defaultExport == nil {
				g.Id("obj").Op(":=").Id("vm").Dot("NewObject").Call()
			} else {
				g.Id("obj").Op(":=").Add(defaultExport)
			}
			g.Empty()

			c.foreach(c.src.Properties, func(p *coder) {
				if p.src.Modifiers.Contains(idl.ModifierDefault) {
					return
				}

				var setr *jen.Statement

				if p.src.Modifiers.Writeable() {
					setr = p.asToValueCall(jen.Id("adapter").Dot(p.as(jsSetter)))
				} else {
					setr = gojaUndefined()
				}

				g.If(jen.Id("err").Op(":=").Id("obj").Dot("DefineAccessorProperty").Call(
					jen.Lit(p.src.Name),
					jen.Id("vm").Dot("ToValue").Call(
						jen.Id("adapter").Dot(p.as(jsGetter))),
					setr, gojaFlagFalse(), gojaFlagTrue()),
					jen.Id("err").Op("!=").Id("nil")).
					Block(jen.Id("panic").Call(jen.Id("err")))
				g.Empty()
			})

			g.Id("mi").Dot("exports").Dot("Default").Op("=").Id("obj")
			g.Empty()

			g.Return().Id("mi")
		})
}

func (c *coder) namespaceSketch(out *jen.Group) {
	out.Comment("init invokes module registration.")
	out.Func().Id("init").Params().Block(jen.Id("register").Call(jen.Id("newModule")))

	out.Comment("newModule returns new goModuleImpl instance.")
	out.Func().
		Id("newModule").
		Params(jen.Id("_").Qual(modulesPath, "VU")).
		Params(jen.Id("goModule")).
		Block(jen.Return().Id("new").Call(jen.Id("goModuleImpl")))

	src := c.toModuleInterface()

	src.goInterfaceSketch(out)
}
