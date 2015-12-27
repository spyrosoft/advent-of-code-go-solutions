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
	MIN = 1
)

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func stringToInteger( string string ) int {
	integer, error := strconv.Atoi( string )
	exitOnError( error )
	return integer
}

func buildList( integers ...int ) *list.List {
	newList := list.New()
	for _, integer := range integers {
		newList.PushBack( integer )
	}
	return newList
}

func listNth( list *list.List, nth int ) *list.Element {
	element := list.Front()
	for i := 0; i < nth; i++ {
		element = element.Next()
	}
	return element
}

func copyList( listToCopy *list.List ) *list.List {
	listCopy := list.New()
	for element := listToCopy.Front(); element != nil; element = element.Next() {
		listCopy.PushBack( element.Value )
	}
	return listCopy
}

func printList( listToPrint *list.List ) {
	fmt.Print( "( " )
	for element := listToPrint.Front(); element != nil; element = element.Next() {
		fmt.Print( element.Value, " " )
	}
	fmt.Println( ")" )
}

func convert2DListTo2DArray( twoDimensionalList *list.List ) [][]int {
	convertedArray := make( [][]int, twoDimensionalList.Len() )
	i := 0
	for element := twoDimensionalList.Front(); element != nil; element = element.Next() {
		secondDimension := make( []int, element.Value.(*list.List).Len() )
		secondDimensionIndex := 0
		for secondDimensionElement := element.Value.(*list.List).Front(); secondDimensionElement != nil; secondDimensionElement = secondDimensionElement.Next() {
			secondDimension[ secondDimensionIndex ] = secondDimensionElement.Value.(int)
			secondDimensionIndex++
		}
		convertedArray[ i ] = secondDimension
		i++
	}
	return convertedArray
}

func buildIngredients() *list.List {
	file, error := os.Open( "input.txt" )
	exitOnError( error )
	defer file.Close()
	scanner := bufio.NewScanner( file )
	
	nonAlphaNumericRegex := regexp.MustCompile( "[^a-zA-Z0-9- ]" )
	ingredients := list.New()
	
	for scanner.Scan() {
		alphaNumericInput := nonAlphaNumericRegex.ReplaceAllString( scanner.Text(), "" )
		input := strings.Split( alphaNumericInput, " " )
		
		capacity := stringToInteger( input[ 2 ] )
		durability := stringToInteger( input[ 4 ] )
		flavor := stringToInteger( input[ 6 ] )
		texture := stringToInteger( input[ 8 ] )
		
		ingredients.PushBack( buildList( capacity, durability, flavor, texture ) )
	}
	
	return ingredients
}

func buildPermutations( numberOfPossibilities int, numberOfAttributes int, ingredients [][]int ) *list.List {
	permutations := list.New()
	firstPermutation := list.New()
	for _, _ = range ingredients {
		firstPermutation.PushBack( MIN )
	}
	firstPermutation.Front().Value = numberOfPossibilities - ( numberOfAttributes - 1 )
	permutations.PushBack( firstPermutation )
	
	carry := 0
	for {
		permutation := copyList( permutations.Back().Value.(*list.List) )
		lastElement := listNth( permutation, len( ingredients ) - 1 )
		lastElementValue := lastElement.Value.(int)
		if lastElementValue > MIN {
			carry += lastElementValue - MIN
			lastElement.Value = MIN
		}
		for i := len( ingredients ) - 1; i >= 0; i-- {
			currentElement := listNth( permutation, i )
			if currentElement.Value.(int) > MIN {
				carry += MIN
				currentElement.Value = currentElement.Value.(int) - 1
				currentElement.Next().Value = currentElement.Next().Value.(int) + carry
				carry = 0
				break
			}
		}
		if permutation.Front().Value == MIN { break }
		permutations.PushBack( permutation )
	}
	
	return permutations
}

func findOptimalIngredientTotal( permutations []int, ingredients [][]int ) int {
	optimalPermutationValue := 0
	
	capacityIndex := 0
	durabilityIndex := 1
	flavorIndex := 2
	textureIndex := 3
	
	capacityTotal := 0
	durabilityTotal := 0
	flavorTotal := 0
	textureTotal := 0
	
	for i, ingredient := range ingredients {
		capacityTotal += permutations[ i ] * ingredient[ capacityIndex ]
		durabilityTotal += permutations[ i ] * ingredient[ durabilityIndex ]
		flavorTotal += permutations[ i ] * ingredient[ flavorIndex ]
		textureTotal += permutations[ i ] * ingredient[ textureIndex ]
	}
	
	if capacityTotal < 0 { capacityTotal = 0 }
	if durabilityTotal < 0 { durabilityTotal = 0 }
	if flavorTotal < 0 { flavorTotal = 0 }
	if textureTotal < 0 { textureTotal = 0 }
	
	permutationProduct := capacityTotal * durabilityTotal * flavorTotal * textureTotal
	
	if permutationProduct > optimalPermutationValue {
		optimalPermutationValue = permutationProduct
	}
	
	return optimalPermutationValue
}

func main() {
	ingredientsList := buildIngredients()
	ingredients := convert2DListTo2DArray( ingredientsList )
	permutationsList := buildPermutations( 100, len( ingredients ), ingredients )
	permutations := convert2DListTo2DArray( permutationsList )
	fmt.Println( permutations )
	//optimalPermutationValue := findOptimalIngredientTotal( permutations, ingredients )
	//fmt.Println( optimalPermutationValue )
}