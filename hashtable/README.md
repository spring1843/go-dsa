# Hash Table

Hash tables are a fundamental data structure that operates based on key-value pairs and enables constant-time operations for lookup, insertion, and deletion. The keys used in hash tables are immutable and can be a simple string or integer in basic usage. However, in more complex applications, a hashing function along with different collision resolution methods such as separate chaining, linear probing, quadratic probing, and double hashing can be used to ensure efficient performance.

## Implementation

In Go, hash tables are implemented as maps, which is a built-in data type of the language. To declare a map, the data type for the key and the value must be specified, and the map needs to be initialized using the make function before it can be used. Below is an example of how to declare a map with string keys and integer values:

```Go
package main

import (
	"fmt"
)

func main() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4

	for i, v := range numbers {
		fmt.Printf("%s:%d", i, v)
	}
}
```

When using maps in Go, it is crucial to remember that the order of the items stored in the map is not preserved, unlike arrays and slices. Relying on the order of the contents of a map can lead to unexpected issues, such as code behaving inconsistently, and intermittent failures.

## Complexity

Hash tables provide O(1) time complexity for inserting, deletion and searching operations.

## Application

Hash tables are widely used in algorithms to cache and memoize data for fast constant access times. This advantage in performance makes hash tables more suitable than [Arrays](../arrays) and even [Binary Search Trees](../tree), if there is no need to keep the data ordered, they can't however be used for finding a range, min or max of data.

Compilers use hash tables to generate a symbol table, to keep track of variable declarations.

## Rehearsal

### Find Missing Number

Given an unsorted array of numbers like {7,5,3,4,1,2,0,-1} return the missing integer like 6. [Solution](missing_number.go) [Test](missing_number.go)

### List Elements Summing Up to K

Given a list, output the indices of the first two elements that sum up to K. [Solution](sum_up_to_k.go) [Test](sum_up_to_k_test.go)

### Fastest Way to Cut a Brick Wall

Given a two dimensional array of numbers representing the length of each brick in each row of a wall, return the position at which we could do the least amount of cuts (by not having to cut the touch points where bricks meet). [Solution](cut_brick_wall.go) [Test](cut_brick_wall_test.go)

### Find Anagrams

Given a dictionary, return lists of words that are anagrams of each other. [Solution](find_anagrams.go), [Tests](find_anagrams_test.go)

### Find Max Points on the Same Line

Given multiple coordinates of points like {[1,1], [2,2], [3,3], [4,5]}, return the maximum number of points that are on the same line like 3. [Solution](max_lines_on_line.go) [Test](max_lines_on_line_test.go)
