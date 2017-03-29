package compiler

import (
	"bytes"
	"regexp"
	"strings"
)

var stringModifierRe = regexp.MustCompile(`((\[(LSHIFT|LCTRL|LALT|LMETA|RSHIFT|RCTRL|RALT|RMETA)\]>)+(\w+))+`)
var charModifierRe = regexp.MustCompile(`((\w+)+(\[(LSHIFT|LCTRL|LALT|LMETA|RSHIFT|RCTRL|RALT|RMETA)\]:)+(\w+))+`)

// Compile actually builds the scancode sequence
func Compile(s string) string {
	l := newLinkedList()

	for _, word := range strings.Split(s, " ") {
		l.insertToken(tokenize(word))
	}

	node := l.getFirst()
	var buf bytes.Buffer
	for node.Next != nil {
		d := node.Data
		if d.StringModifier.Modifier != NOMOD {
			buf.WriteString(generateHIDPayloadForString(d.StringModifier))
		} else if d.CharacterModifier.Modifier != NOMOD {
			buf.WriteString(generateHIDPayloadForCharacter(d.CharacterModifier))
		} else {
			buf.WriteString(generateHIDPayloadForStandardString(d.RealString))
		}

		if node.Next != nil {
			buf.WriteString(generateHIDPayloadForSpace())
		}
		node = *node.Next
	}

	return buf.String()
}

func tokenize(s string) token {
	if matchStringModifier(s) {
		return tokenizeStringModifier(s)
	}

	if matchCharModifier(s) {
		return tokenizeCharModifier(s)
	}

	return token{
		RawString:         s,
		RealString:        s,
		StringModifier:    stringModifier{StringRef: "", Modifier: NOMOD},
		CharacterModifier: characterModifier{CharRef: ' ', Modifier: NOMOD},
	}
}

func matchStringModifier(s string) bool {
	return stringModifierRe.MatchString(s)
}

func matchCharModifier(s string) bool {
	return charModifierRe.MatchString(s)
}

func tokenizeCharModifier(s string) token {
	t := token{}

	submatch := charModifierRe.FindStringSubmatch(s)
	t.RawString = s
	t.RealString = submatch[2] + submatch[5]

	var m modifier
	switch submatch[4] {
	case "LCTRL":
		m = LCTRL
	case "LSHIFT":
		m = LSHIFT
	case "LALT":
		m = LALT
	case "LMETA":
		m = LMETA
	case "RCTRL":
		m = RCTRL
	case "RSHIFT":
		m = RSHIFT
	case "RALT":
		m = RALT
	case "RMETA":
		m = RMETA
	}

	t.CharacterModifier = characterModifier{CharRef: submatch[5][0], Modifier: m, Left: submatch[2], Right: submatch[5]}
	t.StringModifier = stringModifier{StringRef: "", Modifier: NOMOD}

	return t
}

func tokenizeStringModifier(s string) token {
	t := token{}

	submatch := stringModifierRe.FindStringSubmatch(s)
	t.RawString = s
	t.RealString = submatch[4]

	var m modifier
	switch submatch[3] {
	case "LCTRL":
		m = LCTRL
	case "LSHIFT":
		m = LSHIFT
	case "LALT":
		m = LALT
	case "LMETA":
		m = LMETA
	case "RCTRL":
		m = RCTRL
	case "RSHIFT":
		m = RSHIFT
	case "RALT":
		m = RALT
	case "RMETA":
		m = RMETA
	}

	t.StringModifier = stringModifier{StringRef: submatch[4], Modifier: m}
	t.CharacterModifier = characterModifier{CharRef: ' ', Modifier: NOMOD}

	return t
}
