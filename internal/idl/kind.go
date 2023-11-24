package idl

// Kind describes declaration kind.
type Kind int

const (
	// KindUnknown is a Kind enum value for unknown declaration.
	KindUnknown Kind = iota
	// KindClass is a Kind enum value for class declaration.
	KindClass
	// KindInterface is a Kind enum value for interface declaration.
	KindInterface
	// KindFunction is a Kind enum value for function declaration.
	KindFunction
	// KindVariable is a Kind enum value for variable declaration.
	KindVariable
	// KindNamespace is a Kind enum value for namespace declaration.
	KindNamespace
	// KindMethod is a Kind enum value for method declaration.
	KindMethod
	// KindConstructor is a Kind enum value for constructor declaration.
	KindConstructor
	// KindProperty is a Kind enum value for property declaration.
	KindProperty
	// KindParameter is a Kind enum value for parameter declaration.
	KindParameter
)

//go:generate enumer -trimprefix Kind -transform snake-upper -json -text -values -type Kind
