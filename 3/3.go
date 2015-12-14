package main

import (
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func main() {
	var x, y, totalDeliveries int = 0, 0, 0
	deliveries := make( map[string]uint32 )
	
	deliveries[ "0:0" ]++
	
	input, error := ioutil.ReadAll( os.Stdin )
	exitOnError( error )
	
	for _, character := range input {
		if character == '^' {
			y++
		} else if character == '>' {
			x++
		} else if character == 'v' {
			y--
		} else if character == '<' {
			x--
		}
		deliveries[ strconv.Itoa( x ) + ":" + strconv.Itoa( y ) ]++
	}
	
	for key, value := range deliveries {
		fmt.Print(key)
		fmt.Print("::")
		fmt.Println(value)
		totalDeliveries++
	}
	
	fmt.Println( totalDeliveries )
}