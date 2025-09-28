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
	// Does matrix multiplication 
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
func modInverse(a, m int) int{
	a = a % m
	for x := 1; x < m; x++{
		if (a*x) % m == 1{
			return x
		}
	}
	return -1
}

// Compute inverse key matrix modulo 26
func inverseKeyMatrix(key [][]int) [][]int{
	// Get the key values for the formula
	a, b := key[0][0], key[0][1]
	c, d := key[1][0], key[1][1]

	// Calculate the determinant
	determinant := (a*d - b*c) % 26
	// Ensures the deteminant is positive
	if determinant < 0{
		determinant += 26
	}

	// Gets the inverse of the determinant 
	determinant_inverse := modInverse(determinant, 26)
	if determinant_inverse == -1{
		panic("Key Matrix is not invertiable modulo 26!!!!!")
	}

	// Gets the inverse key matrix
	inverseKeyMatrix := [][]int{
		{d * determinant_inverse % 26, (-b) * determinant_inverse % 26},
		{(-c) * determinant_inverse % 26, a * determinant_inverse % 26},
	}

	// Make all elements positive modulo 26
	for i := range inverseKeyMatrix{
		for j := range inverseKeyMatrix[i]{
			inverseKeyMatrix[i][j] = (inverseKeyMatrix[i][j] + 26) % 26
		}
	}

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