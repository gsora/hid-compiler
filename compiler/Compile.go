package compiler

import (
	"github.com/gsora/hid-compiler/linkedlist"
	"github.com/gsora/hid-compiler/modifiers"
	"github.com/gsora/hid-compiler/types"
	"regexp"
	"strings"
)

var stringModifierRe = regexp.MustCompile(`((\[(LSHIFT|LCTRL|LALT|LMETA|RSHIFT|RCTRL|RALT|RMETA)\]>)+(\w+))+`)
var charModifierRe = regexp.MustCompile(`((\w+)+(\[(LSHIFT|LCTRL|LALT|LMETA|RSHIFT|RCTRL|RALT|RMETA)\]:)+(\w+))+`)

func Compile(s string) linkedlist.LinkedList {
	l := linkedlist.NewLinkedList()

	for _, word := range strings.Split(s, " ") {
		l.InsertToken(tokenize(word))
	}

	return l
}

func tokenize(s string) types.Token {
	if matchStringModifier(s) {
		return tokenizeStringModifier(s)
	}

	if matchCharModifier(s) {
		return tokenizeCharModifier(s)
	}

	return types.Token{
		RawString:         s,
		RealString:        s,
		StringModifier:    modifiers.StringModifier{StringRef: "", Modifier: modifiers.NOMOD},
		CharacterModifier: modifiers.CharacterModifier{CharRef: ' ', Modifier: modifiers.NOMOD},
	}
}

func matchStringModifier(s string) bool {
	return stringModifierRe.MatchString(s)
}

func matchCharModifier(s string) bool {
	return charModifierRe.MatchString(s)
}

func tokenizeCharModifier(s string) types.Token {
	t := types.Token{}

	submatch := charModifierRe.FindStringSubmatch(s)
	t.RawString = s
	t.RealString = submatch[2] + submatch[5]

	var m modifiers.Modifier
	switch submatch[4] {
	case "LCTRL":
		m = modifiers.LCTRL
	case "LSHIFT":
		m = modifiers.LSHIFT
	case "LALT":
		m = modifiers.LALT
	case "LMETA":
		m = modifiers.LMETA
	case "RCTRL":
		m = modifiers.RCTRL
	case "RSHIFT":
		m = modifiers.RSHIFT
	case "RALT":
		m = modifiers.RALT
	case "RMETA":
		m = modifiers.RMETA
	}

	t.CharacterModifier = modifiers.CharacterModifier{CharRef: submatch[5][0], Modifier: m, Left: submatch[2], Right: submatch[5]}
	t.StringModifier = modifiers.StringModifier{StringRef: "", Modifier: modifiers.NOMOD}

	return t
}

func tokenizeStringModifier(s string) types.Token {
	t := types.Token{}

	submatch := stringModifierRe.FindStringSubmatch(s)
	t.RawString = s
	t.RealString = submatch[4]

	var m modifiers.Modifier
	switch submatch[3] {
	case "LCTRL":
		m = modifiers.LCTRL
	case "LSHIFT":
		m = modifiers.LSHIFT
	case "LALT":
		m = modifiers.LALT
	case "LMETA":
		m = modifiers.LMETA
	case "RCTRL":
		m = modifiers.RCTRL
	case "RSHIFT":
		m = modifiers.RSHIFT
	case "RALT":
		m = modifiers.RALT
	case "RMETA":
		m = modifiers.RMETA
	}

	t.StringModifier = modifiers.StringModifier{StringRef: submatch[4], Modifier: m}
	t.CharacterModifier = modifiers.CharacterModifier{CharRef: ' ', Modifier: modifiers.NOMOD}

	return t
}
