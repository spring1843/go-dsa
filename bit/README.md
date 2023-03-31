# Bit Manipulation

Bit manipulation or bitwise operations refers to the process of manipulating individual or groups of bits in a computer's memory using logical operations to solve a problem. CPUs with support for specialized bitwise instructions are designed to execute bit-level operations efficiently, including AND, OR, XOR, SHIFT, and other similar operations. As these instructions are implemented in hardware, they offer excellent performance and are widely used in high-performance computing applications.

## Arithmetic by Shifting

Left shifting can be seen as dividing by 2 or 2 raised to the power of a given number, and right shifting can be seen multiplying by 2 raised to the power of a given number. For example `a << b` can be seen as dividing a by 2^b, and `a >> b` can be seen as multiplying a by 2^b times.

```Go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 5)  // Prints 32. 1 * 2^5 = 32
	fmt.Println(2 << 1)  // Prints 4. 2 * 2^1 = 4
	fmt.Println(2 << 2)  // Prints 8. 2 * 2^2 = 8
	fmt.Println(8 >> 1)  // Prints 4. 8 / 2^1 = 4
	fmt.Println(32 >> 5) // Prints 1. 32 / 2^5 = 1
	fmt.Println(8 >> 2)  // Prints 2. 8 / 2^2 = 2
}
```


# Implementation

Go provides below operators that can be used in bit manipulation:

```Go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 0b101010
	b := 0b010101
	c := 0b010011

	printBinary(^c)       // -10100       Negation
	printBinary(a & c)    // 10           And
	printBinary(a | b)    // 111111       Or
	printBinary(a ^ b)    // 111111       Xor
	printBinary(a ^ c)    // 111001       Xor
	printBinary(a << 2)   // 10101000     Left shift
	printBinary(a >> 2)   // 1010         Right shift
	printBinary(a ^ (^a)) // -1	    Xor with negate of self
}

func printBinary(n int) {
	fmt.Println(strconv.FormatInt(int64(n), 2))
}
```

## Complexity

Bit manipulation operations are characterized by a constant time complexity of O(1). This high level of performance renders them an optimal choice as a replacement for other approaches, especially when working with large data sets. As a result, they are frequently utilized in algorithmic design to optimize the execution of certain operations.

## Application

Bit manipulation techniques are widely utilized in diverse fields of computing, such as cryptography, data compression, network protocols, and databases, to name a few.

## Rehearsal

### Division without multiplication or division operators

Divide x by y, two integers without using the built-in `/` or `*` operators.  [Solution](division_without_operators.go), [Tests](division_without_operators_test.go)

### Middle without division

Given two integers min and max like `1` and `5`, return an integer like `3` that is in the middle of the two. [Solution](middle_without_division.go), [Tests](middle_without_division_test.go)

### Addition without using any operators

Add x by y, two integers without using the built-in + or any other operators. [Solution](addition_without_operators.go), [Tests](addition_without_operators_test.go)

### Maximum without if conditions

Write max, a function that returns the largest of two numbers without using a comparison or if conditions. [Solution](max_function_without_conditions.go), [Tests](max_function_without_conditions_test.go)

### Oddly Repeated Number

Given an array of integers that are all repeated an even number of times except one, find the oddly repeated element. [Solution](oddly_repeated_number.go), [Tests](oddly_repeated_number_test.go)
