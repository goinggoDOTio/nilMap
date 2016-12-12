package main

import "fmt"

func main() {
	//Initializing an nil map - DONT DO THIS!
	//var myMap map[string]int
	//myMap["hello"] = 1
	//fmt.Println(myMap)
	//panic: assignment to entry in nil map
	//Because there is no way to get out of a nil in a map. Map there is no append like the slice

	//Initializing a map the right, or one of the right ways
	myRightMap := make(map[string]int)
	myRightMap["helloAgain"] = 1
	fmt.Println(myRightMap)
	//map[helloAgain:1]
}
