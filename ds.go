package prism

// DecoratedString is a string with methods for coloring
type DecoratedString string

// InBlack returns this DecoratedString with black text
func (ds DecoratedString) InBlack() DecoratedString {
	return InBlack(string(ds))
}

// InRed returns this DecoratedString with red text
func (ds DecoratedString) InRed() DecoratedString {
	return InRed(string(ds))
}

// InGreen returns this DecoratedString with green text
func (ds DecoratedString) InGreen() DecoratedString {
	return InGreen(string(ds))
}

// InYellow returns this DecoratedString with yellow text
func (ds DecoratedString) InYellow() DecoratedString {
	return InYellow(string(ds))
}

// InBlue returns this DecoratedString with blue text
func (ds DecoratedString) InBlue() DecoratedString {
	return InBlue(string(ds))
}

// InMagenta returns this DecoratedString with magenta text
func (ds DecoratedString) InMagenta() DecoratedString {
	return InMagenta(string(ds))
}

// InCyan returns this DecoratedString with cyan text
func (ds DecoratedString) InCyan() DecoratedString {
	return InCyan(string(ds))
}

// InWhite returns this DecoratedString with white text
func (ds DecoratedString) InWhite() DecoratedString {
	return InWhite(string(ds))
}

// OnBlack returns this decorated string with a black background
func (ds DecoratedString) OnBlack() DecoratedString {
	return OnBlack(string(ds))
}

// OnRed returns this DecoratedString with a red background
func (ds DecoratedString) OnRed() DecoratedString {
	return OnRed(string(ds))
}

// OnGreen returns this DecoratedString with a green background
func (ds DecoratedString) OnGreen() DecoratedString {
	return OnGreen(string(ds))
}

// OnYellow returns this DecoratedString with a yellow background
func (ds DecoratedString) OnYellow() DecoratedString {
	return OnYellow(string(ds))
}

// OnBlue returns this DecoratedString with a blue background
func (ds DecoratedString) OnBlue() DecoratedString {
	return OnBlue(string(ds))
}

// OnMagenta returns this DecoratedString with a magenta background
func (ds DecoratedString) OnMagenta() DecoratedString {
	return OnMagenta(string(ds))
}

// OnCyan returns this DecoratedString with a cyan background
func (ds DecoratedString) OnCyan() DecoratedString {
	return OnCyan(string(ds))
}

// OnWhite returns this DecoratedString with a white background
func (ds DecoratedString) OnWhite() DecoratedString {
	return OnWhite(string(ds))
}

// Underlined returns this DecoratedString underlined
func (ds DecoratedString) Underlined() DecoratedString {
	return Underlined(string(ds))
}

// InBold returns this DecoratedString in bold
func (ds DecoratedString) InBold() DecoratedString {
	return InBold(string(ds))
}

func (ds DecoratedString) Inverted() DecoratedString {
	return Inverted(string(ds))
}
