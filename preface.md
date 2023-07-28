# Preface

## Purpose

My passion for programming began at a young age, and my first encounter with solving a programming challenge at ten years old left me exhilarated. This excitement ignited a lifelong pursuit of coding. As I progressed to studying computer science at university, my goal was to delve into more profound and efficient coding techniques. Among the subjects that were offered, the ones that resonated most with this objective were Data Structures and Algorithms.

Algorithms are a ubiquitous part of today's formal education. The step-by-step procedures that students take to solve equations is an example that also relates to the origin of the term itself. More than a thousand years ago Khwarazmi, a Persian polymath and mathematician wrote a book on Algebra discussing solving equations among other things. The Latin translation of his name "Algoritmi" eventually gave rise to the term "Algorithm".

Today, Data Structures and Algorithms form the foundational knowledge in computer science and related professions. Mastery of these concepts empowers individuals to analyze, design, and discuss solutions for intricate problems, making them vital skills assessed in technical job interviews.

Learning Data Structures and Algorithms is a journey that demands extensive reading, contemplation, and practice. Furthermore, this subject is also susceptible to being easily forgotten over time. To tackle this challenge, I embarked on compiling concise notes and saving the problems I solved for my personal use.

This collection became a valuable resource, and I sought to share its benefits with others. While contemplating publishing a book, I realized that offering a public code repository has several advantages. Primarily a codebase unlike a book is executable and testable, therefore much less prone to errors. Moreover, a public code repository can be continually improved by the community whereas a book is fixed in time.

## Choosing Go as the Language

Go initially aimed to be a successor for C++, an ambitious goal given that operating systems like Linux and Windows and programming languages like Java and Python are written in C and C++. Go has matured over the past decade into a robust and solid option, finding its niche particularly in systems engineering where reliability and performance are essential.

Go is simple. This simplicity is evident in several aspects. The language features few [keywords](https://go.dev/ref/spec#Keywords) that programmers from any background would likely recognize. Its syntax is minimal, offering a single type of loop, and no brackets are needed for conditions in if or for statements. The Go standard library, written in Go and assembly language, is frequently used not only to understand the internals of Go but also to learn how to write idiomatic Go code. Go maintains backward compatibility for its standard API, and significant changes require approval from the three original authors, ensuring that the language does not become a "Frankenstein".

Go is compiled into machine language for multiple operating systems. This eliminates the need to ship the code itself like Ruby or download a virtual machine like Java and runtime libraries like Node.js, making it convenient for distribution and execution.

Go excels in handling concurrency, allowing seamless utilization of multiple CPU cores. While being faster than dynamic programming languages like Python, Go still maintains garbage collection for memory management which theoretically makes it slower than languages with no garbage collection like Rust.

Being a strongly typed language, Go ensures code stability, while offering the modern conveniences found in modern dynamic scripting languages like Node.js.

Go is a new language, and it comes with built-in tooling for testing, fuzzing, benchmarking, formatting, vetting, race detection, and package management, all of which are secondary or external additions to much older languages like C, C++ and Java. It also comes with built-in understanding of more modern concepts like code repositories.

Go is well organized. Go is strict in how a Go project should be formatted and organized in terms of directories, and files. This makes all Go projects similar to each other and therefore easy to navigate.

## Common Objections to Go

1. Fast Development: Some developers initially find it slower to work with Go due to its unique syntax and conventions. However, once developers become familiar with the language and its ecosystem, Go allows for fast and efficient code writing. While it may require slightly more time to write code in a strongly typed and compiled language like Go, the trade-off is increased code stability and fewer issues in production.
2. Too much error handling: Handling errors in Go is a deliberate design choice for ensuring program stability. Go developers invest time in ensuring their programs handle errors appropriately during execution. This practice leads to more robust and reliable code.
3. No generics: This is no longer the case. Go has support for generics starting from version 1.18.
4. Not suitable for limited environments: While Go is a general purpose programming language, it may not be suitable for all environments. For example, it may not be suitable for embedded systems without an operating system. Additionally, Go might not be the best choice for creating user interfaces in frontend web or mobile development. However, Go can be used to create WebAssembly and libraries for other languages.

## Target Audience

This resource is suitable for those with at least a beginner's understanding of Go and exposure to Data Structures and Algorithms. It is designed to be concise and to-the-point, assuming readers already possess basic knowledge of both subjects.

For individuals unfamiliar with Go, the [Tour of Go](https://go.dev/tour/welcome/1) is an excellent starting point. As for Data Structures and Algorithms, taking an instructor-led course is recommended for an introduction.

## How to use

* Read the README.md file of each subject.
* Solve the rehearsal problems.
* Read the problem description.
* Copy the contents of each `*_test.go` file into [Go Playground](https://go.dev/play/) or your favorite development environment.
* Change the first line from `package SOME_PACKAGE_NAME` to `package main`.
* Look at the sample input in the `*_test.go` file.
* Create a solution function by looking at the function signature in the test file.
* Implement your own solution.
* Review the provided solution and compare it to your own.
* Come back and rehearse.

## Note on Problem Sources

While I created a few rehearsal problems in this collection, the majority are well-known challenges found in multiple books and online resources. Due to their age and widespread use, pinpointing their origins can be challenging. The additional resources listed next contain many of these problems, and even they do not mention the origin of these problems in many cases.

## Additional Resources

* [Introduction to Algorithms](https://amzn.to/3q6S4TK)
* [Data Structures and Algorithm Analysis in Java.](https://amzn.to/3DugGsP)
* [The Big Book of Coding Interviews in Python](https://amzn.to/3rC0ToT)
* [Daily Coding Problem](https://amzn.to/3OvDjmH)
* [LeetCode](https://leetcode.com/)
