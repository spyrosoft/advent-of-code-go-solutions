package main

import (
	"fmt"
	"strconv"
)

func lookAndSay( input string ) string {
	currentStreak := 1
	output := ""
	for i, digit := range input {
		if ( i + 1 < len( input ) ) {
			if rune( input[ i + 1 ] ) == digit {
				currentStreak++
			} else {
				output += strconv.Itoa( currentStreak ) + string( digit )
				currentStreak = 1
			}
		} else {
			output += strconv.Itoa( currentStreak ) + string( digit )
			currentStreak = 1
		}
	}
	return output
}

func main() {
	//seed := "1"
	seed := "1321131112"
	for i := 0; i < 40; i++ {
		seed = lookAndSay( seed )
	}
	fmt.Println( len( seed ) )
}