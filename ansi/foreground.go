package ansi


// ForegroundColor represents an foreground color ansi escape code
type ForegroundColor Code


const (
	// BlackText is the ansi escape Code for black text
	BlackText = ForegroundColor("30")
	// RedText is the ansi escape code for a red text
	RedText = ForegroundColor("31")
	// GreenText is the ansi escape code for a green text
	GreenText = ForegroundColor("32")
	// YellowText is the ansi escape code for a yellow text
	YellowText = ForegroundColor("33")
	// BlueText is the ansi escape code for a blue text
	BlueText = ForegroundColor("34")
	// MagentaText is the ansi escape code for a magenta text
	MagentaText = ForegroundColor("35")
	// CyanText is the ansi escape code for  text
	CyanText = ForegroundColor("36")
	// WhiteText is the ansi escape code for a white text
	WhiteText = ForegroundColor("37")
	// DefaultText is the ansi escape code for the default text color
	DefaultText = "39"
)
