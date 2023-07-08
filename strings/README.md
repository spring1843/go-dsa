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

* [Solution](longest_dictionary_word.go), [The longest Dictionary Word Containing Key](longest_dictionary_word_test.go)
* [Look and Tell](look_and_tell_test.go), [Solution](look_and_tell.go)
* [In Memory Database](in_memory_database_test.go), [Solution](in_memory_database.go)
* [Number in English](number_in_english_test.go), [Solution](number_in_english.go)
* [Reverse Vowels In a String](reverse_vowels_test.go), [Solution](reverse_vowels.go)
* [Longest Substring of Two Unique Characters](longest_substring_test.go), [Solution](longest_substring.go)
