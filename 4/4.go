package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
)

func main() {
	input := "iwrupvqb"
	counter := 1
	for {
		md5Sum := md5.Sum( []byte( input + strconv.Itoa( counter ) ) )
		if md5Sum[ 0 ] == 0 && md5Sum[ 1 ] == 0 && md5Sum[ 2 ] < 16 {
			fmt.Println( "Winner!!" )
			fmt.Println( counter )
			break
		}
		counter++
	}
}