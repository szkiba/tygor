package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/szkiba/tygor/internal/idl"
)

func (c *coder) jsAdapterCode(out *jen.Group) {
	rem(out, "%s converts %s to %s.", c.as(jsAdapter), c.as(goInterface), c.as(jsInterface))
	out.Type().Id(c.as(jsAdapter)).Struct(jen.Id("adaptee").Id(c.as(goInterface)))

	out.Empty()
	out.Var().
		Id("_").
		Id(c.as(jsInterface)).
		Op("=").
		Parens(jen.Op("*").Id(c.as(jsAdapter))).
		Call(jen.Id("nil"))

	c.foreach(c.src.Methods, func(m *coder) {
		m.jsAdapterMethodCode(c, out)
	})

	c.foreach(c.src.Properties, func(p *coder) {
		p.jsAdapterGetterCode(c, out)

		if p.src.Modifiers.Writeable() {
			p.jsAdapterSetterCode(c, out)
		}
	})
}

func (c *coder) jsAdapterMethodCode(parent *coder, out *jen.Group) {
	if c.src.Kind == idl.KindConstructor {
		return
	}

	rem(out, "%s is a %s adapter method.", c.as(jsMethod), parent.as(jsInterface))
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(jsAdapter))).Id(c.as(jsMethod)).
		ParamsFunc(jsFunctionParams).ParamsFunc(jsFunctionReturns).
		BlockFunc(func(g *jen.Group) {
			var vars *jen.Statement
			var ret *jen.Statement

			if c.src.Type == "void" {
				vars = jen.Id("err")
				ret = gojaUndefined()
			} else {
				vars = jen.List(jen.Id("v"), jen.Id("err"))
				ret = c.asToValueCall(jen.Id("v"))
			}

			g.Add(vars.
				Op(":=").
				Id("self").
				Dot("adaptee").
				Dot(c.as(goMethod)).
				CallFunc(func(g *jen.Group) {
					idx := 0
					c.foreach(c.src.Parameters, func(p *coder) {
						g.Add(
							p.asValueToCall(
								jen.Id("call").Dot("Argument").Call(jen.Lit(idx)),
								jen.Id("vm"),
							),
						)
						idx++
					})
				}))

			g.Add(jen.If(jen.Id("err").Op("!=").Id("nil")).
				Block(jen.Panic(jen.Id("err"))))

			g.Empty()
			g.Add(jen.Return().Add(ret))
		})
}

func (c *coder) jsAdapterGetterCode(parent *coder, out *jen.Group) {
	rem(out, "%s is a %s property getter adapter method.", c.as(jsGetter), parent.as(jsInterface))
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(jsAdapter))).Id(c.as(jsGetter)).
		ParamsFunc(jsFunctionParams).ParamsFunc(jsFunctionReturns).
		BlockFunc(func(g *jen.Group) {
			g.Add(
				jen.List(jen.Id("v"), jen.Id("err")).
					Op(":=").
					Id("self").Dot("adaptee").Dot(c.as(goGetter)).Call(),
			)
			g.Add(
				jen.If(jen.Id("err").Op("!=").Id("nil")).
					Block(jen.Panic(jen.Id("err"))),
			)
			g.Empty()
			g.Add(jen.Return().Add(c.asToValueCall(jen.Id("v"))))
		})
}

func (c *coder) jsAdapterSetterCode(parent *coder, out *jen.Group) {
	rem(out, "%s is a %s property setter adapter method.", c.as(jsSetter), parent.as(jsInterface))
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(jsAdapter))).Id(c.as(jsSetter)).
		ParamsFunc(jsFunctionParams).ParamsFunc(jsFunctionReturns).
		BlockFunc(func(g *jen.Group) {
			g.Add(
				jen.Id("err").Op(":=").Id("self").
					Dot("adaptee").
					Dot(c.as(goSetter)).
					Call(c.asValueToCall(jen.Id("call").Dot("Argument").Call(jen.Lit(0)), jen.Id("vm"))),
			)
			g.Add(
				jen.If(jen.Id("err").Op("!=").Id("nil")).
					Block(jen.Panic(jen.Id("err"))),
			)
			g.Empty()
			g.Add(jen.Return(gojaUndefined()))
		})
}

func (c *coder) jsFromCode(out *jen.Group) {
	iname := c.as(jsInterface)

	rem(out, "%sFrom returns a %s based on a %s.", iname, iname, c.as(goInterface))
	out.Func().
		Id(iname + "From").
		Params(jen.Id("adaptee").Id(c.as(goInterface))).
		Params(jen.Id(iname)).
		Block(
			jen.Return().
				Op("&").
				Id(c.as(jsAdapter)).
				Values(jen.Id("adaptee").Op(":").Id("adaptee")),
		)
}

func (c *coder) goAdapterCode(out *jen.Group) {
	rem(out, "%s converts goja Object to %s.", c.as(goAdapter), c.as(goInterface))
	out.Type().Id(c.as(goAdapter)).
		Struct(jen.Id("adaptee").Op("*").Qual(gojaPath, "Object"), jen.Id("vm").Op("*").Qual(gojaPath, "Runtime"))

	out.Var().
		Id("_").
		Id(c.as(goInterface)).
		Op("=").
		Parens(jen.Op("*").Id(c.as(goAdapter))).
		Call(jen.Id("nil"))

	c.foreach(c.src.Methods, func(m *coder) {
		m.goAdapterMethodCode(c, out)
	})

	c.foreach(c.src.Properties, func(p *coder) {
		p.goAdapterGetterCode(c, out)

		if p.src.Modifiers.Writeable() {
			p.goAdapterSetterCode(c, out)
		}
	})
}

func (c *coder) goAdapterMethodCode(parent *coder, out *jen.Group) {
	var method string
	if c.src.Kind == idl.KindConstructor {
		method = c.as(goFactory)
		rem(out, "%s is a %s facotry method.", method, c.src.Name)
	} else {
		method = c.as(goMethod)
		rem(out, "%s is a %s adapter method.", method, c.src.Name)
	}

	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(goAdapter))).Id(method).
		Params(c.parameters()).
		Params(c.asType(), errorType()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("fun"), jen.Id("ok")).
				Op(":=").
				Qual(gojaPath, "AssertFunction").
				Call(jen.Id("self").Dot("adaptee").Dot("Get").Call(jen.Lit(c.src.Name)))
			g.If(jen.Op("!").Id("ok")).
				Block(
					jen.Return().
						List(
							c.asZero(),
							jen.Qual(fmtPath, "Errorf").
								Call(jen.Lit("%w: "+c.src.Name), jen.Qual(errorsPath, "ErrUnsupported")),
						),
				)
			g.Empty()

			if c.src.Type == "void" {
				g.List(jen.Id("_"), jen.Id("err")).
					Op(":=").
					Id("fun").
					Call(jen.Id("self").Dot("adaptee"))

				g.Empty()

				g.Return().List(jen.Id("err"))
			} else {
				g.List(jen.Id("res"), jen.Id("err")).
					Op(":=").
					Id("fun").
					Call(jen.Id("self").Dot("adaptee"))
				g.If(jen.Id("err").Op("!=").Id("nil")).
					Block(jen.Return().List(c.asZero(), jen.Id("err")))

				g.Empty()

				g.Return().List(c.asValueToCall(
					jen.Id("res"),
					jen.Id("self").Dot("vm"),
				), jen.Id("nil"))
			}
		})
}

func (c *coder) goAdapterGetterCode(parent *coder, out *jen.Group) {
	rem(out, "%s is a %s property getter adapter method.", c.as(goGetter), parent.as(goInterface))
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(goAdapter))).Id(c.as(goGetter)).
		Params().
		Params(c.asType(), errorType()).BlockFunc(func(g *jen.Group) {
		g.Return().
			List(
				c.asValueToCall(jen.Id("self").Dot("adaptee").Dot("Get").Call(jen.Lit(c.src.Name)), jen.Id("self").Dot("vm")),
				jen.Id("nil"))
	})
}

func (c *coder) goAdapterSetterCode(parent *coder, out *jen.Group) {
	rem(out, "%s is a %s property setter adapter method.", c.as(goSetter), parent.as(goInterface))
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(goAdapter))).Id(c.as(goSetter)).
		Params(jen.Id("v").Add(c.asType())).
		ParamsFunc(func(g *jen.Group) {
			g.Add(jen.Id("error"))
		}).BlockFunc(func(g *jen.Group) {
		g.Return().Id("self").Dot("adaptee").Dot("Set").
			Call(jen.Lit(c.src.Name), jen.Id("v"))
	})
}

func (c *coder) goFromCode(out *jen.Group) {
	rem(out, "%sFrom returns a %s from goja Object.", c.as(goInterface), c.as(goInterface))
	out.Func().
		Id(c.as(goInterface)+"From").
		Params(jen.Id("adaptee").Op("*").Qual(gojaPath, "Object"), jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")).
		Params(jen.Id(c.as(goInterface))).
		Block(
			jen.Return().
				Op("&").
				Id(c.as(goAdapter)).
				Values(jen.Id("adaptee").Op(":").Id("adaptee"), jen.Id("vm").Op(":").Id("vm")),
		)
}

func (c *coder) goToObjectCode(out *jen.Group) {
	rem(out, "%sToObject returns a goja Object from %s.", c.as(goInterface), c.as(goInterface))
	out.Func().
		Id(c.as(goInterface)+"ToObject").
		Params(jen.Id("v").Id(c.as(goInterface)), jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")).
		Params(jen.Op("*").Qual(gojaPath, "Object")).
		Block(
			jen.Id("obj").
				Op(":=").
				Id("vm").
				Dot("NewObject").
				Call(),
			jen.Empty(),
			jen.Id("err").
				Op(":=").
				Id(c.as(jsInterface)+"To").
				Call(jen.Id(c.as(jsInterface)+"From").Call(jen.Id("v")), jen.Id("obj"), jen.Id("vm")),
			jen.If(jen.Id("err").Op("!=").Id("nil")).
				Block(jen.Id("panic").Call(jen.Id("err"))),
			jen.Empty(),
			jen.Return().Id("obj"))
}

func (c *coder) jsConstructorCode(out *jen.Group) {
	ctor := newCoder(c.src.Constructors[0], c.naming)

	rem(out, "%s creates new %s instance.", c.as(goConstructor), c.as(goInterface))
	out.Type().
		Id(c.as(goConstructor)).
		Func().
		Params(ctor.parameters()).
		Params(jen.Id(c.as(goInterface)), jen.Id("error"))

	rem(out, "%s creates %s JavaScript constructor.", c.as(jsConstructorFactory), c.src.Name)
	out.Func().
		Id(c.as(jsConstructorFactory)).
		Params(jen.Id("ctor").Id(c.as(goConstructor))).
		Params(
			jen.Func().
				Params(jen.Id("call").Qual(gojaPath, "ConstructorCall"), jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")).
				Params(jen.Op("*").Qual(gojaPath, "Object")),
		).
		Block(
			jen.Return().
				Func().
				Params(jen.Id("call").Qual(gojaPath, "ConstructorCall"), jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")).
				Params(jen.Op("*").Qual(gojaPath, "Object")).
				BlockFunc(func(g *jen.Group) {
					g.List(jen.Id("adaptee"), jen.Id("err")).
						Op(":=").
						Id("ctor").
						CallFunc(func(g *jen.Group) {
							idx := 0
							c.foreach(c.src.Constructors[0].Parameters, func(p *coder) {
								g.Add(
									p.asValueToCall(
										jen.Id("call").Dot("Argument").Call(jen.Lit(idx)),
										jen.Id("vm"),
									),
								)
								idx++
							})
						})
					g.If(jen.Id("err").Op("!=").Id("nil")).
						Block(jen.Id("panic").Call(jen.Id("err")))
					g.Empty()

					g.Id("adapter").Op(":=").Id(c.as(jsInterface) + "From").Call(jen.Id("adaptee"))
					g.Empty()

					g.If(
						jen.Id("err").
							Op(":=").
							Id(c.as(jsInterface)+"To").
							Call(jen.Id("adapter"), jen.Id("call").Dot("This"), jen.Id("vm")),
						jen.Id("err").Op("!=").Id("nil"),
					).
						Block(jen.Id("panic").Call(jen.Id("err")))
					g.Empty()

					g.Return().Id("nil")
				}))
}

func (c *coder) jsToCode(out *jen.Group) {
	rem(
		out,
		"%sTo setup %s JavaScript object from %s.",
		c.as(jsInterface),
		c.src.Name,
		c.as(jsInterface),
	)
	out.Func().
		Id(c.as(jsInterface)+"To").
		Params(
			jen.Id("src").
				Id(c.as(jsInterface)),
			jen.Id("obj").
				Op("*").
				Qual(gojaPath, "Object"),
			jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")).
		Params(jen.Id("error")).
		BlockFunc(func(g *jen.Group) {
			count := len(c.src.Methods) + len(c.src.Properties)
			if count == 0 {
				g.Return().Id("nil")
				return
			}

			c.foreach(c.src.Methods, func(m *coder) {
				if count--; count == 0 {
					g.Return().
						Id("obj").
						Dot("Set").
						Call(jen.Lit(m.src.Name), jen.Id("src").Dot(m.as(jsMethod)))
				} else {
					g.If(jen.Id("err").Op(":=").Id("obj").Dot("Set").Call(jen.Lit(m.src.Name), jen.Id("src").Dot(m.as(jsMethod))),
						jen.Id("err").Op("!=").Id("nil")).
						Block(jen.Return().Id("err"))
					g.Empty()
				}
			})

			c.foreach(c.src.Properties, func(prop *coder) {
				var setr *jen.Statement

				if prop.src.Modifiers.Writeable() {
					setr = jen.Id("vm").
						Dot("ToValue").
						Call(jen.Id("src").Dot(prop.as(jsSetter)))
				} else {
					setr = jen.Qual(gojaPath, "Undefined").Call()
				}

				if count--; count == 0 {
					g.Return().Id("obj").Dot("DefineAccessorProperty").
						Call(
							jen.Lit(prop.src.Name),
							jen.Id("vm").
								Dot("ToValue").
								Call(jen.Id("src").Dot(prop.as(jsGetter))),
							setr,
							jen.Qual(gojaPath, "FLAG_FALSE"),
							jen.Qual(gojaPath, "FLAG_TRUE"))
				} else {
					g.If(
						jen.Id("err").Op(":=").Id("obj").Dot("DefineAccessorProperty").
							Call(
								jen.Lit(prop.src.Name),
								jen.Id("vm").Dot("ToValue").Call(jen.Id("src").Dot(prop.as(jsGetter))),
								setr,
								jen.Qual(gojaPath, "FLAG_FALSE"),
								jen.Qual(gojaPath, "FLAG_TRUE")),
						jen.Id("err").Op("!=").Id("nil")).
						Block(jen.Return().Id("err"))
					g.Empty()
				}
			})
		})
}
