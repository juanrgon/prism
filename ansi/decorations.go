package ansi

// Decoration represents an ansi decoration
type Decoration Code

const (
	// Bold is the ansi escape code for bold text
	Bold = Decoration(Code(string(3*iota + 1)))
	// Underlined is the ansi escape code for underlined text
	Underlined
	// Inverted is the ansi escape code for swapping the foreground and background colors
	Inverted
)
