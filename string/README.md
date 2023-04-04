# String

Strings are an ubiquitous data structure that typically manifest as built-in data types within programming languages. However, beneath the surface, they are essentially arrays of characters that enable the storage and manipulation of textual data.

## Implementation

In Go strings are a data type. Behind the scene strings are a slice of bytes. The `strings` package provides several useful convenience functions. Examples include:

* [Index](https://golang.org/pkg/strings/#Index), [Contains](https://golang.org/pkg/strings/#Contains), [HasPrefix](https://golang.org/pkg/strings/#HasPrefix), [HasSuffix](https://golang.org/pkg/strings/#HasSuffix)
* [Split](https://golang.org/pkg/strings/#Split), [Fields](https://golang.org/pkg/strings/#Split), [Join](https://golang.org/pkg/strings/#Join)
* [Repeat](https://golang.org/pkg/strings/#Repeat)
* [ReplaceAll](https://golang.org/pkg/strings/#ReplaceAll)
* [Title](https://golang.org/pkg/strings/#Title), [ToLower](https://golang.org/pkg/strings/#ToLower), [ToUpper](https://golang.org/pkg/strings/#ToUpper)
* [Trim](https://golang.org/pkg/strings/#Trim), [TrimSpace](https://golang.org/pkg/strings/#TrimSpace), [TrimSuffix](https://golang.org/pkg/strings/#TrimSuffix), [TrimPrefix](https://golang.org/pkg/strings/#TrimPrefix)

When you iterate through a string in Go using the `range` keyword, every element becomes a [rune](https://blog.golang.org/strings#TOC_5.) which is the same as `int32`. If you are writing code that deals with characters it's often easier to define your function parameters and variables as rune. The following code shows how to iterate through a string.

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

The Go standard library includes a [regular expressions](https://golang.org/pkg/regexp/) package that unlike most other languages comes with a guarantee to run in O(n) time.

## Application

Strings are used to store words, characters, sentences and etc.

## Rehearsal

### The longest Dictionary Word Containing Key

Given a key like "car", and a dictionary like {"rectify", "race", "archeology", "racoon"} return the longest dictionary word that contains every letter of the key like "archeology". [Solution](longest_dictionary_word.go), [Tests](longest_dictionary_word_test.go)

### Look and Tell

Given a depth, return the output of look and tell an algorithm where each line reads the last line. For example "1" is read as "11" (one one), and "11" is read as "21" (two ones). [Solution](look_and_tell.go), [Tests](look_and_tell_test.go)

### In Memory Database

Write an in memory database that stores string key value pairs and supports SET, GET, EXISTS, and UNSET commands. It should also allow transactions with BEGIN, COMMIT and ROLLBACK commands. [Solution](in_memory_database.go), [Tests](in_memory_database_test.go)

### Number in English

Given a number like 34, return how the number  would be read in English e.g. (Thirty Four) for integers smaller than one Trillion. [Solution](number_in_english.go), [Tests](number_in_english_test.go)

### Reverse Vowels In a String

Given a string e.g. "coat", reverse the order in which vowels appear e.g. "caot". [Solution](reverse_vowels.go), [Tests](reverse_vowels_test.go)

### Longest Substring of Two Unique Characters

Given a string like "aabbc" return the longest substring of two unique characters like "aabb". [Solution](longest_substring.go), [Test](longest_substring_test.go)
