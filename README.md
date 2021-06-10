# Golang Toy Robot

I am using this as a tool to learn golang.

The app is a CLI that models a robot moving about on a 5x5 table.

## Valid commands

```
PLACE 1,1,NORTH
MOVE
LEFT
RIGHT
REPORT
```

## Status

1. Get something that works, without worrying too much about go idioms.
2. CURRENT - Update to be idiomatic go code, unless there is a good reason.
3. Polish and finish.

### Questions

#### Should I use err more?

For example, my parsing functions return `nil` to indiciate a failure to parse. Should this be the `err` pattern instead which seems standard in go?

#### Is immutable style good?

I do a lot of pass by value, which duplicates the value to pass it into the function and will not mutate the original value. This seems like a good thing, but will have performance implications.

Is idiomatic go using pass by ref and attempting safe mutation? Should I change to this style?

#### Code structure?

The code structure is based on my Haskell toy robot. However, I think perhaps a feature based package breakdown might be nice, ie parsing + processing.

#### Constant hacks?

I am using some non-standard? techniques to avoid modelling some types as strings. Ie. the commands and direction. Not sure how unidiomatic these are.
