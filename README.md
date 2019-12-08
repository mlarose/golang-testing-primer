Golang testing primer
---------------------

# Introduction

This project was written as part of a training session aiming to provide guidance on how 
write unit tests in Golang. 


# Get started

## Requisites

This tutorial examples require docker and golang to be installed. 
It was tested using Go 1.13 for darwin.

## CLI help

The included makefile aims to facilitate trying out the examples

- make test: executes all the test
- make bench: executes all the benchmark
- make godoc: runs a dockerized version of godoc that you can open locally http://localhost:6060

## Tip

The tests are separated in arrange / act / assert sections to conveniently expose to the
reader what's the preconditions or setup for the test, the functionality being tested, and the
expected results of the test. This is also a good pattern to follow to ensure your tests are 
readable and maintainable.

## Section

This tutorial is organized in three sections. Feel free to read the examples and comments.

### Basics - Introduction to the Golang testing tooling

- [example code](./basics/add.go) 
- [tests](./basics/add_test.go)

### Assert - Introduction to stretchr/testify/assert

- [example code](./assert/sub.go)
- [tests](./assert/sub_test.go)

### Mocking - Introduction to mocking with stretchr/testify/mock

- [example code](./mocking/unit.go)
- [external dependency impl](./mocking/external-dependency.go)
- [tests](./mocking/unit_test.go)
- [roll our own mock example](./mocking/mock_1_test.go)
- [mock using testify/mock](./mocking/mock_2_test.go)
- [Base64.Encoding partial mock](./mocking/mock-encoder_test.go)

