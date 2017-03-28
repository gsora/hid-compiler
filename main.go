package main

import (
	"fmt"
	"github.com/gsora/hid-compiler/compiler"
	"github.com/gsora/hid-compiler/modifiers"
	//"github.com/gsora/hid-compiler/types"
)

var s1 = "this is a string"
var s2 = "this is a [LCTRL]>string too"
var s3 = "th[LSHIFT]:is is a string [LCTRL]>again"

func main() {

	l := compiler.Compile(s3)

	i := l.GetFirst()

	for i.Next != nil {
		d := i.Data
		fmt.Println("Found element:")
		fmt.Println("Raw string:", d.RawString)
		fmt.Println("Real string:", d.RealString)
		if d.StringModifier.Modifier != modifiers.NOMOD {
			fmt.Println("String modifier:", d.StringModifier.Modifier.String())
			fmt.Println("String to modify:", d.StringModifier.StringRef)
			fmt.Println(modifiers.GenerateHIDPayloadForString(d.StringModifier))
		} else if d.CharacterModifier.Modifier != modifiers.NOMOD {
			fmt.Println("Char modifier:", d.CharacterModifier.Modifier.String())
			fmt.Println("Char to modify:", string(d.CharacterModifier.CharRef))
			fmt.Println("Char left:", string(d.CharacterModifier.Left))
			fmt.Println("Char right:", string(d.CharacterModifier.Right))
			fmt.Println(modifiers.GenerateHIDPayloadForCharacter(d.CharacterModifier))
		} else {
			fmt.Println(modifiers.GenerateHIDPayloadForStandardString(d.RealString))
		}

		fmt.Println("")
		i = *i.Next
	}

}
