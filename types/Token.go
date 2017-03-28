package types

import (
	"github.com/gsora/hid-compiler/modifiers"
)

type Token struct {
	RawString         string
	RealString        string
	StringModifier    modifiers.StringModifier
	CharacterModifier modifiers.CharacterModifier
}
