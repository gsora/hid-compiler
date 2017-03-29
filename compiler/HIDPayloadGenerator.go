package compiler

import (
	"bytes"
	"fmt"
	"strings"
)

type hidPayload struct {
	Modifier   modifier
	Character0 string
}

const hidEmpty = "\\0\\0\\0\\0\\0\\0\\0\\0"

func (h *hidPayload) String() string {
	if h.Modifier == NOMOD {
		return fmt.Sprintf("\\0\\0\\%s\\0\\0\\0\\0\\0", h.Character0)
	}
	return fmt.Sprintf("\\%s\\0\\%s\\0\\0\\0\\0\\0", scancodeForModifier(h.Modifier.String()), h.Character0)
}

func generateHIDPayloadForString(s stringModifier) string {
	var ret bytes.Buffer

	for count, character := range s.StringRef {
		st := hidPayload{s.Modifier, scancodeForModifier(strings.ToUpper(string(character)))}
		ret.WriteString(st.String())
		if count != len(s.StringRef) {
			ret.WriteString(hidEmpty)
		}
	}

	return ret.String()

}

func generateHIDPayloadForCharacter(c characterModifier) string {
	var ret bytes.Buffer

	ret.WriteString(generateHIDPayloadForStandardString(c.Left))

	st := hidPayload{c.Modifier, scancodeForModifier(strings.ToUpper(string(c.CharRef)))}
	ret.WriteString(st.String())
	ret.WriteString(hidEmpty)

	ret.WriteString(generateHIDPayloadForStandardString(c.Left[1:]))

	return ret.String()

}

func generateHIDPayloadForStandardString(s string) string {
	var ret bytes.Buffer

	for count, character := range s {
		st := hidPayload{NOMOD, scancodeForModifier(strings.ToUpper(string(character)))}
		ret.WriteString(st.String())
		if count != len(s) {
			ret.WriteString(hidEmpty)
		}
	}

	return ret.String()

}
