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
