// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// )

// func generatePassword(length int) string {
// 	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
// 	password := make([]byte, length)
// 	for i := range password {
// 		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
// 		password[i] = chars[randomIndex.Int64()]
// 	}
// 	return string(password)
// }

// func main() {
// 	fmt.Println(generatePassword(12))
// }

// The code is mostly correct, but it lacks proper error handling and the big package import, which is necessary for big.NewInt. Here is an improved version with error handling and the required import.

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generatePassword(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	password := make([]byte, length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return ""
		}
		password[i] = chars[randomIndex.Int64()]
	}
	return string(password)
}

func main() {
	fmt.Println(generatePassword(12))
}

// Improvements:
// Error Handling: Added error handling for the rand.Int function to ensure any issues are reported.
// Import math/big: Added the necessary import for the big package.
// This version ensures that any errors during the random number generation are handled and reported.
