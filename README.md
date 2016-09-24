# drgomesp/cargo

> *An efficient and robust Go dependency injection container* – by **[Daniel Ribeiro](https://github.com/drgomesp)**

[![License](https://img.shields.io/badge/liecense-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/drgomesp/cargo?status.svg)](https://godoc.org/github.com/drgomesp/cargo) [![Current Project](https://img.shields.io/badge/target%20release-0.1.0-ff69cc.svg)](https://github.com/drgomesp/cargo/projects/1)
 [![Go Report Card](https://goreportcard.com/badge/github.com/drgomesp/cargo)](https://goreportcard.com/report/github.com/drgomesp/cargo) [![Build Status](https://travis-ci.org/drgomesp/cargo.svg?branch=master)](https://travis-ci.org/drgomesp/cargo) 

 

___

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [How it Works](#how-it-works)
  1. [Builder](#builder)

### Introduction

**cargo** is a library that provides a powerful way of handling objects and 
 their dependencies, by using the *Container*. The container works
 by implementing the [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) 
 pattern via constructor injection, resulting in explicit dependencies and the achievement 
 of the [Inversion of Control](https://en.wikipedia.org/wiki/Inversion_of_control) principle.

### Installation

```bash
$ go get github.com/drgomesp/cargo
```

### How it Works

#### Builder

The `container.Container` is what provides an API to register and retrieve services

Defining services is very simple:

```go
type Foo struct {} 

container := container.NewContainer()
container.Register("foo", definition.NewDefinition(&Foo{}))
```

If you prefer not to use the composite literal expression, you can define a constructor function and use it as a literal:

```go
type Foo struct {} 
func NewFoo() *Foo {
    return &Foo{}
}

container.Register("foo", definition.NewDefinition(NewFoo))
```

To check if a service definition exists, simply do:

```go
container.HasDefinition("foo")
```

And to get a definition back to work with:

```go
if def, err := container.GetDefinition("foo"); err != nil {
    fmt.Println(def.ID)
}
```

