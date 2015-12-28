package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main() {
	containers := populateContainers()
	MAX_VOLUME := 150
	//MAX_VOLUME := 30
	fmt.Println( howManyCombinationsOfContainers( containers, MAX_VOLUME ) )
}

func populateContainers() []int {
	stdin, error := ioutil.ReadAll( os.Stdin )
	panicOnError( error )
	containers := convertStringArrayToIntArray( strings.Split( string( stdin ), "\n" ) )
	sort.Ints( containers )
	reverseIntArray( containers )
	return containers
}

func howManyCombinationsOfContainers( containers []int, MAX_VOLUME int ) int {
	var i, totalCombos, currentTotal int
	permutation := make( []bool, len( containers ) )
	permutation[ i ] = true

	for {
		currentTotal = containersTotal( containers, permutation )
		if currentTotal == MAX_VOLUME {
			totalCombos++
		}
		if i == len( containers ) - 1 {
			clearContiguousIFromRight( permutation )
			i = firstOnFromRight( permutation )
			if i == -1 { break }
			permutation[ i ] = false
		}
		if currentTotal >= MAX_VOLUME {
			permutation[ i ] = false
		}
		i++
		permutation[ i ] = true
	}
	return totalCombos
}

func clearContiguousIFromRight( permutation []bool ) {
	contiguous := true
	i := len( permutation )
	for {
		i--
		if i == -1 { break }
		if permutation[ i ] == false {
			contiguous = false
		} else if !contiguous {
			break
		} else {
			permutation[ i ] = false
		}
	}
}

func firstOnFromRight( permutation []bool ) int {
	for i := len( permutation ) - 1; i >= 0; i-- {
		if permutation[ i ] {
			return i
		}
	}
	return -1
}

func containersTotal( containers []int, permutation []bool ) int {
	if len( containers ) != len( permutation ) {
		panic( "Contiainers is not the same length as permutation." )
	}
	total := 0
	for i, onOrOff := range permutation {
		if onOrOff {
			total += containers[ i ]
		}
	}
	return total
}




func convertStringArrayToIntArray( stringArray []string ) []int {
	intArray := make( []int, len( stringArray ) )
	for i, string := range stringArray {
		intArray[ i ] = stringToInt( string )
	}
	return intArray
}

func reverseIntArray( arrayToReverse []int ) {
	for i, _ := range arrayToReverse {
		if i >= len( arrayToReverse ) / 2 { break }
		tmp := arrayToReverse[ i ]
		arrayToReverse[ i ] = arrayToReverse[ len( arrayToReverse ) - 1 - i ]
		arrayToReverse[ len( arrayToReverse ) - 1 - i ] = tmp
	}
}

func stringToInt( string string ) int {
	integer, error := strconv.Atoi( string )
	panicOnError( error )
	return integer
}

func panicOnError( error error ) { if error != nil { panic( error ) } }