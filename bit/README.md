# Bit Manipulation

Bit manipulation or bitwise operations refers to manipulating individual or groups of bits in a computer's memory using logical operations to solve a problem. CPUs supporting specialized bitwise instructions are designed to execute bit-level operations efficiently, including AND, OR, XOR, SHIFT, and other similar operations. As these instructions are implemented in hardware, they offer excellent performance and are widely used in high-performance computing applications.

```ASCII
[Figure 1] Truth tables and illustration of bit manipulation operations

AND	1100	OR	1100	XOR	1100	Negation	1100	L Shift	1100	R Shift	1100
	1010		1010		1010
	====		====		====			====		====		====
	1000		1110		0110			0011		1000		0110
```

## Implementation

Go provides the below operators for bit manipulation:

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

	printBinary(a & c)    // Prints "10" (Binary AND operation)
	printBinary(a | b)    // Prints "111111" (Binary OR operation)
	printBinary(a ^ b)    // Prints "111111" (Binary XOR operation)
	printBinary(^c)       // Prints "-10100" (Binary Negation operation)
	printBinary(a << 2)   // Prints "10101000" (Binary Left Shift operation)
	printBinary(a >> 2)   // Prints "1010" (Binary Right Shift operation)
	printBinary(a ^ (^a)) // Prints "-1" (XOR with its own negation)
}

// printBinary prints an int as its binary string using strconv.FormatInt.
func printBinary(n int) {
	fmt.Println(strconv.FormatInt(int64(n), 2))
}
```

## Arithmetic by Shifting

Left shifting can be viewed as a multiplication operation by 2 raised to the power of a specified number. Right shifting can be viewed as a division operation by 2 raised to the power of a specified number. For instance, a << b can be interpreted as multiplying a by 2^b, and a >> b can be interpreted as dividing a by 2^b.

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

The XOR operation can be used for basic cryptography. A message can be encrypted by being XORed with a key. This encrypted message can be shared with someone else who knows the same key. If they XOR the key with the encrypted message, they will obtain the original plaintext message. This method is not secure enough because the key is relatively easy to guess from the encrypted message. The following example demonstrates this process:

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

Bit manipulation operations are characterized by a constant time complexity. This high level of performance renders them an optimal choice to replace other approaches, especially when working with large data sets. As a result, they are frequently used to achieve better performance.

## Application

Bit manipulation techniques are widely utilized in diverse fields of computing, such as cryptography, data compression, network protocols, and databases, to name a few. Each bitwise operation has its own qualities that make it useful in different scenarios.

AND extracts bit(s) from a larger number. For example, to check if a certain bit is set in a number, the number can be ANDed with a mask that has only that bit set to 1, and if the result is not 0, then that bit was set. Another application is to clear or reset certain bits in a number by ANDing with a mask with those bits set to 0.

OR can be useful to set or turn on certain bits in a binary number. For example with a variable flag, a binary number representing various options, a particular flag can be set by ORing the variable with a binary number where only the corresponding bit for that flag is 1. This will turn on the flag in the variable without affecting other flags.

XOR can be used for encryption, decryption, error detection, and correction. It can also be used to swap two variables without using a temporary variable. Additionally, XOR can be used to solve problems related to finding unique elements in a list or array or to check whether two data sets have overlapping elements.

Negation can be used to invert a set of flags or find the two's complement of a number.

## Rehearsal

* [Division without multiplication or division operators](division_without_operators_test.go), [Solution](division_without_operators.go)
* [Middle without division](middle_without_division_test.go), [Solution](middle_without_division.go)
* [Addition without using plus (+) or any other arithmetic operators](addition_without_operators_test.go), [Solution](addition_without_operators.go)
* [Power of Two](./is_power_of_two_test.go), [Solution](./is_power_of_two.go)
* [Maximum without if conditions](max_function_without_conditions_test.go), [Solution](max_function_without_conditions.go)
* [Oddly Repeated Number](oddly_repeated_number_test.go), [Solution](oddly_repeated_number.go)
