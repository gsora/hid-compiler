# `hid-compiler` [![Go Report Card](https://goreportcard.com/badge/github.com/gsora/hid-compiler)](https://goreportcard.com/report/github.com/gsora/hid-compiler)

An experimental "compiler".
Write and compile a small language to Linux's USB HID keyboard scancodes.

## Why

Sometimes, you need a quick solution to write USB HID scancodes, and the process **will** be painful.

Why write scancodes by hand when you can create a small language and a compiler to do this job for you instead? `;)`

I'm not a compiler designer, or a language designer, I'm just a guy who likes to tinker with stuff -- I know this is far from being called a real compiler.

## Syntax

`hid-compiler` will accept any **ASCII symbol**.

An **ASCII string** is defined by one or more ASCII symbol:

 - "a" is an accepted string
 - " " is an accepted string
 - "" is not an accepted string (empty, thus does not contain any ASCII symbol)
 - "hello" is an accepted string
 
### Key modifiers
 
`hid-compiler` supports the bare minimum *key modifiers* available on a standard keyboard:

 - `LCTRL` (left control)
 - `LSHIFT` (left shift)
 - `LALT` (left alt)
 - `LMETA` (left Super/Windows/Command key)
 - `RCTRL` (right control)
 - `RSHIFT` (right shift)
 - `RALT` (right alt)
 - `RMETA` (right Super/Windows/Command key)
 
`hid-compiler` supports **string modifiers** and **character modifiers**.

### String modifiers

A **string modifier** applies a key modifier operation on a string, and it's defined as follows:

```
[key-modifier]>STRING
```

Where "STRING" is a well-formed ASCII string as defined previously.

### Character modifiers

A **character modifier** applies a key modifier operation on the character on the rightmost side of the declaration; it's defined as follows:

```
[key-modifier]:STRING
```

The modifier operation will be applied only on the first character of the "STRING" ASCII string (the "S").

```
ST[key-modifier]:RING
```

The modifier operation will be applied on the "R" character of the "STRING" ASCII string.

## Intuitive idea and compiling phases

The input string is parsed as a *linked list* composed of `tokens`.

A `token` holds the original, un-parsed string, the parsed string without any modifier, either a `stringModifier` or a `characterModifier`.

Both `stringModifier` and `characterModifier` holds a reference to the string or character they have to modify, and a `modifier`.

A `modifier` is a simple data type, mostly created for stylish purposes -- I could have used a simple `int`, but that would have been ugly.


While a `stringModifier` is conceptually easier to understand -- it only holds a string and a `modifier` -- `characterModifier` is more complex because of the way it binds to a character.

This kind of modifier can appear everywhere in the string, thus I needed to memorize the sub-string before the modifier (`Left`) and after (`Right`), including the character to modify.

The real action starts in the `Compile(s string)`, in `compiler/Compile.go`.

This function starts by creating a linked list; each word divided a space will be tokenized by `tokenize(s string)` and put into the list.

`tokenize(s string)` calls both `matchStringModifier(s string)` and `matchCharModifier(s string)` to check if the string contains a character modifier, or a string modifier.

Each of these function will return a `token` containig one of the two modifier fields non-nil.

If both functions don't return, it means the string doesn't contain any modifier, and thus can be parsed withouth any special operation-- the resulting `token` contains the modifier fields set to `nil`.

After the tokenizing phase, it's now time to actually translate the tokens to HID scancodes.

Since there are 3 types of strings currently built in, 3 functions has been written to accomplish the translation:

 + `generateHIDPayloadForString(s stringModifier)`
 + `generateHIDPayloadForCharacter(c characterModifier)`
 + `generateHIDPayloadForCharacter(s string)`
 + `generateHIDPayloadForSpace()`
 
To make the translation easier, the `hidPayload` type has been defined.
`hidPayload` holds a reference to the Modifier to apply to a character (which is a `string` anyway), and the character itself.

A `String()` method for `hidPayload` has been defined too, which returns the hex-formatted scancode.

In a nutshell, these function will create the hex-formatted string by concatenating the `String()` output of all the `hidPayload` created with the content of the `characterModifier` or `stringModifier` (if any) passed as argument.

`generateHIDPayloadForCharacter(c characterModifier)` is a little bit different, because it has to account for the left and right part of the original string.

`generateHIDPayloadForSpace()` is much simpler: it just creates a "space" character, encoded.

These functions will be called while cycling through the list, and their output will be concatenated.
