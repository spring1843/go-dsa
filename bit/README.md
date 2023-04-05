# Bit Manipulation

Bit manipulation or bitwise operations refers to the process of manipulating individual or groups of bits in a computer's memory using logical operations to solve a problem. CPUs with support for specialized bitwise instructions are designed to execute bit-level operations efficiently, including AND, OR, XOR, SHIFT, and other similar operations. As these instructions are implemented in hardware, they offer excellent performance and are widely used in high-performance computing applications.


```ASCII
[Figure 1] Truth tables and illustration of bit manipulation operations

AND	1100	OR	1100	XOR	1100	Negation	1100	L Shift	1100	R Shift	1100
	1010		1010		1010
	====		====		====			====		====		====
	1000		1110		0110			0011		1000		0110
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

	printBinary(a & c)    // 10           And
	printBinary(a | b)    // 111111       Or
	printBinary(a ^ b)    // 111111       Xor
	printBinary(^c)       // -10100       Negation
	printBinary(a << 2)   // 10101000     Left shift
	printBinary(a >> 2)   // 1010         Right shift
	printBinary(a ^ (^a)) // -1	    Xor with negate of self
}

func printBinary(n int) {
	fmt.Println(strconv.FormatInt(int64(n), 2))
}
```

## Arithmetic by Shifting

Left shifting can be viewed as a multiplication operation by 2 raised to the power of a specified number, while right shifting can be viewed as a division operation by 2 raised to the power of a specified number. For instance, a << b can be interpreted as multiplying a by 2^b, and a >> b can be interpreted as dividing a by 2^b.

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

## Cryptography and Other XOR applications

The XOR operation can be used to create perform basic cryptography. By XORing a message with a key, we can generate an encrypted message. This encrypted message can be shared with someone else who knows the same key. If they XOR the key with the encrypted message, they will obtain the original plaintext message. This method is not secure enough because the key is relatively easy to guess from the encrypted message. The following example demonstrates this process:

```Go
package main

import "fmt"

func main() {
	key := []byte("secret key")
	message := []byte("hello world")

	encryptedMsg := xorCrypt(key, message)
	fmt.Printf("Encrypted message: %v\n", encryptedMsg) // Prints [27 0 15 30 10 84 87 4 23 21 23]
	fmt.Printf("Decrypted message: %s", xorCrypt(key, encryptedMsg)) // Prints hello world
}

func xorCrypt(key, message []byte) []byte {
	for i, char := range message {
		message[i] = key[i%len(key)] ^ char
	}
	return message
}
```

## Complexity

Bit manipulation operations are characterized by a constant time complexity of O(1). This high level of performance renders them an optimal choice as a replacement for other approaches, especially when working with large data sets. As a result, they are frequently utilized in algorithmic design to optimize the execution of certain operations.

## Application

Bit manipulation techniques are widely utilized in diverse fields of computing, such as cryptography, data compression, network protocols, and databases, to name a few. Each specific bitwise operation has it's own qualities that makes them useful in different scenarios.

AND is used to extract specific bit or set of bits from a larger number. For example, to check if a certain bit is set in a number, we can AND the number with a mask that has only that bit set to 1, and if the result is not 0, then that bit was set. Another application is to clear or reset certain bits in a number by ANDing with a mask that has those bits set to 0.

OR can be useful in solving problems where we want to "set" or "turn on" certain bits in a binary number. For example, if we have a variable flags which is a binary number representing various options, we can set a particular flag by ORing the variable with a binary number where only the corresponding bit for that flag is 1. This will turn on the flag in the variable without affecting any other flags.

XOR can be used for encryption and decryption, as well as error detection and correction. It can also be used to swap two variables without using a temporary variable. Additionally, XOR can be used to solve problems related to finding unique elements in a list or array, or to check whether two sets of data have any overlapping elements

Negation can be used to invert a set of flags or to find the two's complement of a number. In computer architecture, negation is often used in the implementation of logical and arithmetic operations.

## Rehearsal

### Division without multiplication or division operators

Divide x by y, two integers without using the built-in `/` or `*` operators.  [Solution](division_without_operators.go), [Tests](division_without_operators_test.go)

### Middle without division

Given two integers min and max like `1` and `5`, return an integer like `3` that is in the middle of the two. [Solution](middle_without_division.go), [Tests](middle_without_division_test.go)

### Addition without using plus (+) or any other arithmetic operators

Add x by y, two integers without using the built-in + or any other arithmetic operators. [Solution](addition_without_operators.go), [Tests](addition_without_operators_test.go)

### Maximum without if conditions

Write max, a function that returns the largest of two numbers without using a comparison or if conditions. [Solution](max_function_without_conditions.go), [Tests](max_function_without_conditions_test.go)

### Oddly Repeated Number

Given an array of integers that are all repeated an even number of times except one, find the oddly repeated element. [Solution](oddly_repeated_number.go), [Tests](oddly_repeated_number_test.go)
