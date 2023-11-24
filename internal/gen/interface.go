package gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/szkiba/tygor/internal/idl"
)

func (c *coder) jsInterfaceCode(out *jen.Group) {
	iface := c.as(jsInterface)

	rem(out, "%s is the go binding for the JavaScript %s type.", iface, c.src.Name)
	tsdoc(out, c.src)

	out.Type().Id(iface).InterfaceFunc(func(g *jen.Group) {
		c.foreach(c.src.Methods, func(m *coder) {
			m.jsMethodCode(g)
		})

		c.foreach(c.src.Properties, func(p *coder) {
			p.jsPropertyCode(g)
		})
	})
}

func (c *coder) jsMethodCode(out *jen.Group) {
	if c.src.Kind == idl.KindConstructor {
		return
	}

	rem(out, "%s is the go binding for the JavaScript %s method.", c.as(jsMethod), c.src.Name)
	tsdoc(out, c.src)

	out.Id(c.as(jsMethod)).Params(callParam(), vmParam()).Params(valueType())
	out.Empty()
}

func (c *coder) jsPropertyCode(out *jen.Group) {
	c.jsGetterCode(out)

	if c.src.Modifiers.Writeable() {
		c.jsSetterCode(out)
	}
}

func (c *coder) jsGetterCode(out *jen.Group) {
	rem(
		out,
		"%s is the go getter binding for the JavaScript %s property.",
		c.as(jsGetter),
		c.src.Name,
	)
	tsdoc(out, c.src)

	out.Id(c.as(jsGetter)).Params(callParam(), vmParam()).Params(valueType())
	out.Empty()
}

func (c *coder) jsSetterCode(out *jen.Group) {
	rem(
		out,
		"%s is the go setter binding for the JavaScript %s property.",
		c.as(jsSetter),
		c.src.Name,
	)
	tsdoc(out, c.src)

	out.Id(c.as(jsSetter)).Params(callParam(), vmParam()).Params(valueType())
	out.Empty()
}

func (c *coder) goInterfaceCode(out *jen.Group) {
	rem(
		out,
		"%s is the go representation of the JavaScript %s type.",
		c.as(goInterface),
		c.src.Name,
	)
	tsdoc(out, c.src)

	out.Type().Id(c.as(goInterface)).InterfaceFunc(func(g *jen.Group) {
		c.foreach(c.src.Methods, func(m *coder) {
			m.goMethodCode(g)
		})

		c.foreach(c.src.Properties, func(p *coder) {
			p.goPropertyCode(g)
		})
	})
}

func (c *coder) goMethodCode(out *jen.Group) {
	var method string
	if c.src.Kind == idl.KindConstructor {
		method = c.as(goFactory)
		rem(out, "%s is the go factory method for the %s type.", method, c.src.Name)
	} else {
		method = c.as(goMethod)
		rem(out, "%s is the go representation of the %s method.", method, c.src.Name)
		tsdoc(out, c.src)
	}

	out.Id(method).Params(c.parameters()).Params(c.asType(), jen.Id("error"))
	out.Empty()
}

func (c *coder) goPropertyCode(out *jen.Group) {
	c.goGetterCode(out)

	if c.src.Modifiers.Writeable() {
		c.goSetterCode(out)
	}
}

func (c *coder) goGetterCode(out *jen.Group) {
	rem(out, "%s is the go getter method for the %s property.", c.as(goGetter), c.src.Name)
	tsdoc(out, c.src)

	out.Id(c.as(goGetter)).Params().Params(c.asType(), jen.Id("error"))
	out.Empty()
}

func (c *coder) goSetterCode(out *jen.Group) {
	rem(out, "%s is the go setter method for the %s property.", c.as(goSetter), c.src.Name)
	tsdoc(out, c.src)

	out.Id(c.as(goSetter)).Params(jen.Id("v").Add(c.asType())).Params(jen.Id("error"))
	out.Empty()
}

func (c *coder) goInterfaceSketch(out *jen.Group) {
	rem(out, "%s is an empty implementation of %s.", c.as(goInterfaceImpl), c.as(goInterface))
	out.Type().Id(c.as(goInterfaceImpl)).Struct()

	out.Var().
		Id("_").
		Id(c.as(goInterface)).
		Op("=").
		Parens(jen.Op("*").Id(c.as(goInterfaceImpl))).
		Call(jen.Id("nil"))

	c.foreach(c.src.Methods, func(m *coder) {
		m.goMethodSketch(c, out)
	})

	c.foreach(c.src.Properties, func(p *coder) {
		p.goGetterSketch(c, out)

		if p.src.Modifiers.Writeable() {
			p.goSetterSketch(c, out)
		}
	})
}

func (c *coder) goMethodSketch(parent *coder, g *jen.Group) {
	var method string
	if c.src.Kind == idl.KindConstructor {
		method = c.as(goFactory)
	} else {
		method = c.as(goMethod)
	}

	rem(g, "%s is a %s method implementation.", method, parent.as(goInterface))
	g.Func().Params(jen.Id("self").Op("*").Id(parent.as(goInterfaceImpl))).Id(method).
		Params(c.parameters()).
		Params(c.asType(), errorType()).
		BlockFunc(func(g *jen.Group) {
			g.Return().ListFunc(func(g *jen.Group) {
				if zero := c.asZero(); zero != nil {
					g.Add(zero)
				}
				g.Add(jen.Qual(errorsPath, "ErrUnsupported"))
			})
		})
}

func (c *coder) goGetterSketch(parent *coder, out *jen.Group) {
	rem(
		out,
		"%s is a %s property getter method implementation.",
		c.as(goGetter),
		parent.as(goInterface),
	)
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(goInterfaceImpl))).Id(c.as(goGetter)).
		Params().
		Params(c.asType(), errorType()).BlockFunc(func(g *jen.Group) {
		g.Return().List(c.asZero(), jen.Qual(errorsPath, "ErrUnsupported"))
	})
}

func (c *coder) goSetterSketch(parent *coder, out *jen.Group) {
	rem(out, "%s is a %s property setter method.", c.as(goSetter), parent.as(goInterface))
	out.Func().Params(jen.Id("self").Op("*").Id(parent.as(goInterfaceImpl))).Id(c.as(goSetter)).
		Params(c.asParameter()).
		ParamsFunc(func(g *jen.Group) {
			g.Add(jen.Id("error"))
		}).BlockFunc(func(g *jen.Group) {
		g.Return().Qual(errorsPath, "ErrUnsupported")
	})
}
