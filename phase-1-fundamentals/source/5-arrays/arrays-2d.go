package main

import "fmt"

func arrays2D(){
	temp:=1
	var v [5][5]int
	for i:=0; i<len(v); i++{
		for j:=0; j<len(v[0]); j++{
			v[i][j] = temp
			temp += 1
		}
	}

	fmt.Println("\nThe 2d array is :\n", v)
}