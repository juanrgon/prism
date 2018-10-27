package ansi

// BackgroundColor represents an background color ansi escape code
type BackgroundColor Code

const (
	// BlackBackground is the ansi escape code for a black background
	BlackBackground = BackgroundColor("40")
	// RedBackground is the ansi escape code for a red background
	RedBackground= BackgroundColor("41")
	// GreenBackground is the ansi escape code for a green background
	GreenBackground= BackgroundColor("42")
	// YellowBackground is the ansi escape code for a yellow background
	YellowBackground= BackgroundColor("43")
	// BlueBackground is the ansi escape code for a blue background
	BlueBackground= BackgroundColor("44")
	// MagentaBackground is the ansi escape code for a magenta background
	MagentaBackground= BackgroundColor("45")
	// CyanBackground is the ansi escape code for a cyan background
	CyanBackground= BackgroundColor("46")
	// WhiteBackground is the ansi escape code for a white background
	WhiteBackground= BackgroundColor("47")
	// DefaultBackground is the ansi escape code for the default background color
	DefaultBackground = "49"
)
