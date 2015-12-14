package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Instruction struct {
	OnOffToggle uint32 `json:"on-off-toggle"`
	StartX uint32 `json:"start-x"`
	StartY uint32 `json:"start-y"`
	EndX uint32 `json:"end-x"`
	EndY uint32 `json:"end-y"`
}

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func resetInstructionLessToMore( instruction Instruction ) Instruction {
	lessToMoreInstruction := Instruction{}
	lessToMoreInstruction.OnOffToggle = instruction.OnOffToggle
	if instruction.StartX < instruction.EndX {
		lessToMoreInstruction.StartX = instruction.StartX
		lessToMoreInstruction.EndX = instruction.EndX
	} else {
		lessToMoreInstruction.StartX = instruction.EndX
		lessToMoreInstruction.EndX = instruction.StartX
	}
	if instruction.StartY < instruction.EndY {
		lessToMoreInstruction.StartY = instruction.StartY
		lessToMoreInstruction.EndY = instruction.EndY
	} else {
		lessToMoreInstruction.StartY = instruction.EndY
		lessToMoreInstruction.EndY = instruction.StartY
	}
	return lessToMoreInstruction
}

func runLightInstructions( instructions []Instruction, lights [][]bool ) {
	for _, initialInstruction := range instructions {
		instruction := resetInstructionLessToMore( initialInstruction )
		fmt.Println(instruction)
		for xi := instruction.StartX; xi <= instruction.EndX; xi++ {
			for yi := instruction.StartY; yi <= instruction.EndY; yi++ {
				if instruction.OnOffToggle == 0 {
					lights[ xi ][ yi ] = false
				}	else if instruction.OnOffToggle == 1 {
					lights[ xi ][ yi ] = true
				} else if instruction.OnOffToggle == 2 {
					lights[ xi ][ yi ] = !lights[ xi ][ yi ]
				}
			}
		}
	}
}

func countLightsOn( lights [][]bool ) {
	totalLightsOn := 0
	for _, lightsRow := range lights {
		for _, light := range lightsRow {
			if ( light ) {
				totalLightsOn++
			}
		}
	}
	fmt.Print( "Total lights on: " )
	fmt.Println( totalLightsOn )
}

func main() {
	width, height := 1000, 1000
	lights := make( [][]bool, height )
	lightsRows := make( []bool, height * width )
	for i := range lights {
		lights[ i ], lightsRows = lightsRows[ :width ], lightsRows[ width: ]
	}
	
	stdin, error := ioutil.ReadAll( os.Stdin )
	exitOnError( error )
	instructions := []Instruction{}
	error = json.Unmarshal( stdin, &instructions )
	
	exitOnError( error )
	runLightInstructions( instructions, lights )
	countLightsOn( lights )
}