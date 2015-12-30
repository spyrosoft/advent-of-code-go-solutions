package main

import (
	"fmt"
	"time"
)

func main() {
	MIN_PRESENTS := 36000000
	door := 1
	totalPresents := 0
	startBenchmark := time.Now().Unix()
	for {
		totalPresents = 0
		for elf := door; elf > 0; elf-- {
			if door % elf == 0 {
				totalPresents += 10 * elf
			}
		}
		if MIN_PRESENTS <= totalPresents { break }
		door++
	}
	fmt.Println( "Door:", door, "Total Presents:", totalPresents )
	stopBenchmark := time.Now().Unix()
	fmt.Println( "Total Seconds to Calculate:", stopBenchmark - startBenchmark )
	//This took 3144 seconds for me - I was going to sleep anyway
	//This could be optimized using go routines and a better algorithm
}