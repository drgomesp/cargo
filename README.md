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

by **[Daniel Ribeiro](twitter.com/drgomesp)**

[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/badge/liecense-MIT-blue.svg 
[docs]: https://godoc.org/github.com/drgomesp/cargo
[docs_badge]: https://godoc.org/github.com/drgomesp/cargo?status.svg
[release]: https://github.com/drgomesp/cargo/releases
[release_badge]: https://img.shields.io/github/release/drgomesp/cargo.svg
[report]: https://goreportcard.com/report/github.com/drgomesp/cargo
[report_badge]: https://goreportcard.com/badge/github.com/drgomesp/cargo
[build]: https://travis-ci.org/drgomesp/cargo
[build_badge]: https://travis-ci.org/drgomesp/cargo.svg?branch=develop
[coverage]: https://coveralls.io/github/drgomesp/cargo?branch=develop
[coverage_badge]: https://coveralls.io/repos/github/drgomesp/cargo/badge.svg?branch=develop