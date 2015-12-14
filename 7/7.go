package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"regexp"
)

var (
	instructions map[string][]string
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func buildInstructions() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		wireRegex := regexp.MustCompile( "[^ ]+$" )
		instructionRegex := regexp.MustCompile( "^(.+) ->" )
		wire := wireRegex.FindString( scanner.Text() )
		instruction := strings.Split( strings.Trim( instructionRegex.FindString( scanner.Text() ), " ->" ), " " )
		instructions[ wire ] = instruction
	}
}

func evaluateWire( wire string ) string {
	
	if len( instructions[ wire ] ) == 1 {
		
		integerRegex := regexp.MustCompile( "^[0-9]+$" )
		if len( integerRegex.FindString( instructions[ wire ][ 0 ] ) ) > 0 {
			return instructions[ wire ][ 0 ]
		} else {
			return evaluateWire( instructions[ wire ][ 0 ] )
		}
		
	} else if instructions[ wire ][ 0 ] == "NOT" {
		
		return strconv.Itoa( int( !strconv.ParseUint( evaluateWire( instructions[ wire ][ 1 ] ), 10, 16 ) ) )
		
	} else if instructions[ wire ][ 1 ] == "AND" {
		
	}
	
	return strconv.Itoa(4)
}

func main() {
	var a, b uint16 = 0, 1
	fmt.Println( a & b )
	//buildInstructions()
	//evaluateWire( "a" )
}