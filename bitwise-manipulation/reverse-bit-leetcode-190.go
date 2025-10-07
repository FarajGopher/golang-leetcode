package main

import "fmt"

func reverseBits(num uint32) uint32 {
    var result uint32 = 0
	// Iterate through all 32 bits of the integer
    for i := 0; i < 32; i++ {
        // Shift result left to make space for the next bit
        result <<= 1
        // Add the least significant bit of num to result
        result |= num & 1
		fmt.Printf("num: %032b result: %032b\n", num, result)
        // Shift num right to process the next bit
        num >>= 1
    }

    return result
}

func main() {
    var n uint32 = 43261596
    fmt.Println(reverseBits(n)) // Output: 964176192
}
