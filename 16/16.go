package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

var (
	knownSue = map[string]uint16{ "children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1 }
	inputSueRegex = make( map[string]*regexp.Regexp )
)

func main() {
	stdin, error := ioutil.ReadAll( os.Stdin )
	panicOnError( error )
	buildInputSueRegex()
	sues := strings.Split( string( stdin ), "\n" )
	fmt.Println( matchSue( sues ) )
}

func buildInputSueRegex() {
	for identifier, _ := range knownSue {
		inputSueRegex[ identifier ] = regexp.MustCompile( identifier + ": ([0-9]+)" )
	}
}

func matchSue( sues []string ) uint16 {
	sueRegex := regexp.MustCompile( "Sue ([0-9]+): " )
	
	var correct, incorrect bool
	var sueIndex uint16
	for _, sue := range sues {
		correct = false
		incorrect = false
		
		// Skip lines which do not have a Sue
		sueIndexMatch := sueRegex.FindStringSubmatch( sue )
		if len( sueIndexMatch ) < 2 { continue }
		
		sueIndex = stringToUint16( sueRegex.FindStringSubmatch( sue )[ 1 ] )
		
		for identifier, regex := range inputSueRegex {
			identifierMatch := regex.FindStringSubmatch( sue )
			if len( identifierMatch ) > 1 {
				if stringToUint16( identifierMatch[ 1 ] ) == knownSue[ identifier ] {
					correct = true
				} else {
					incorrect = true
				}
			}
		}
		if correct && !incorrect { break }
	}
	return sueIndex
}




func stringToUint16( string string ) uint16 {
	integer, error := strconv.Atoi( string )
	panicOnError( error )
	uint16 := uint16( integer )
	if int( uint16 ) != integer {
		panic( "Call to stringToUint16 on a larger number than fits." )
	}
	return uint16
}

func panicOnError( error error ) { if error != nil { panic( error ) } }