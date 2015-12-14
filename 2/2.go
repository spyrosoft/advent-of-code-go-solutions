package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Boxes struct {
	Dimensions [][ 3 ]uint32 `json:"dimensions"`
}

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func main() {
	var boxes = Boxes{}
	input, error := ioutil.ReadAll( os.Stdin )
	exitOnError( error )
	error = json.Unmarshal( input, &boxes )
	exitOnError( error )
	var totalSquareFeet uint32 = 0
	for _, box := range boxes.Dimensions {
		length := box[ 0 ]
		width := box[ 1 ]
		height := box[ 2 ]
		side1 := length*width
		side2 := width*height
		side3 := length*height
		least := side1
		
		if side2 < least {
			least = side2
		}
		if side3 < least {
			least = side3
		}
		totalSquareFeet += ( 2*side1 + 2*side2 + 2*side3 ) + least
	}
	fmt.Println( totalSquareFeet )
}