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

func main() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	totalCharacters := 0
	
	for scanner.Scan() {
		currentLine := scanner.Text()
		totalCharacters += len( currentLine )
	}
	fmt.Println( totalCharacters )
}

/* And the corresponding JavaScript:
a = [...input wrapped in brackets with commas...]
for (var i = 0; i < a.length; i++) { t+=a[i].length; }
6195 - 4845 = 1350