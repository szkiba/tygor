package idl

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/dop251/goja"
	"github.com/jmespath/go-jmespath"
	"github.com/stretchr/testify/assert"
)

func testExtract(t *testing.T, filename, source, query string) Declarations { //nolint:unparam
	t.Helper()

	fn, err := newExtractFunc(extractorScriptName, extractorScript, testRuntime(t))

	assert.NoError(t, err)

	str, err := fn(filename, source)

	assert.NoError(t, err)

	var data interface{}

	err = json.Unmarshal([]byte(str), &data)

	assert.NoError(t, err)

	result, err := jmespath.Search(query, data)

	assert.NoError(t, err)

	bin, err := json.Marshal(result)

	assert.NoError(t, err)

	var decls []*Declaration

	err = json.Unmarshal(bin, &decls)

	assert.NoError(t, err)

	return decls
}

type console struct {
	t *testing.T
}

func (c *console) valueString(v goja.Value) string {
	c.t.Helper()

	m, ok := v.(json.Marshaler)
	if !ok {
		return v.String()
	}

	b, err := json.Marshal(m)
	if err != nil {
		return v.String()
	}

	return string(b)
}

func (c *console) format(args ...goja.Value) string {
	c.t.Helper()

	var strs strings.Builder
	for i := 0; i < len(args); i++ {
		if i > 0 {
			strs.WriteString(" ")
		}
		strs.WriteString(c.valueString(args[i]))
	}

	return strs.String()
}

func (c *console) log(args ...goja.Value) {
	c.t.Helper()

	c.t.Log(c.format(args...))
}

func (c *console) err(args ...goja.Value) {
	c.t.Helper()

	c.t.Error(c.format(args...))
}

//nolint:errcheck,gosec
func testRuntime(t *testing.T) *goja.Runtime {
	t.Helper()

	runtime := goja.New()

	c := &console{t}

	obj := runtime.NewObject()

	obj.Set("debug", c.log)
	obj.Set("log", c.log)
	obj.Set("info", c.log)
	obj.Set("warn", c.log)
	obj.Set("error", c.err)
	obj.Set("dir", c.err)

	runtime.Set("console", obj)

	return runtime
}
