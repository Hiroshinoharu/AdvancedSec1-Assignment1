package main

import "fmt"
import "unicode"


// Declares the matrixes nesscary 

// Matrix needed to encrypt/decrypt the text
var keyMatrix = [][]int{
	{0, 0},
	{0, 0},
}

// This the matrix needed to put our plain/cipher text in
var textMatrix = [][]int{
	{0},
	{0},
}

// This function gets the key matrix
func getKeyMatrix(key string) [][]int{
	k := 0
	for i := range keyMatrix{
		for j := range keyMatrix[i] {
			keyMatrix[i][j] = int(unicode.ToUpper(rune(key[k]))- 'A')
			k++
		}
	}
	return keyMatrix
}

// Encrypt or decrypt a 2x1 block using a given key matrix
func multiplyMatrix(key [][]int, text [][]int) [][]int {
	result := [][]int{
		{0},
		{0},
	}
	// Does matrix multiplication (ROW X COLUMN)
	for i := 0; i < 2; i++ {
		for j := 0; j < 1; j++ {
			sum := 0
			for k := 0; k < 2; k++ {
				sum += key[i][k] * text[k][j]
			}
			result[i][j] = (sum % 26 + 26) % 26 // ensure positive modulo
		}
	}
	return result
}

// Find modular inverse of a number modulo m
// modInverse computes the modular multiplicative inverse of 'a' under modulo 'm'.
// It returns an integer 'x' such that (a * x) % m == 1, if such an integer exists.
// If the modular inverse does not exist, the function returns -1.
// This function uses a brute-force approach and is suitable for small values of 'm'.
func modInverse(a, m int) int{
	a = a % m
	for x := 1; x < m; x++{
		if (a*x) % m == 1{
			return x
		}
	}
	return -1
}

// inverseKeyMatrix computes the modular inverse of a 2×2 key matrix
// used in a Hill cipher. The result is the matrix K⁻¹ such that
// (K * K⁻¹) ≡ I (mod 26). It panics if the key matrix is not invertible.
func inverseKeyMatrix(key [][]int) [][]int {
	// Extract the 2x2 matrix elements:
	//   | a  b |
	//   | c  d |
	a, b := key[0][0], key[0][1]
	c, d := key[1][0], key[1][1]

	// Compute the determinant: det = a*d - b*c
	// This value must be coprime with 26 for the matrix to be invertible mod 26.
	determinant := (a*d - b*c) % 26

	// Ensure the determinant is positive (mod 26 should be in [0,25])
	if determinant < 0 {
		determinant += 26
	}

	// Find the modular multiplicative inverse of the determinant
	// such that (determinant * determinant_inverse) ≡ 1 (mod 26).
	determinant_inverse := modInverse(determinant, 26)
	if determinant_inverse == -1 {
		// If there is no modular inverse, the key matrix cannot be used to decrypt.
		panic("Key Matrix is not invertible modulo 26!")
	}

	// Apply the formula for the inverse of a 2×2 matrix:
	// K⁻¹ = (det⁻¹) * [ d  -b
	//                   -c  a ]  (all calculations mod 26)
	inverseKeyMatrix := [][]int{
		{d * determinant_inverse % 26, (-b) * determinant_inverse % 26},
		{(-c) * determinant_inverse % 26, a * determinant_inverse % 26},
	}

	// Convert any negative numbers into their positive modular equivalents
	// so every element is between 0 and 25.
	for i := range inverseKeyMatrix {
		for j := range inverseKeyMatrix[i] {
			inverseKeyMatrix[i][j] = (inverseKeyMatrix[i][j] + 26) % 26
		}
	}

	// Return the 2×2 inverse key matrix mod 26.
	return inverseKeyMatrix
}

// Hill Cipher Encryption
func EncryptUsingHillCipher(plain_text string, key string) string{
	result := ""

	// Get the key matrix from the key string
	getKeyMatrix(key)

	fmt.Print("Printing Key Matrix for debugging: ")
	// Print the the matrixes for debugging purposes
	for i := range keyMatrix{
		for j := range keyMatrix[i] {
			fmt.Print(keyMatrix[i][j], " ")
		}
	}
	fmt.Println()

	// Convert the plain_text into a slice of runes
	plainRunes := []rune(plain_text)

	// Pad the plaintext if its length is not a multiple of 2
	if len(plainRunes)%2 != 0 {
		plainRunes = append(plainRunes, 'X')
	}

	// Process the plaintext in blocks of 2
	for b := 0; b < len(plainRunes); b += 2 {
		// Fill the textMatrix for this block
		for i := 0; i < 2; i++ {
			textMatrix[i][0] = int(unicode.ToUpper(plainRunes[b+i])) - 65
		}
		// Encrypt the block
		cipher := multiplyMatrix(keyMatrix,textMatrix)
		// Append the cipher text for this block
		for i := 0; i < 2; i++ {
			result += string(rune(cipher[i][0] + 65))
		}
	}

	return result
}

func decryptHillCipher(cipher_text string, key string) string{
	result := " "
	keyMatrix := getKeyMatrix(key)
	inverseKey := inverseKeyMatrix(keyMatrix)

	// Gets the the string and converts into a slice of runes
	cipheRunes := []rune(cipher_text)

	// Process the ciphertext in blocks of 2
	for b:=0; b < len(cipheRunes); b += 2 {
		// Fill the text matrix for this block
		for i := 0; i < 2; i++{
			textMatrix[i][0] = int(unicode.ToUpper(cipheRunes[b+i])) - 65
		}
		// Gets the plain text matrix to decrypt 
		plain := multiplyMatrix(inverseKey, textMatrix)
		// Fill the result string
		for i := 0; i < 2; i++{
			result += string(rune(plain[i][0] + 65))
		}
	}
	return  result
}

func main(){
	plain_text := "dogs"
	key := "hill"

	fmt.Printf("Encrypting %s with key %s\n", plain_text, key)

	encrypted_text := EncryptUsingHillCipher(plain_text, key)

	fmt.Printf("\nCipher Text: %s\n", encrypted_text)

	// Decrypt the text nows
	encrypted_text = "owrr"

	fmt.Printf("\nDecrypting %s with key %s\n", encrypted_text, key)

	new_plain_text := decryptHillCipher(encrypted_text, key)

	fmt.Printf("\nPlaintext: %s\n", new_plain_text)
}