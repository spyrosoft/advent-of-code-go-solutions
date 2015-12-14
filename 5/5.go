package main

import (
	"fmt"
	"os"
	"bufio"
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func isBadSequence( previousCharacter rune, currentCharacter rune ) bool {
	if ( previousCharacter == 'a' && currentCharacter == 'b' || previousCharacter == 'c' && currentCharacter == 'd' || previousCharacter == 'p' && currentCharacter == 'q' || previousCharacter == 'x' && currentCharacter == 'y' ) {
		return true;
	}
	return false
}

func main() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	vowels := "aoeui"
	niceStrings := 0
	
	previousCharacter := '~'
	vowelCount := 0
	doubleLetter := false
	badSequence := false
	
	for scanner.Scan() {
		previousCharacter = '~'
		vowelCount = 0
		doubleLetter = false
		badSequence = false
		
		for _, character := range scanner.Text() {
			if isBadSequence( previousCharacter, character ) {
				badSequence = true
			}
			if character == previousCharacter {
				doubleLetter = true
			}
			for _, vowel := range vowels {
				if character == vowel {
					vowelCount++
				}
			}
			previousCharacter = character
		}
		if !badSequence && vowelCount >= 3 && doubleLetter {
			niceStrings++
		}
	}
	fmt.Println( "Nice Strings:" )
	fmt.Println( niceStrings )
}