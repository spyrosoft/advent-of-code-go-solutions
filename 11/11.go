package main

import (
	"fmt"
	"os"
)

var (
	validPasswords = []string{"abccddzv", "bcddcfzz", "abefggkk"}
	invalidPasswords = []string{"hijklmmn", "abbceffg", "abbcegjk"}
	runeIncrementLookup = map[rune]rune{'a': 'b', 'b': 'c', 'c': 'd', 'd': 'e', 'e': 'f', 'f': 'g', 'g': 'h', 'h': 'j', 'j': 'k', 'k': 'm', 'm': 'n', 'n': 'p', 'p': 'q', 'q': 'r', 'r': 's', 's': 't', 't': 'u', 'u': 'v', 'v': 'w', 'w': 'x', 'x': 'y', 'y': 'z', 'z': 'a'}
)

func checkForStraight( first int, second int, third int ) bool {
	if first + 1 == second && second + 1 == third {
		return true
	}
	return false
}

func checkForDoubleLetter( first int, second int ) bool {
	if first == second {
		return true
	}
	return false
}

func validatePassword( password string ) bool {
	if len( password ) != 8 {
		return false
	}
	straightExists := false
	allValidCharacters := true
	numberOfDoubleLetters := 0
	previousDoubleLetterIndex := 0
	for i, rune := range password {
		if i >= 2 {
			if checkForStraight( int( password[ i - 2 ] ), int( password[ i - 1 ] ), int( password[ i ] ) ) {
				straightExists = true
			}
		}
		if i >= 1 {
			if numberOfDoubleLetters == 0 || previousDoubleLetterIndex < i - 1 {
				if checkForDoubleLetter( int( password[ i - 1 ] ), int( password[ i ]  ) ) {
					numberOfDoubleLetters++
					previousDoubleLetterIndex = i
				}
			}
		}
		if rune == 'i' || rune == 'o' || rune == 'l' {
			allValidCharacters = false
		}
	}
	if straightExists && allValidCharacters && numberOfDoubleLetters >= 2 {
		return true
	}
	return false
}

func testValidPasswords() {
	for _, password := range validPasswords {
		if !validatePassword( password ) {
			fmt.Println( "Incorrectly invalid password: ", password )
			os.Exit( 1 )
		}
	}
}

func testInvalidPasswords() {
	for _, password := range invalidPasswords {
		if validatePassword( password ) {
			fmt.Println( "Incorrectly valid password: ", password )
			os.Exit( 1 )
		}
	}
}

func incrementPassword( password string ) string {
	passwordRunes := []rune( password )
	for i := len( password ) - 1; i >= 0; i-- {
		passwordRunes[ i ] = runeIncrementLookup[ rune( password[ i ] ) ]
		if rune( password[ i ] ) != 'a' {
			break
		}
	}
	password = string( passwordRunes )
	return password
}

func nextValidPassword( password string ) string {
	for !validatePassword( password ) {
		password = incrementPassword( password )
	}
	return password
}

func main() {
	testValidPasswords()
	testInvalidPasswords()
	fmt.Println( "Tests complete." )
	password := "hepxcrrq"
	for {
		password = incrementPassword( password )
		if validatePassword( password ) {
			break
		}
	}
	fmt.Println( "Next valid password: ", password )
}