# Data Structures and Algorithms in Go

[![tests](https://github.com/spring1843/go-dsa/actions/workflows/tests.yaml/badge.svg)](https://github.com/spring1843/go-dsa/actions/workflows/tests.yaml)
[![GitHub License](https://img.shields.io/badge/License-Apache%202.0-ff69b4.svg)](https://github.com/aws/karpenter/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/spring1843/go-dsa)](https://goreportcard.com/report/github.com/spring1843/go-dsa)
[![Coverage Report](https://coveralls.io/repos/github/spring1843/go-dsa/badge.svg?branch=main)](https://coveralls.io/github/spring1843/go-dsa?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/spring1843/go-dsa.svg)](https://pkg.go.dev/github.com/spring1843/go-dsa)

The primary objective of this project is to offer a valuable learning and practicing resource for Data Structures and Algorithms in the Go programming language. The project aims to provide concise and executable materials for individuals seeking to develop their knowledge in this area.

This learning platform offers several advantages over traditional learning methods such as books and videos:

* It continues to evolve and improve over time, providing learners with community contributed and reviewed content
* It is executable on various platforms, which minimizes the possibility of errors in the provided solutions
* It provides solutions for all problems with 100% test coverage to prove the accuracy of each solution
* It allows learners to test their skills in their preferred programming environment, making it more convenient and flexible
* It is freely available to anyone interested in enhancing their knowledge in the subject area.

## Table of Contents

* Time and Space Analysis
* Data Structures
    * [Arrays](./array)
    * [Strings](./string)
    * [Linked Lists](./linkedlist)
    * [Stacks](./stack)
    * [Queues](./queue)
    * [Hash Tables](./hashtable)
    * [Trees](./tree)
    * [Heaps](./heap)
* Algorithms
    * [Recursion](./recursion)
    * [Divide and Conquer](dnc)
    * [Bit Manipulation](./bit)
    * [Backtracking](./backtracking)
    * [Graphs](./graph)
    * [Greedy Algorithms](./greedy)
    * [Dynamic Programming](./dp)

## How to use

* Read the README.md file of each subject
* Copy the contents of each `*_test.go` file into [Go Playground](https://go.dev/play/)
* Implement your own solution and then compare it with the provided solution
* Read the provided solution and compare it to your own solution

To run examples locally [install Go](https://go.dev/doc/install), and run `go get -u github.com/spring1843/go-dsa`.

## Outline

For each area of study the goal is to provide a README.md file and at least 5 solved problems.

Each README.md will contain the following sections:

* **Implementation**, How the data structure or algorithm can generally be implemented with code examples in Go
* **Complexity**, Time and space complexity analysis of typical data structure operations
* **Application**, Problems commonly solved using the data structure
* **Rehearsal**, Example problem to practice along with links to tests that provide 100% coverage and example input and outputs
