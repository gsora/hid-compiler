package compiler

type modifier int

//go:generate stringer -type=modifier
const (
	LCTRL modifier = 1 << iota
	LSHIFT
	LALT
	LMETA
	RCTRL
	RSHIFT
	RALT
	RMETA
	NOMOD
)

type stringModifier struct {
	StringRef string
	Modifier  modifier
}

type characterModifier struct {
	CharRef  byte
	Modifier modifier
	Left     string
	Right    string
}

type token struct {
	RawString         string
	RealString        string
	StringModifier    stringModifier
	CharacterModifier characterModifier
}
