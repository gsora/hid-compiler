package modifiers

import (
	"bytes"
	"fmt"
	"strings"
)

type HIDPayload struct {
	Modifier   Modifier
	Character0 string
}

const HIDEmpty = "\\0\\0\\0\\0\\0\\0\\0\\0"

func (h *HIDPayload) String() string {
	if h.Modifier == NOMOD {
		return fmt.Sprintf("\\0\\0\\%s\\0\\0\\0\\0\\0", h.Character0)
	}
	return fmt.Sprintf("\\%s\\0\\%s\\0\\0\\0\\0\\0", scancodeForModifier(h.Modifier.String()), h.Character0)
}

func GenerateHIDPayloadForString(s StringModifier) string {
	var ret bytes.Buffer

	for count, character := range s.StringRef {
		st := HIDPayload{s.Modifier, scancodeForModifier(strings.ToUpper(string(character)))}
		ret.WriteString(st.String())
		if count != len(s.StringRef) {
			ret.WriteString(HIDEmpty)
		}
	}

	return ret.String()

}

func GenerateHIDPayloadForCharacter(c CharacterModifier) string {
	var ret bytes.Buffer

	ret.WriteString(GenerateHIDPayloadForStandardString(c.Left))

	st := HIDPayload{c.Modifier, scancodeForModifier(strings.ToUpper(string(c.CharRef)))}
	ret.WriteString(st.String())
	ret.WriteString(HIDEmpty)

	ret.WriteString(GenerateHIDPayloadForStandardString(c.Left[1:]))

	return ret.String()

}

func GenerateHIDPayloadForStandardString(s string) string {
	var ret bytes.Buffer

	for count, character := range s {
		st := HIDPayload{NOMOD, scancodeForModifier(strings.ToUpper(string(character)))}
		ret.WriteString(st.String())
		if count != len(s) {
			ret.WriteString(HIDEmpty)
		}
	}

	return ret.String()

}
