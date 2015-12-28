package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

var (
	DIMENSIONS = 100
	//DIMENSIONS = 6
)

func main() {
	lightGrid := initializeLightGrid()
	for i := 0; i < 100; i++ {
		lightGrid = animate( lightGrid )
	}
	fmt.Println( countLightsOn( lightGrid ) )
}

func animate( oldLightGrid [][]bool ) ( lightGrid [][]bool ) {
	lightGrid = copyLightGrid( oldLightGrid )
	for rowIndex, row := range lightGrid {
		for columnIndex, cell := range row {
			howManyNeighborsAreOn := countHowManyNeighborsAreOn( oldLightGrid, rowIndex, columnIndex )
			if rowIndex == 1 && columnIndex == 2 {
			}
			if cell {
				if howManyNeighborsAreOn == 2 || howManyNeighborsAreOn == 3 {
					lightGrid[ rowIndex ][ columnIndex ] = true
				} else {
					lightGrid[ rowIndex ][ columnIndex ] = false
				}
			} else {
				if howManyNeighborsAreOn == 3 {
					lightGrid[ rowIndex ][ columnIndex ] = true
				}
			}
		}
	}
	return
}

func countHowManyNeighborsAreOn( lightGrid [][]bool, y int, x int ) (neighbors int) {
	if y > 0 {
		//.?.
		//.+.
		//...
		if lightGrid[ y - 1 ][ x ] { neighbors++ }
		//?x.
		//.+.
		//...
		if x > 0 { if lightGrid[ y - 1 ][ x - 1 ] { neighbors++ } }
		//xx?
		//.+.
		//...
		if x < DIMENSIONS - 1 { if lightGrid[ y - 1 ][ x + 1 ] { neighbors++ } }
	}
	if x > 0 {
		//xxx
		//?+.
		//...
		if lightGrid[ y ][ x - 1 ] { neighbors++ }
		//xxx
		//x+.
		//?..
		if y < DIMENSIONS - 1 { if lightGrid[ y + 1 ][ x - 1 ] { neighbors++ } }
	}
	if y < DIMENSIONS - 1 {
		//xxx
		//x+.
		//x?.
		if lightGrid[ y + 1 ][ x ] { neighbors++ }
		//xxx
		//x+.
		//xx?
		if x < DIMENSIONS - 1 { if lightGrid[ y + 1 ][ x + 1 ] { neighbors++ } }
	}
	if x < DIMENSIONS - 1 {
		//xxx
		//x+?
		//xxx
		if lightGrid[ y ][ x + 1 ] { neighbors++ }
	}
	return
}

func countLightsOn( lightGrid [][]bool ) ( total int ) {
	for _, row := range lightGrid {
		for _, cell := range row {
			if cell {
				total++
			}
		}
	}
	return
}

func initializeLightGrid() [][]bool {
	lightGrid := buildLightGrid()
	populateLightGrid( lightGrid )
	return lightGrid
}

func populateLightGrid( lightGrid [][]bool ) {
	stdin, error := ioutil.ReadAll( os.Stdin )
	panicOnError( error )
	instructions := strings.Split( string( stdin ), "\n" )
	for rowIndex, instructionRow := range instructions {
		for columnIndex, instruction := range instructionRow {
			if instruction == '#' {
				lightGrid[ rowIndex ][ columnIndex ] = true
			} else if instruction == '.' {
				lightGrid[ rowIndex ][ columnIndex ] = false
			}
		}
	}
}

func buildLightGrid() ( lightGrid [][]bool ) {
	lightGrid = make( [][]bool, DIMENSIONS )
	lightGridRows := make( []bool, DIMENSIONS * DIMENSIONS )
	for i := range lightGrid {
		lightGrid[ i ], lightGridRows = lightGridRows[ :DIMENSIONS ], lightGridRows[ DIMENSIONS: ]
	}
	return
}

func copyLightGrid( oldLightGrid [][]bool ) ( lightGrid [][]bool ) {
	//This could be optimized by creating two light grids initially and toggling between them
	lightGrid = buildLightGrid()
	for rowIndex, row := range lightGrid {
		for columnIndex, _ := range row {
			lightGrid[ rowIndex ][ columnIndex ] = oldLightGrid[ rowIndex ][ columnIndex ]
		}
	}
	return
}





func printGrid( grid [][]bool ) {
	for _, row := range grid {
		fmt.Print( "[" )
		for _, cell := range row {
			if cell { fmt.Print( "#" ) } else { fmt.Print( "." ) }
		}
		fmt.Println( "]" )
	}
	fmt.Println()
}

func panicOnError( error error ) { if error != nil { panic( error ) } }