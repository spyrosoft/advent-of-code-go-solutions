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
	happinessLookup = make( map[string]map[string]int )
	people = list.New()
	optimalHappiness = 0
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func personExists( person string ) bool {
	for element := people.Front(); element != nil; element = element.Next() {
		if element.Value == person {
			return true
		}
	}
	return false
}

func buildHappinessLookupAndPeopleList() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner( file )
	
	for scanner.Scan() {
		input := strings.Split( scanner.Text(), " " )
		
		happinessRating, error := strconv.Atoi( input[ 3 ] )
		exitOnError( error )
		if input[ 2 ] == "lose" {
			happinessRating = 0 - happinessRating
		}
		_, personExistsInHappiness := happinessLookup[ input[ 0 ] ]
		if !personExistsInHappiness {
			happinessLookup[ input[ 0 ] ] = make( map[string]int )
		}
		happinessLookup[ input[ 0 ] ][ input[ 10 ] ] = int( happinessRating )
		
		if !personExists( input[ 0 ] ) {
			people.PushBack( input[ 0 ] )
		}
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

func copyList( listToCopy *list.List ) *list.List {
	listCopy := list.New()
	for element := listToCopy.Front(); element != nil; element = element.Next() {
		listCopy.PushBack( element.Value )
	}
	return listCopy
}

func findOptimalHappiness( arrangement *list.List, remainingPeople *list.List ) {
	if remainingPeople.Len() == 0 {
		arrangement.PushBack( arrangement.Front().Value )
		happiness := 0
		for person := arrangement.Front(); person != nil; person = person.Next() {
			if person.Next() == nil { break }
			happiness += happinessLookup[ person.Value.(string) ][ person.Next().Value.(string) ]
			happiness += happinessLookup[ person.Next().Value.(string) ][ person.Value.(string) ]
		}
		if happiness > optimalHappiness {
			optimalHappiness = happiness
		}
		return
	}
	
	for person := remainingPeople.Front(); person != nil; person = person.Next() {
		arrangementCopy := copyList( arrangement )
		arrangementCopy.PushBack( person.Value.(string) )
		nextRemainingPeople := copyListExceptStringValue( remainingPeople, person.Value.(string) )
		findOptimalHappiness( arrangementCopy, nextRemainingPeople )
	}
}

func main() {
	buildHappinessLookupAndPeopleList()
	firstPerson := people.Front()
	remainingPeople := copyListExceptStringValue( people, firstPerson.Value.(string) )
	arrangement := list.New()
	arrangement.PushBack( firstPerson.Value )
	findOptimalHappiness( arrangement, remainingPeople )
	fmt.Println( "Result: ", optimalHappiness )
}