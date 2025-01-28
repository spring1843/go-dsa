# String

A string is a ubiquitous data structure, typically a built-in data type in programming languages. However, beneath the surface, strings are essentially [slices](../array/) of characters that enable textual data storage and manipulation.

## Implementation

In Go, [strings](https://go.dev/blog/strings) are a data type. Behind the scenes strings are an immutable slice of bytes. Since Go is a UTF-8 compliant language, each character in Go can take up to 4 bytes of storage.

The `strings` package provides several useful convenience functions. Examples include:

* [Index](https://golang.org/pkg/strings/#Index), [Contains](https://golang.org/pkg/strings/#Contains), [HasPrefix](https://golang.org/pkg/strings/#HasPrefix), [HasSuffix](https://golang.org/pkg/strings/#HasSuffix)
* [Split](https://golang.org/pkg/strings/#Split), [Fields](https://golang.org/pkg/strings/#Split), [Join](https://golang.org/pkg/strings/#Join)
* [Repeat](https://golang.org/pkg/strings/#Repeat)
* [ReplaceAll](https://golang.org/pkg/strings/#ReplaceAll)
* [Title](https://golang.org/pkg/strings/#Title), [ToLower](https://golang.org/pkg/strings/#ToLower), [ToUpper](https://golang.org/pkg/strings/#ToUpper)
* [Trim](https://golang.org/pkg/strings/#Trim), [TrimSpace](https://golang.org/pkg/strings/#TrimSpace), [TrimSuffix](https://golang.org/pkg/strings/#TrimSuffix), [TrimPrefix](https://golang.org/pkg/strings/#TrimPrefix)

When a iterating every character in a string in Go using the `range` keyword, every element becomes a [rune](https://blog.golang.org/strings#TOC_5.) which is an alias for the type `int32`. If the code being written works with many single-character strings, it is better to define variables and function parameters as `rune` rather than convert them many times. The following code shows how to iterate through a string.

```Go
package main

import "fmt"

/*
 main outputs the rune (int32) value of each character:

 Char #0 "a" has value 97
 Char #1 "A" has value 65
 Char #2 "和" has value 21644
 Char #5 "平" has value 24179
 Char #8 "😊" has value 128522
*/
func main() {
	for i, r := range "aA𓅚😊" {
		fmt.Printf("Char #%d %q has value %d\n", i, string(r), r)
	}
}
```

A very common tool to use for manipulating strings in Go is the `fmt.Sprtinf` function. This is specially useful when converting many values into a string.

```Go
package main

import "fmt"

func main() {
	number := 1
	value := 1.1
	name := "foo"

	output := fmt.Sprintf("%d %f %s", number, value, name)
	fmt.Println(output) // 1 1.100000 foo
}
```

### Regular Expressions

Unlike many other programming languages, in Go [regular expressions](https://golang.org/pkg/regexp/) are [guaranteed](https://swtch.com/~rsc/regexp/regexp1.html) to have O(n) time complexity where n is the length of the input, making them a viable and practical option for pattern matching in a string.

Here is an example of how you can find a pattern using regular expressions in Go. Given a string return the string if it contains a fish word. A fish word is a word that starts with `fi` optionally followed by other character(s) and ends with `sh`. Examples include {`fish`, `finish`}.

```GO
package main

import (
	"fmt"
	"regexp"
)

var fishPattern = regexp.MustCompile(`(?i).*fi\w*sh\b`)

// main outputs [selfish][shellfish][fish][finish][Finnish]
func main() {
	inputs := []string{"shift", "selfish", "shellfish", "fish dish", "finish", "Finnish"}

	for _, input := range inputs {
		matches := fishPattern.FindAllString(input, -1)
		if len(matches) > 0 {
			fmt.Print(matches)
		}
	}
}
```

## Complexity

Since strings are slices of bytes, the time complexity of string operations should be similar to [arrays](../array/). Reading a character at a given index is O(1), but since strings are immutable modifying them involves creating a new string making it a O(n) operation. Go standard library includes `strings.Builder` for more efficient string building.

The space complexity to store a string depends on the type of characters. This following example shows how we can index a string and print the hexadecimal value of every byte in it.

```GO
package main

import "fmt"

// main Outputs 41 f0 93 85 9a f0 9f 98 8a.
func main() {
	input := "A𓅚😊"
	for i := 0; i < len(input); i++ {
		fmt.Printf("%x ", input[i])
	}
}
```

The output of the above code indicates that 9 bytes are used to store the 3 input characters. 1 byte for the first character and 4 bytes for each of the remaining two.

## Application

Strings store words, characters, sentences, etc.

## Rehearsal

* [The Longest Dictionary Word Containing Key](./longest_dictionary_word_test.go), [Solution](./longest_dictionary_word.go)
* [Look and Tell](./look_and_tell_test.go), [Solution](./look_and_tell.go)
* [In Memory Database](./in_memory_database_test.go), [Solution](./in_memory_database.go)
* [Number in English](./number_in_english_test.go), [Solution](./number_in_english.go)
* [Reverse Vowels in a String](./reverse_vowels_test.go), [Solution](./reverse_vowels.go)
* [Roman Numerals](./roman_numerals_test.go), [Solution](./roman_numerals.go)
* [Longest Substring of Two Unique Characters](./longest_substring_test.go), [Solution](./longest_substring.go)
