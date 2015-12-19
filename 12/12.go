package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"strconv"
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
	isNumeric := regexp.MustCompile( "[-0-9]+" )
	allNumbers := isNumeric.FindAllString( string( input ), -1 )
	total := 0
	for _, number := range allNumbers {
		currentNumber, error := strconv.Atoi( number )
		exitOnError( error )
		total += currentNumber
	}
	fmt.Println( total )
}