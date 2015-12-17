package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func main() {
	input, error := ioutil.ReadAll( os.Stdin )
	exitOnError( error )
	var floorNumber int32 = 0
	for index, character := range input {
		if character == '(' {
			floorNumber++
		} else {
			floorNumber--
		}
		if floorNumber < 0 {
			fmt.Println( "We did it!" )
			fmt.Println( index + 1 )
			break
		}
	}
}