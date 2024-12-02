package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {

	leftList := []int{}
	rightList := []int{}

	totalDistance := 0
	similarityScore := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error Opening Input File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	selectList := true

    scanner.Split(bufio.ScanWords)   // use scanwords

    for scanner.Scan() {
    	input, err := strconv.Atoi(scanner.Text())
    	if err != nil {
    		fmt.Println("Conversion Error")
    	}
        if selectList {
        	leftList = append(leftList, input)
        	selectList = false
        	//fmt.Println("Left: ", input)
        } else {
        	rightList = append(rightList, input)
        	selectList = true
        	//fmt.Println("Right: ", input)
        }
    }

    sort.Ints(leftList)
    sort.Ints(rightList)

    for i, _ := range leftList {
    	if leftList[i] <= rightList[i] {
    		totalDistance += rightList[i] - leftList[i]
    	} else {
    		totalDistance += leftList[i] - rightList[i]
    	}
    }

    for i, xl := range leftList {
    	counter := 0
    	for _, xr := range rightList {
    		if xl == xr {
    			counter++
    		}
    	}
    	similarityScore += leftList[i] * counter
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    fmt.Println("Total Distance: ", totalDistance)
    fmt.Println("Similarity Score: ", similarityScore)
}