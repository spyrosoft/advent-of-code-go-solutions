package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"container/list"
)

var (
	destinations = list.New()
	distances = make( map[string]uint16 )
	shortestDistance uint16 = 65535
)


func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}


func destinationExists( destination string ) bool {
	for element := destinations.Front(); element != nil; element = element.Next() {
		if element.Value == destination {
			return true
		}
	}
	return false
}

func buildDistancesAndDestinations() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner( file )
	
	for scanner.Scan() {
		inputValues := strings.Split( scanner.Text(), " " )
		
		if !destinationExists( inputValues[ 0 ] ) {
			destinations.PushBack( inputValues[ 0 ] )
		}
		if !destinationExists( inputValues[ 2 ] ) {
			destinations.PushBack( inputValues[ 2 ] )
		}

		distance, error := strconv.Atoi( inputValues[ 4 ] )
		exitOnError( error )
		distances[ inputValues[ 0 ] + inputValues[ 2 ] ] = uint16( distance )
		distances[ inputValues[ 2 ] + inputValues[ 0 ] ] = uint16( distance )
	}
}

func copyListExceptStringValue( listToCopy *list.List, exceptStringValue string ) *list.List {
	listCopy := list.New()
	for element := listToCopy.Front(); element != nil; element = element.Next() {
		if element.Value != exceptStringValue {
			listCopy.PushBack( element.Value )
		}
	}
	return listCopy
}

func findShortestPath( currentDistance uint16, currentDestination string, remainingDestinations *list.List ) {
	if remainingDestinations.Len() == 0 {
		if currentDistance < shortestDistance {
			shortestDistance = currentDistance
		}
		return
	}
	
	for destination := remainingDestinations.Front(); destination != nil; destination = destination.Next() {
		newRemainingDestinations := copyListExceptStringValue( remainingDestinations, destination.Value.(string) )
		findShortestPath( currentDistance + distances[ currentDestination + destination.Value.(string) ], destination.Value.(string), newRemainingDestinations )
	}
}

func main() {
	buildDistancesAndDestinations()
	for destination := destinations.Front(); destination != nil; destination = destination.Next() {
		initialDestinationsList := copyListExceptStringValue( destinations, destination.Value.(string) )
		findShortestPath( 0, destination.Value.(string), initialDestinationsList )
	}
	fmt.Println( "Result: ", shortestDistance )
}