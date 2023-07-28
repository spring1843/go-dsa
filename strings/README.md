# String

A string is a ubiquitous data structure, typically a built-in data type in programming languages. However, beneath the surface, strings are essentially arrays of characters that enable textual data storage and manipulation.

## Implementation

In Go, strings are a data type. Behind the scenes strings are a slice of bytes. The `strings` package provides several useful convenience functions. Examples include:

* [Index](https://golang.org/pkg/strings/#Index), [Contains](https://golang.org/pkg/strings/#Contains), [HasPrefix](https://golang.org/pkg/strings/#HasPrefix), [HasSuffix](https://golang.org/pkg/strings/#HasSuffix)
* [Split](https://golang.org/pkg/strings/#Split), [Fields](https://golang.org/pkg/strings/#Split), [Join](https://golang.org/pkg/strings/#Join)
* [Repeat](https://golang.org/pkg/strings/#Repeat)
* [ReplaceAll](https://golang.org/pkg/strings/#ReplaceAll)
* [Title](https://golang.org/pkg/strings/#Title), [ToLower](https://golang.org/pkg/strings/#ToLower), [ToUpper](https://golang.org/pkg/strings/#ToUpper)
* [Trim](https://golang.org/pkg/strings/#Trim), [TrimSpace](https://golang.org/pkg/strings/#TrimSpace), [TrimSuffix](https://golang.org/pkg/strings/#TrimSuffix), [TrimPrefix](https://golang.org/pkg/strings/#TrimPrefix)

When a string is iterated in Go using the `range` keyword, every element becomes a [rune](https://blog.golang.org/strings#TOC_5.) which is an alias for the type `int32`. If the code being written works with many single-character strings, it is better to define variables and function parameters as `rune`. The following code shows how to iterate through a string.

```Go
package main

import (
	"fmt"
)

func main() {
	for i, r := range "abcd" {
		fmt.Printf("Char #%d %q has value %d\n", i, string(r), r)
	}
}
```

## Complexity

Strings have the same complexity as [arrays](../array/) and slices in Go.

Unlike many other programming languages, [regular expressions](https://golang.org/pkg/regexp/) are guaranteed to have O(n) time complexity in Go, allowing efficient string pattern matching.

## Application

Strings store words, characters, sentences, etc.

## Rehearsal

* [The Longest Dictionary Word Containing Key](./longest_dictionary_word_test.go), [Solution](./longest_dictionary_word.go)
* [Look and Tell](./look_and_tell_test.go), [Solution](./look_and_tell.go)
* [In Memory Database](./in_memory_database_test.go), [Solution](./in_memory_database.go)
* [Number in English](./number_in_english_test.go), [Solution](./number_in_english.go)
* [Reverse Vowels In a String](./reverse_vowels_test.go), [Solution](./reverse_vowels.go)
* [Roman Numerals](./roman_numerals_test.go), [Solution](./roman_numerals.go)
* [Longest Substring of Two Unique Characters](./longest_substring_test.go), [Solution](./longest_substring.go)
