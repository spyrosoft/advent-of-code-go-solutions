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
	instructions = make(map[string][]string)
	memorizedWires = make(map[string]string)
	MAX_UINT_16 uint16 = 65535
	POWERS_OF_TWO = []uint16{32768, 16384, 8192, 4096, 2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	INTEGER_REGEX = regexp.MustCompile( "^[0-9]+$" )
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


func evaluateNotInstruction( notInstruction string ) string {
	wireStringNumber := evaluateWire( notInstruction )
	wireUint64, error := strconv.ParseUint( wireStringNumber, 10, 16 )
	exitOnError( error )
	wireUint16 := uint16( wireUint64 )
	notWire := MAX_UINT_16 - wireUint16
	return strconv.Itoa( int( notWire ) )
}

func convertNumberToBinaryArray( numberToConvert string ) []bool {
	binaryArray := make( []bool, 16 )
	number64, error := strconv.Atoi( numberToConvert )
	exitOnError( error )
	number := uint16( number64 )
	for i, powerOfTwo := range POWERS_OF_TWO {
		if number >= powerOfTwo {
			number -= powerOfTwo
			binaryArray[ i ] = true
		} else {
			binaryArray[ i ] = false
		}
	}
	return binaryArray
}

func andBinaryArrays( first []bool, second []bool ) []bool {
	if len( first ) != len( second ) {
		panic( "andBinaryArrays arguments need to be the same length." )
	}
	andResult := make( []bool, 16 )
	for i, _ := range first {
		if first[ i ] && second[ i ] {
			andResult[ i ] = true
		} else {
			andResult[ i ] = false
		}
	}
	return andResult
}

func convertBinaryArrayToNumber( binary []bool ) string {
	number := 0
	for index, bit := range binary {
		if bit {
			number += int( POWERS_OF_TWO[ index ] )
		}
	}
	return strconv.Itoa( number )
}

func evaluateAndInstruction( andInstructionFirst string, andInstructionSecond string ) string {
	firstBinary := convertNumberToBinaryArray( andInstructionFirst )
	secondBinary := convertNumberToBinaryArray( andInstructionSecond )
	andBinary := andBinaryArrays( firstBinary, secondBinary )
	return convertBinaryArrayToNumber( andBinary )
}

func orBinaryArrays( first []bool, second []bool ) []bool {
	if len( first ) != len( second ) {
		panic( "andBinaryArrays arguments need to be the same length." )
	}
	andResult := make( []bool, 16 )
	for i, _ := range first {
		if first[ i ] || second[ i ] {
			andResult[ i ] = true
		} else {
			andResult[ i ] = false
		}
	}
	return andResult
}

func evaluateOrInstruction( andInstructionFirst string, andInstructionSecond string ) string {
	firstBinary := convertNumberToBinaryArray( andInstructionFirst )
	secondBinary := convertNumberToBinaryArray( andInstructionSecond )
	orBinary := orBinaryArrays( firstBinary, secondBinary )
	return convertBinaryArrayToNumber( orBinary )
}

func rshift( numberInput string, shiftInput string ) string {
	number, error := strconv.Atoi( numberInput )
	exitOnError( error )
	shift, error := strconv.Atoi( shiftInput )
	exitOnError( error )
	result := number>>uint16( shift )
	return strconv.Itoa( int( result ) )
}

func lshift( numberInput string, shiftInput string ) string {
	number, error := strconv.Atoi( numberInput )
	exitOnError( error )
	shift, error := strconv.Atoi( shiftInput )
	exitOnError( error )
	result := number<<uint16( shift )
	return strconv.Itoa( int( result ) )
}

func memorizeWire( wire string ) string {
	if len( INTEGER_REGEX.FindString( wire ) ) > 0 {
		return wire
	} else if len( instructions[ wire ] ) == 1 {
		return evaluateWire( instructions[ wire ][ 0 ] )
	} else if instructions[ wire ][ 0 ] == "NOT" {
		return evaluateNotInstruction( evaluateWire( instructions[ wire ][ 1 ] ) )
	} else if instructions[ wire ][ 1 ] == "AND" {
		return evaluateAndInstruction( evaluateWire( instructions[ wire ][ 0 ] ), evaluateWire( instructions[ wire ][ 2 ] ) )
	} else if instructions[ wire ][ 1 ] == "OR" {
		return evaluateOrInstruction( evaluateWire( instructions[ wire ][ 0 ] ), evaluateWire( instructions[ wire ][ 2 ] ) )
	} else if instructions[ wire ][ 1 ] == "RSHIFT" {
		return rshift( evaluateWire( instructions[ wire ][ 0 ] ), evaluateWire( instructions[ wire ][ 2 ] ) )
	} else if instructions[ wire ][ 1 ] == "LSHIFT" {
		return lshift( evaluateWire( instructions[ wire ][ 0 ] ), evaluateWire( instructions[ wire ][ 2 ] ) )
	} else {
		panic( "Invalid instruction: " + wire )
	}
}

func evaluateWire( wire string ) string {
	fmt.Print( "wire: " )
	fmt.Println( wire )
	previouslyMemorizedWire, memorizedWireExists := memorizedWires[ wire ]
	if memorizedWireExists {
		return previouslyMemorizedWire
	} else {
		memorizedWires[ wire ] = memorizeWire( wire )
		return memorizedWires[ wire ]
	}
}

func main() {
	buildInstructions()
	fmt.Println( "Result: " + evaluateWire( "a" ) )
}