package doc

// Format describes output formt.
type Format int

const (
	// FormatUnknown is a Format enum value for unknown format.
	FormatUnknown Format = iota
	// FormatMarkdown is a Format enum value for markdown output format.
	FormatMarkdown
	// FormatHTML is a Format enum value for HTML output format.
	FormatHTML
)

//go:generate enumer -trimprefix Format -transform snake-upper -json -text -values -type Format
