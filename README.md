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
