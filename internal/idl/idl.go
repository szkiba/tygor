// Package idl contains TypeScript declaration based IDL API.
package idl

// Tags represents TypeDoc tags.
type Tags map[string][]string

// Declaration contains delaration properties.
type Declaration struct {
	Name         string       `json:"name"`
	Doc          string       `json:"doc,omitempty"`
	Kind         Kind         `json:"kind,omitempty"`
	Type         string       `json:"type,omitempty"`
	Modifiers    Modifiers    `json:"modifiers,omitempty"`
	Methods      Declarations `json:"methods,omitempty"`
	Properties   Declarations `json:"properties,omitempty"`
	Constructors Declarations `json:"constructors,omitempty"`
	Tags         Tags         `json:"tags,omitempty"`
	Source       string       `json:"source,omitempty"`

	Parameters Declarations `json:"parameters,omitempty"`
}

// Declarations represent a list of related declarations.
type Declarations []*Declaration

// Module describes TypeScript module.
type Module struct {
	Name         string       `json:"name,omitempty"`
	Declarations Declarations `json:"declarations,omitempty"`
}

// Get returns Declaration by name.
func (d Declarations) Get(name string) *Declaration {
	for _, decl := range d {
		if decl != nil && decl.Name == name {
			return decl
		}
	}

	return nil
}
