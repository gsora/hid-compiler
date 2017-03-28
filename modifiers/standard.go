package modifiers

type Modifier int

//go:generate stringer -type=Modifier
const (
	LCTRL Modifier = 1 << iota
	LSHIFT
	LALT
	LMETA
	RCTRL
	RSHIFT
	RALT
	RMETA
	NOMOD
)

type StringModifier struct {
	StringRef string
	Modifier  Modifier
}

type CharacterModifier struct {
	CharRef  byte
	Modifier Modifier
	Left     string
	Right    string
}
