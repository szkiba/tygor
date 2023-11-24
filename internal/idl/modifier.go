package idl

// Modifier represents declaration modifier.
type Modifier int

const (
	// ModifierUnknown is a Modifier enum value for unknown modifier.
	ModifierUnknown Modifier = iota
	// ModifierConst is a Modifier enum value for const modifier.
	ModifierConst
	// ModifierReadonly is a Modifier enum value for readonly modifier.
	ModifierReadonly
	// ModifierAsync is a Modifier enum value for async modifier.
	ModifierAsync
	// ModifierDefault is a Modifier enum value for default modifier.
	ModifierDefault
)

//go:generate enumer -trimprefix Modifier -transform snake-upper -json -text -values -type Modifier

// Modifiers holds all modifiers.
type Modifiers []Modifier

// Writeable returns true if Modifiers doesn't contains const or readonly.
func (m Modifiers) Writeable() bool {
	for _, mod := range m {
		if mod == ModifierConst || mod == ModifierReadonly {
			return false
		}
	}

	return true
}

// Contains returns true if Modifiers contains a given modifier.
func (m Modifiers) Contains(modifier Modifier) bool {
	for _, mod := range m {
		if mod == modifier {
			return true
		}
	}

	return false
}

// Default returns true if Modifiers contains default modifier.
func (m Modifiers) Default() bool {
	return m.Contains(ModifierDefault)
}
