//nolint:revive
package idl

import (
	"encoding/json"
)

func Parse(filename, source string) (*Module, error) {
	str, err := extract(filename, source)
	if err != nil {
		return nil, err
	}

	var decls []*Declaration

	if err = json.Unmarshal([]byte(str), &decls); err != nil {
		return nil, err
	}

	mod := new(Module)

	mod.Declarations = decls

	for _, decl := range decls {
		if decl.Kind == KindNamespace {
			mod.Name = decl.Name

			break
		}
	}

	return mod, nil
}
