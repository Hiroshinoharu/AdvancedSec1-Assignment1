// Program: caesarCipher.go
// Description: This program encrypts the string using caesar's cipher
// Date: 23/09/2025

package main

import "fmt"
import "unicode"

// caesar_cipher applies a Caesar cipher to the input text, shifting each letter by the specified amount.
// Uppercase and lowercase letters are shifted independently, wrapping around the alphabet if necessary.
// Non-alphabetic characters are not handled and may produce unexpected results.
// Parameters:
//   text  - the input string to be ciphered
//   shift - the number of positions to shift each letter
// Returns:
//   The ciphered string with each letter shifted by the specified amount.
func caesar_cipher(text string, shift int) string{
	var result string = " ";

	// Traverse the text
	for _,ch := range text {
		if (unicode.IsUpper(ch)) {
			// Upper Case characters
			// This is shifted by taking the Unicode and shifting by the specified shift value
			//Its moded by 26 to show the 26 alphabets
			// 'A' in unicode is 65 
			result += string('A' + ((ch -'A'+ rune(shift)) % 26))
		} else{
			// Lowercase characters are shifted
			// 'a' = 97 for context
			result += string('a' + ((ch -'a' + rune(shift)) % 26))
		}
	}

	// Returns the ciphered text
	return result
}

// decrypt_cipher_text decrypts a given Caesar ciphered text using the specified shift value.
// It shifts each alphabetical character backwards by the shift amount, preserving case.
// Non-alphabetical characters are left unchanged.
// Parameters:
//   ciphered_text - the encrypted string to be decrypted.
//   shift - the integer shift value used for decryption.
// Returns:
//   The decrypted string.
func decrypt_cipher_text(ciphered_text string, shift int) string{
	var result string = "";

	for _, ch := range ciphered_text {
		if unicode.IsUpper(ch) {
			// Decyrpting is just take away from the shift key
			result += string('A' + ((ch -'A'-rune(shift)+26)%26))
		} else if unicode.IsLower(ch) {
			result += string('a' + ((ch -'a'-rune(shift)+26)%26))
		} else{
			// Keep other characters unchanged
			result += string(ch)
		}
	}

	return result
}

func main() {
	var plaintext string = "Hello"
	var shift_value = 3
	var ciphered_text = caesar_cipher(plaintext, shift_value)

	var encrypted_text = "RQH YDULDWLRQ WR WKH VWDQGDUG FDHVDU FLSKHU LV ZKHQ WKH DOSKDEHW LV \"NHBHG\" EB XVLQJ D ZRUG. LQ WKH WUDGLWLRQDO YDULHWB, RQH FRXOG ZULWH WKH DOSKDEHW RQ WZR VWULSV DQG MXVW PDWFK XS WKH VWULSV DIWHU VOLGLQJ WKH ERWWRP VWULS WR WKH OHIW RU ULJKW. WR HQFRGH, BRX ZRXOG ILQG D OHWWHU LQ WKH WRS URZ DQG VXEVWLWXWH LW IRU WKH OHWWHU LQ WKH ERWWRP URZ. IRU D NHBHG YHUVLRQ, RQH ZRXOG QRW XVH D VWDQGDUG DOSKDEHW, EXW ZRXOG ILUVW ZULWH D ZRUG (RPLWWLQJ GXSOLFDWHG OHWWHUV) DQG WKHQ ZULWH WKH UHPDLQLQJ OHWWHUV RI WKH DOSKDEHW. IRU WKH HADPSOH EHORZ, L XVHG D NHB RI \"UXPNLQ.FRP\" DQG BRX ZLOO VHH WKDW WKH SHULRG LV UHPRYHG EHFDXVH LW LV QRW D OHWWHU. BRX ZLOO DOVR QRWLFH WKH VHFRQG \"P\" LV QRW LQFOXGHG EHFDXVH WKHUH ZDV DQ P DOUHDGB DQG BRX FDQ'W KDYH GXSOLFDWHV"

	var decyrpted_text = decrypt_cipher_text(encrypted_text, 3)

	fmt.Print("Plantext: ", plaintext)
	fmt.Print("\nShift Key: ", shift_value)
	fmt.Print("\nCiphered Text:", ciphered_text, "\n")
	fmt.Print("\nDecrypted Text: ", decyrpted_text, "\n")
}