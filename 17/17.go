package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"container/list"
)

func main() {
	MAX_TOTAL := 150
	stdin, error := ioutil.ReadAll( os.Stdin )
	panicOnError( error )
	containers := convertStringArrayToUint8Array( strings.Split( string( stdin ), "\n" ) )
	i := 0
	total := 0
	combitanionsOfContainers := 0
	stack := list.New()
	limit := 100
	for {
		total += int( containers[ i ] )
		if total == MAX_TOTAL {
			combitanionsOfContainers++
		}
		if total >= MAX_TOTAL {
			total -= int( containers[ i ] )
			stack.Remove( stack.Back() )
		}
		stack.PushBack( i )
		i++
		fmt.Println("i",i)
		if i == len( containers ) {
			printList( stack )
			if stack.Len() == 1 {
				fmt.Println( "--", containers[ stack.Front().Value.(int) ] )
				if stack.Front().Value.(int) + 1 >= len( containers ) { break }
				stack.PushFront( stack.Front().Value.(int) + 1 )
			}
			stack.Remove( stack.Back() )
			i = stack.Front().Value.(int)
		}
		
		limit++
		if limit > 1000 {
			println( "Limit reached." )
			break
		}
	}
	fmt.Println( combitanionsOfContainers )
}



func printList( listToPrint *list.List ) {
	fmt.Print( "( " )
	for element := listToPrint.Front(); element != nil; element = element.Next() {
		fmt.Print( element.Value, " " )
	}
	fmt.Println( ")" )
}

func convertStringArrayToUint8Array( stringArray []string ) []uint8 {
	uint8Array := make( []uint8, len( stringArray ) )
	for i, string := range stringArray {
		uint8Array[ i ] = stringToUint8( string )
	}
	return uint8Array
}

func stringToUint8( string string ) uint8 {
	integer, error := strconv.Atoi( string )
	panicOnError( error )
	uint8 := uint8( integer )
	if int( uint8 ) != integer {
		panic( "Call to stringToUint8 on a non uint8 string." )
	}
	return uint8
}

func panicOnError( error error ) { if error != nil { panic( error ) } }