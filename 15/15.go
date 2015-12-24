package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"container/list"
	"regexp"
)

var (
	permutationsArray [][]int
	ingredientsList = list.New()
	ingredientsArray [][]int
	nonAlphaNumericRegex = regexp.MustCompile( "[^a-zA-Z0-9- ]" )
	optimalPermutationValue = 0
	capacityIndex = 0
	durabilityIndex = 1
	flavorIndex = 2
	textureIndex = 3
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func convertIngredientsListToArray() {
	ingredientsArray = make( [][]int, ingredientsList.Len() )
	i := 0
	for element := ingredientsList.Front(); element != nil; element = element.Next() {
		ingredientsArray[ i ] = element.Value.([]int)
		i++
	}
}

func buildIngredients() {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner( file )
	
	for scanner.Scan() {
		alphaNumericInput := nonAlphaNumericRegex.ReplaceAllString( scanner.Text(), "" )
		input := strings.Split( alphaNumericInput, " " )
		
		capacity, error := strconv.Atoi( input[ 2 ] )
		exitOnError( error )
		durability, error := strconv.Atoi( input[ 4 ] )
		exitOnError( error )
		flavor, error := strconv.Atoi( input[ 6 ] )
		exitOnError( error )
		texture, error := strconv.Atoi( input[ 8 ] )
		exitOnError( error )
		ingredientsList.PushBack( []int{ capacity, durability, flavor, texture } )
	}
	
	convertIngredientsListToArray()
}

func addPermutation( permutations *list.List ) {
	
}

func buildPermutations( numberOfPossibilities int, numberOfAttributes int ) {
	permutations := list.New()
	
	firstPermutation := list.New()
	for i := 0; i < ingredientsList.Len(); i++ {
		firstPermutation.PushBack( 1 )
	}
	firstPermutation.Front().Value = numberOfPossibilities - ( numberOfAttributes - 1 )
	
	permutations.PushBack( firstPermutation )
	addPermutation( permutations )
}

func checkForOptimalIngredients( permutations []int ) {
	capacityTotal := 0
	durabilityTotal := 0
	flavorTotal := 0
	textureTotal := 0
	
	fmt.Println( permutations )
	for i := 0; i < ingredientsList.Len(); i++ {
		capacityTotal += permutations[ i ] * ingredientsArray[ i ][ capacityIndex ]
		durabilityTotal += permutations[ i ] * ingredientsArray[ i ][ durabilityIndex ]
		flavorTotal += permutations[ i ] * ingredientsArray[ i ][ flavorIndex ]
		textureTotal += permutations[ i ] * ingredientsArray[ i ][ textureIndex ]
	}
	
	if capacityTotal < 0 { capacityTotal = 0 }
	if durabilityTotal < 0 { durabilityTotal = 0 }
	if flavorTotal < 0 { flavorTotal = 0 }
	if textureTotal < 0 { textureTotal = 0 }
	
	permutationProduct := capacityTotal * durabilityTotal * flavorTotal * textureTotal
	
	if permutationProduct > optimalPermutationValue {
		optimalPermutationValue = permutationProduct
	}
}

// func permuteCookies( numberOfPossibilities int, numberOfAttributes int ) {
// 	max := numberOfPossibilities - ( numberOfAttributes - 1 )
// 	min := 1
// 	permutations := make( []int, numberOfAttributes )
// 	for i := 0; i < numberOfAttributes; i++ {
// 		permutations[ i ] = min
// 	}
// 	permutations[ 0 ] = max
// 	carry := 0
// 	index := numberOfAttributes - 1
// //	checkForOptimalIngredients( permutations )
// 	for {
// 		if permutations[ index ] > min {
// 			carry++
// 			permutations[ index ]--
// 			permutations[ index + 1 ] += carry
// 			index = numberOfAttributes - 1
// 			carry = 0
// 			fmt.Println(permutations)
// //			checkForOptimalIngredients( permutations )
// 		}
// 		if index == 0 && permutations[ index ] == min {
// 			break
// 		}
// 		index--
// 	}
// }

func main() {
	buildIngredients()
	buildPermutations( 8, 4 )
	//permuteCookies( 100, ingredientsList.Len() )
	//permuteCookies( 8, 4 )
	fmt.Println( optimalPermutationValue )
}