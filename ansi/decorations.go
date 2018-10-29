package ansi

// Decoration represents an ansi decoration
type Decoration Code

const (
	// Bold is the ansi escape code for bold text
	Bold = Decoration("1")
	// Underlined is the ansi escape code for underlined text
	Underlined = Decoration("4")
	// Inverted is the ansi escape code for swapping the foreground and background colors
	Inverted = Decoration("7")
)
