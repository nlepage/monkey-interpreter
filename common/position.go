package common

// Position points to somewhere in a file
type Position struct {
	File   string
	Line   int
	Column int
}
