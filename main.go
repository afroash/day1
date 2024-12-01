package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AOC OF CODE BY Ash")
	data, err := readFile("actual-input")
	if err != nil {
		log.Fatal("We could not read the file due to error:", err)
	}
	col1, col2 := sortData(data)
	fmt.Println("The sum of the difference between the two columns is:", findDifference(col1, col2))
	fmt.Println("The similarity score between the two columns is:", similarityScore(col1, col2))

}

func readFile(filepath string) ([][]int, error) {

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("We could not load the file due to error:", err)
	}
	defer f.Close()

	var data [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) != 2 {
			continue
		}
		n1, err1 := strconv.Atoi(line[0])
		n2, err2 := strconv.Atoi(line[1])
		if err1 != nil || err2 != nil {
			continue
		}
		data = append(data, []int{n1, n2})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("We could not scan the file due to error:", err)
	}
	return data, nil
}

// sort the data
func sortData(data [][]int) ([]int, []int) {
	var col1, col2 []int
	for _, v := range data {
		col1 = append(col1, v[0])
		col2 = append(col2, v[1])
	}
	sort.Ints(col1)
	sort.Ints(col2)
	return col1, col2
}

// find difference between two slices and return the sum
func findDifference(col1, col2 []int) int {
	var sum int
	for i := 0; i < len(col1); i++ {
		dif := math.Abs(float64(col2[i]) - float64(col1[i]))
		sum += int(dif)
		//fmt.Println(unsdiff)
		//sum += uint(col2[i]) - uint(col1[i])
	}
	return sum
}

// PART 2
// similarity score. check value in col1 and count occurance in col2. mulitply the occurance by the value in col1 and sum all the values
func similarityScore(col1, col2 []int) int {
	var sum int
	for i := 0; i < len(col1); i++ {
		count := 0
		for j := 0; j < len(col2); j++ {
			if col1[i] == col2[j] {
				count++
			}
		}
		sum += count * col1[i]
	}
	return sum
}
