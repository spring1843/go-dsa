# Hash Table

Hash tables are a fundamental data structure that operates based on key-value pairs and enables constant-time operations for lookup, insertion, and deletion. Hash tables use immutable keys that can be strings or integers among other things. However, in more complex applications, a hashing function, and different collision resolution methods such as separate chaining, linear probing, quadratic probing, and double hashing, can be used.

## Implementation

In Go, hash tables are implemented as maps, a built-in language data type. To declare a map, the data type for the key and the value must be specified. The map needs to be initialized using the make function before it can be used. Below is an example of how to declare a map with string keys and integer values:

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

When using maps in Go, it is crucial to remember that the order of the items stored on the map is not preserved. This is unlike arrays and slices. Relying on the order of map contents can lead to unexpected issues, such as inconsistent code behavior and intermittent failures.

As shown below it is possible in Go to store variables as keys in a map. It is also possible to have a map of only keys with no values.

```Go
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

var (
	person1           = &person{name: "p1", age: 20}
	person2           = &person{name: "p2", age: 30}
	person3           = &person{name: "p3", age: 40}
	allowListedPeople = map[*person]struct{}{
		person1: struct{}{},
		person2: struct{}{},
	}
)

func main() {
	isAllowed(person1) // p1 is allowed
	isAllowed(person2) // p2 is allowed
	isAllowed(person3) // p3 is not allowed
}

func isAllowed(person *person) {
	verb := "is"
	if _, ok := allowListedPeople[person]; !ok {
		verb += " not"
	}
	fmt.Printf("%s %s allowed\n", person.name, verb)
}
```

## Complexity

Hash tables provide O(1) time complexity for inserting, deletion, and searching operations.

## Application

When there is no need to preserve the order of data, hash tables are used for fast O(1) reads and writes. This performance advantage makes hash tables more suitable than [Arrays](../array) and [Binary Search Trees](../tree).

Compilers use hash tables to generate a symbol table to keep track of variable declarations.

## Rehearsal

* [Find Missing Number](./missing_number_test.go), [Solution](./missing_number.go)
* [List Elements Summing Up to K](./sum_up_to_k_test.go), [Solution](./sum_up_to_k.go)
* [Fastest Way to Cut a Brick Wall](./cut_brick_wall_test.go), [Solution](./cut_brick_wall.go)
* [Smallest Missing Positive Integer](./smallest_missing_positive_integer_test.go), [Solution](./smallest_missing_positive_integer.go)
* [Find Anagrams](./find_anagrams_test.go), [Solution](./find_anagrams.go)
* [Find Max Points on the Same Line](./max_points_on_line_test.go), [Solution](./max_points_on_line.go)
