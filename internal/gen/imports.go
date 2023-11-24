package gen

const (
	gojaPath = "github.com/dop251/goja"

	modulesPath = "go.k6.io/k6/js/modules"

	errorsPath = "errors"

	fmtPath = "fmt"

	timePath = "time"
)

var importNames = map[string]string{ //nolint:gochecknoglobals
	gojaPath:    "goja",
	modulesPath: "modules",
	errorsPath:  "errors",
	fmtPath:     "fmt",
	timePath:    "time",
}
