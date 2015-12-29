package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"regexp"
)

func main() {
	moleculeLookup := populateMoleculeLookup()
	fmt.Println( howManyDistinctMolecules( moleculeLookup ) )
}

func howManyDistinctMolecules( moleculeLookup map[string][]string ) ( totalDistinctMolecules int ) {
	initialMolecule := "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
	distinctMolecules := make( map[string]bool )
	
	for atom, expansions := range moleculeLookup {
		atomRegex := regexp.MustCompile( atom )
		matchingIndexes := atomRegex.FindAllStringIndex( initialMolecule, -1 )
		if len( matchingIndexes ) == 0 {
			fmt.Println( "Not Used: ", atom )
			continue
		}
		
		for _, matchingIndex := range matchingIndexes {
			for _, expansion := range expansions {
				fmt.Println( "Expansion:", atom, expansion )
				newMolecule := initialMolecule[ :matchingIndex[ 0 ] ] + expansion + initialMolecule[ matchingIndex[ 1 ]: ]
				distinctMolecules[ newMolecule ] = true
			}
		}
	}
	
	return len( distinctMolecules )
}

func populateMoleculeLookup() map[string][]string {
	stdin, error := ioutil.ReadAll( os.Stdin )
	panicOnError( error )
	instructions := strings.Split( string( stdin ), "\n" )
	moleculeLookup := make( map[string][]string )
	for _, instruction := range instructions {
		var atom, expansion string
		_, error := fmt.Sscanf( instruction, "%s => %s", &atom, &expansion )
		panicOnError( error )
		if !mapKeyExists( moleculeLookup, atom ) {
			moleculeLookup[ atom ] = make( []string, 1 )
			moleculeLookup[ atom ][ 0 ] = expansion
		} else {
			moleculeLookup[ atom ] = append( moleculeLookup[ atom ], expansion )
		}
	}
	return moleculeLookup
}



func mapKeyExists( testMap map[string][]string, testKey string ) bool {
	_, keyExists := testMap[ testKey ]
	return keyExists
}

func panicOnError( error error ) { if error != nil { panic( error ) } }