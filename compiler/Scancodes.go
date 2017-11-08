package compiler

import "log"

func scancodeForModifier(m string) string {
	if val, ok := scancodes[m]; ok {
		return val
	} else {
		log.Fatal("NOT FOUND:", m == "\n")
		return " "
	}
}

var scancodes = map[string]string{
	"LCTRL":              "\x01",
	"LSHIFT":             "\x02",
	"LALT":               "\x04",
	"LMETA":              "\x08",
	"RCTRL":              "\x10",
	"RSHIFT":             "\x20",
	"RALT":               "\x40",
	"RMETA":              "\x80",
	"NONE":               "\x00",
	"ERR_OVF":            "\x01",
	"A":                  "\x04",
	"B":                  "\x05",
	"C":                  "\x06",
	"D":                  "\x07",
	"E":                  "\x08",
	"F":                  "\x09",
	"G":                  "\x0a",
	"H":                  "\x0b",
	"I":                  "\x0c",
	"J":                  "\x0d",
	"K":                  "\x0e",
	"L":                  "\x0f",
	"M":                  "\x10",
	"N":                  "\x11",
	"O":                  "\x12",
	"P":                  "\x13",
	"Q":                  "\x14",
	"R":                  "\x15",
	"S":                  "\x16",
	"T":                  "\x17",
	"U":                  "\x18",
	"V":                  "\x19",
	"W":                  "\x1a",
	"X":                  "\x1b",
	"Y":                  "\x1c",
	"Z":                  "\x1d",
	"1":                  "\x1e",
	"2":                  "\x1f",
	"3":                  "\x20",
	"4":                  "\x21",
	"5":                  "\x22",
	"6":                  "\x23",
	"7":                  "\x24",
	"8":                  "\x25",
	"9":                  "\x26",
	"0":                  "\x27",
	"ENTER":              "\x28",
	"ESC":                "\x29",
	"BACKSPACE":          "\x2a",
	"TAB":                "\x2b",
	" ":                  "\x2c",
	"MINUS":              "\x2d",
	"EQUAL":              "\x2e",
	"LEFTBRACE":          "\x2f",
	"RIGHTBRACE":         "\x30",
	"BACKSLASH":          "\x31",
	"HASHTILDE":          "\x32",
	";":                  "\x33",
	"'":                  "\x34",
	"GRAVE":              "\x35",
	",":                  "\x36",
	".":                  "\x37",
	"SLASH":              "\x38",
	"CAPSLOCK":           "\x39",
	"F1":                 "\x3a",
	"F2":                 "\x3b",
	"F3":                 "\x3c",
	"F4":                 "\x3d",
	"F5":                 "\x3e",
	"F6":                 "\x3f",
	"F7":                 "\x40",
	"F8":                 "\x41",
	"F9":                 "\x42",
	"F10":                "\x43",
	"F11":                "\x44",
	"F12":                "\x45",
	"SYSRQ":              "\x46",
	"SCROLLLOCK":         "\x47",
	"PAUSE":              "\x48",
	"INSERT":             "\x49",
	"HOME":               "\x4a",
	"PAGEUP":             "\x4b",
	"DELETE":             "\x4c",
	"END":                "\x4d",
	"PAGEDOWN":           "\x4e",
	"RIGHT":              "\x4f",
	"LEFT":               "\x50",
	"DOWN":               "\x51",
	"UP":                 "\x52",
	"NUMLOCK":            "\x53",
	"KPSLASH":            "\x54",
	"KPASTERISK":         "\x55",
	"KPMINUS":            "\x56",
	"KPPLUS":             "\x57",
	"KPENTER":            "\x58",
	"KP1":                "\x59",
	"KP2":                "\x5a",
	"KP3":                "\x5b",
	"KP4":                "\x5c",
	"KP5":                "\x5d",
	"KP6":                "\x5e",
	"KP7":                "\x5f",
	"KP8":                "\x60",
	"KP9":                "\x61",
	"KP0":                "\x62",
	"KPDOT":              "\x63",
	"102ND":              "\x64",
	"COMPOSE":            "\x65",
	"POWER":              "\x66",
	"KPEQUAL":            "\x67",
	"F13":                "\x68",
	"F14":                "\x69",
	"F15":                "\x6a",
	"F16":                "\x6b",
	"F17":                "\x6c",
	"F18":                "\x6d",
	"F19":                "\x6e",
	"F20":                "\x6f",
	"F21":                "\x70",
	"F22":                "\x71",
	"F23":                "\x72",
	"F24":                "\x73",
	"OPEN":               "\x74",
	"HELP":               "\x75",
	"PROPS":              "\x76",
	"FRONT":              "\x77",
	"STOP":               "\x78",
	"AGAIN":              "\x79",
	"UNDO":               "\x7a",
	"CUT":                "\x7b",
	"COPY":               "\x7c",
	"PASTE":              "\x7d",
	"FIND":               "\x7e",
	"MUTE":               "\x7f",
	"VOLUMEUP":           "\x80",
	"VOLUMEDOWN":         "\x81",
	"KPCOMMA":            "\x85",
	"RO":                 "\x87",
	"KATAKANAHIRAGANA":   "\x88",
	"YEN":                "\x89",
	"HENKAN":             "\x8a",
	"MUHENKAN":           "\x8b",
	"KPJPCOMMA":          "\x8c",
	"HANGEUL":            "\x90",
	"HANJA":              "\x91",
	"KATAKANA":           "\x92",
	"HIRAGANA":           "\x93",
	"ZENKAKUHANKAKU":     "\x94",
	"KPLEFTPAREN":        "\xb6",
	"KPRIGHTPAREN":       "\xb7",
	"LEFTCTRL":           "\xe0",
	"LEFTSHIFT":          "\xe1",
	"LEFTALT":            "\xe2",
	"LEFTMETA":           "\xe3",
	"RIGHTCTRL":          "\xe4",
	"RIGHTSHIFT":         "\xe5",
	"RIGHTALT":           "\xe6",
	"RIGHTMETA":          "\xe7",
	"MEDIA_PLAYPAUSE":    "\xe8",
	"MEDIA_STOPCD":       "\xe9",
	"MEDIA_PREVIOUSSONG": "\xea",
	"MEDIA_NEXTSONG":     "\xeb",
	"MEDIA_EJECTCD":      "\xec",
	"MEDIA_VOLUMEUP":     "\xed",
	"MEDIA_VOLUMEDOWN":   "\xee",
	"MEDIA_MUTE":         "\xef",
	"MEDIA_WWW":          "\xf0",
	"MEDIA_BACK":         "\xf1",
	"MEDIA_FORWARD":      "\xf2",
	"MEDIA_STOP":         "\xf3",
	"MEDIA_FIND":         "\xf4",
	"MEDIA_SCROLLUP":     "\xf5",
	"MEDIA_SCROLLDOWN":   "\xf6",
	"MEDIA_EDIT":         "\xf7",
	"MEDIA_SLEEP":        "\xf8",
	"MEDIA_COFFEE":       "\xf9",
	"MEDIA_REFRESH":      "\xfa",
	"MEDIA_CALC":         "\xfb",
}
