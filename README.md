<img  width="250" align="right" src="http://2.bp.blogspot.com/-4Yy4UKNvlic/UDacAxBt--I/AAAAAAAAEwU/F-IQc8NGejo/s1600/semi.png" />

# drgomesp/cargo

> *An efficient and robust Go dependency injection container* – by **[Daniel Ribeiro](https://github.com/drgomesp)**

[![License][license_badge]][license]
[![GoDoc][docs_badge]][docs]
[![Latest Release][release_badge]][release]
[![Go Report][report_badge]][report]
[![Build Status][build_badge]][build]
[![Coverage Status][coverage_badge]][coverage]

___

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Getting Started](#getting-started)

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

### Getting Started

#### Creating/Registering Services

There are two main methods used to define services in the container: `container.Register`
and `container.Set`. The first one assumes you already have a pointer to an object instance
to work with, the second needs a **definition**.

Suppose you have an object:

```go
type HttpClient struct {}

client := new(HttpClient)
```

To define that as a service, all you need to do is:

```go
dic := container.NewContainer()
dic.Set("http.client", client)
```

From now on, whenever you need to work with the http client service, you can simply do:

```go
if s, err := dic.Get("http.client"); err != nil {
    panic(err)
}

client := s.(*HttpClient) // the type assertion is required
```

Or, if you do not need errors handling and panic is fine, you can get the same behavior with short synthax:

```go
client := dic.MustGet("http.client").(*HttpClient)
```

by **[Daniel Ribeiro](https://twitter.com/drgomesp)**

[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square

[docs]: https://godoc.org/github.com/drgomesp/cargo
[docs_badge]: https://img.shields.io/badge/godoc-reference-9891ff.svg?style=flat-square

[release]: https://github.com/drgomesp/cargo/releases
[release_badge]: https://img.shields.io/github/release/drgomesp/cargo.svg?style=flat-square

[report]: https://goreportcard.com/report/github.com/drgomesp/cargo
[report_badge]: https://goreportcard.com/badge/github.com/drgomesp/cargo?style=flat-square

[build]: https://travis-ci.org/drgomesp/cargo
[build_badge]: https://img.shields.io/travis/drgomesp/cargo.svg?style=flat-square

[coverage]: https://coveralls.io/github/drgomesp/cargo?branch=develop
[coverage_badge]: https://img.shields.io/coveralls/drgomesp/cargo.svg?style=flat-square
