# Introduction

A Go (Golang) generator.  Combines .c and .go file to produce a new
.go file for `cgo`.  This should allow a developer to use their editor
of choice to edit independently the .c and .go files instead of a
single large comment above a .go import "C" line.

## Usage

At the command line for use with a makefile perhaps:

```
gocombine --cc some-c.cpp --go some-go.go --out generated.gen.go
```

Embedded in Go code:

```
//go:generate gocombine --cc some-c.cpp --go some-go.go --out
generated.gen.go
```

## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt
