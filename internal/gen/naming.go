package gen

import "fmt"

type nameable int

const (
	goInterface nameable = iota
	goMethod
	goParam
	goSetter
	goGetter
	goAdapter
	goConstructor
	goFactory
	goInterfaceImpl
	jsInterface
	jsMethod
	jsParam
	jsSetter
	jsGetter
	jsAdapter
	jsConstructorFactory
)

type naming func(nameable, string) string

var nameFormats = map[nameable]string{ //nolint:gochecknoglobals
	goInterface:          "go%s",
	goMethod:             "%sMethod",
	goSetter:             "%sSetter",
	goGetter:             "%sGetter",
	goParam:              "%sArg",
	goAdapter:            "go%sAdapter",
	goConstructor:        "go%sConstructor",
	goFactory:            "new%s",
	goInterfaceImpl:      "go%sImpl",
	jsInterface:          "js%s",
	jsMethod:             "%sMethod",
	jsSetter:             "%sSetter",
	jsGetter:             "%sGetter",
	jsParam:              "%sArg",
	jsAdapter:            "js%sAdapter",
	jsConstructorFactory: "new%sConstructor",
}

func defaultNaming(what nameable, src string) string {
	format, ok := nameFormats[what]
	if !ok {
		panic(fmt.Sprintf("unknown nameable: %d", what))
	}

	return fmt.Sprintf(format, src)
}
