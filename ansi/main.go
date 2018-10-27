package ansi

// Code is an ansi code e.g. "30" (red text), or "45" (purple background)
type Code string

// Begin is the beginning string of all ansi escape sequence
const Begin = "\033["

// End is the string for ending an ansi escape sequence
const End = "m"

// Reset is the ansi escape sequence for clearing all text modifications
const Reset = Begin + "0" + End
