// Program Name: vigenere_cipher.go
// Description: This program implments vigenere's cipher in GO
// Author: Max Ceban
// Date: 28/09/2025

package main

import "fmt"
import "unicode"

func generateKey(plain_text string, key string) string {
	
	// Declares a slice of runes to help match the key length to the plain_text
	keyRunes := []rune(key) // This rune represents the key
	plainRunes := []rune(plain_text) // This slice represents plaintext
	
	// If the length of the slice of the key matches the plaintext it'll return the rune converted to a string
	if len(keyRunes) == len(plainRunes) {
		return string(keyRunes)
	}
	// Else it'll keep adding the characters until it matches the size of the string
	for i := 0; len(keyRunes) < len(plainRunes); i++ {
		keyRunes = append(keyRunes, keyRunes[i % len(keyRunes)])
	}
	// Returns the value
	return string(keyRunes)
}

func vigernere_encrypt(plain_text string, key string) string{
	result := ""
	key = generateKey(plain_text, key)

	// Convert the plaintext and key into a slice of runes
	keyRunes := []rune(key)
	plainRunes := []rune(plain_text)

	// Iterating through the string block
	for i, ch := range plainRunes{
		if unicode.IsUpper(ch){
			// Add the shift from the key to the encrypted text
			result += string(((ch - 'A') + (keyRunes[i] - 'A')) % 26 + 'A')
		} else if unicode.IsLower(ch){
			result += string(((ch - 'a') + (keyRunes[i] - 'a')) % 26 + 'a')
		} else{
			result += string(ch)
		}
	}

	return result
}

func main(){
	plain_text := "explanation"
	key := "leg"

	fmt.Printf("Plaintext: %s\n", plain_text);
	fmt.Printf("Encrypting: %s using the key %s\n", plain_text, key)

	fmt.Printf("Key for testing purposes: %s\n", generateKey(plain_text,key))

	encrypted_text := vigernere_encrypt(plain_text, key)
	fmt.Printf("Encrypted text: %s\n", encrypted_text)
}