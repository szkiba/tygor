package gen

import "github.com/dave/jennifer/jen"

func runtimeCode(out *jen.Group) {
	goTimeFromDateCode(out)
	out.Empty()
	jsDateFromTimeCode(out)
}

func goTimeFromDateCode(out *jen.Group) {
	out.Func().
		Id("goTimeFromDate").
		Params(jen.Id("v").Id("goja").Dot("Value"), jen.Id("vm").Op("*").Id("goja").Dot("Runtime")).
		Params(jen.Qual("time", "Time")).
		Block(
			jen.List(jen.Id("getTime"), jen.Id("ok")).
				Op(":=").
				Qual(gojaPath, "AssertFunction").
				Call(jen.Id("v").Dot("ToObject").Call(jen.Id("vm")).Dot("Get").Call(jen.Lit("getTime"))),
			jen.If(jen.Op("!").Id("ok")).
				Block(jen.Id("panic").Call(
					jen.Qual(fmtPath, "Errorf").
						Call(jen.Lit("%w: getTime"), jen.Qual(errorsPath, "ErrUnsupported")),
				)),
			jen.Empty(),
			jen.List(jen.Id("ret"), jen.Id("err")).
				Op(":=").
				Id("getTime").
				Call(jen.Id("v")),
			jen.If(jen.Id("err").Op("!=").Id("nil")).
				Block(jen.Id("panic").Call(jen.Id("err"))),
			jen.Empty(),
			jen.Return().Qual("time", "UnixMilli").Call(jen.Id("ret").Dot("ToInteger").Call()))
}

func jsDateFromTimeCode(out *jen.Group) {
	out.Func().
		Id("jsDateFromTime").
		Params(jen.Id("t").Qual(timePath, "Time"), jen.Id("vm").Op("*").Qual(gojaPath, "Runtime")).
		Params(jen.Qual(gojaPath, "Value")).
		Block(
			jen.List(jen.Id("d"), jen.Id("err")).
				Op(":=").
				Id("vm").
				Dot("New").
				Call(
					jen.Id("vm").
						Dot("Get").
						Call(jen.Lit("Date")),
					jen.Id("vm").Dot("ToValue").Call(jen.Id("t").Dot("UnixMilli").Call())),
			jen.If(jen.Id("err").Op("!=").Id("nil")).
				Block(jen.Id("panic").Call(jen.Id("err"))),
			jen.Empty(),
			jen.Return().Id("d"))
}

func runtimeSketch(_ *jen.Group) {}
