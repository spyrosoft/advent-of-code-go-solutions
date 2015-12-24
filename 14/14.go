package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var (
	reindeerDistances = make( map[string][]uint16 )
	reindeerFlying = make( map[string]bool )
	speedIndex = 0
	flightIndex = 1
	restIndex = 2
	flownIndex = 3
	restedIndex = 4
	totalFlightIndex = 5
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func convertStringToUint16( string string ) uint16 {
	int, error := strconv.Atoi( string )
	exitOnError( error )
	return uint16( int )
}

func buildReindeerDistances() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner( file )
	
	for scanner.Scan() {
		input := strings.Split( scanner.Text(), " " )
		
		reindeer := input[ 0 ]
		
		_, reindeerExists := reindeerDistances[ reindeer ]
		if !reindeerExists {
			reindeerDistances[ reindeer ] = make( []uint16, 6 )
		}
		
		reindeerDistances[ reindeer ][ speedIndex ] = convertStringToUint16( input[ 3 ] )
		reindeerDistances[ reindeer ][ flightIndex ] = convertStringToUint16( input[ 6 ] )
		reindeerDistances[ reindeer ][ restIndex ] = convertStringToUint16( input[ 13 ] )
	}
}

func buildReindeerFlying() {
	for reindeer, _ := range reindeerDistances {
		reindeerFlying[ reindeer ] = true
	}
}

func raceRaindeer() {
	for i := 0; i < 2503; i++ {
		for reindeer, distances := range reindeerDistances {
			if reindeerFlying[ reindeer ] {
				distances[ flownIndex ]++
				distances[ totalFlightIndex ]++
				if distances[ flownIndex ] >= distances[ flightIndex ] {
					reindeerFlying[ reindeer ] = false
					distances[ flownIndex ] = 0
				}
			} else {
				distances[ restedIndex ]++
				if distances[ restedIndex ] >= distances[ restIndex ] {
					reindeerFlying[ reindeer ] = true
					distances[ restedIndex ] = 0
				}
			}
		}
	}
}

func main() {
	buildReindeerDistances()
	buildReindeerFlying()
	raceRaindeer()
	for raindeer, distances := range reindeerDistances {
		fmt.Println( raindeer, distances[ totalFlightIndex ] * distances[ speedIndex ] )
	}
}