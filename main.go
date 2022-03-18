package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Gurv33r/GolangGuru/theater"
)

func main() {
	/* a, b, start := []int{0, 0, 2, 100}, []int{1, 10, 3, 77}, time.Now()
	fmt.Println("a =", a, "b =", b)
	bf := SumOfTwoBruteForce(a, b, 12)
	bfduration := time.Since(start)
	op := SumOfTwoOptimized(a, b, 12)
	opduration := time.Since(start).Nanoseconds()
	fmt.Println("Brute force => Expected: true, Actual:", bf, "Time =", bfduration)
	fmt.Println("Optimized => Expected: true, Actual:", op, "Time =", opduration) */
	//numOfAllRunes()
	/* arr := [3][3]int{
		[3]int{0,1,2},
		[3]int{3,4,5},
		[3]int{6,7,8},
	} */
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("input file missing from command-line arguments")
	}
	inputFileName := args[0]
	inFile, err := os.Open(inputFileName)
	defer inFile.Close()
	handle(err)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		request := strings.Fields(scanner.Text())
		numSeatsRequested, err := strconv.Atoi(request[1])
		handle(err)
		reservation, err := theater.Reserve(numSeatsRequested)
		handle(err)
		fmt.Println(request[0], reservation)
	}
}

func SumOfTwoBruteForce(a, b []int, v int) bool {
	for _, i := range a {
		complement := v - i
		for _, j := range b {
			if j == complement {
				return true
			}
		}
	}
	return false
}

func SumOfTwoOptimized(a, b []int, v int) bool {
	complementMap := make(map[int]int)
	for i := 0; i < len(a)+len(b); i++ {
		if i < len(a) {
			complementMap[a[i]] = v - a[i]
		} else {
			return contains(complementMap, b[i-len(a)])
		}
	}
	return false
}

func contains(m map[int]int, target int) bool {
	for _, v := range m {
		if v == target {
			return true
		}
	}
	return false
}

func numOfAllRunes() {
	var c bool
	fmt.Println(c)
	for curr := 'a'; curr <= 'z'; curr++ {
		fmt.Println(curr)
	}
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

/*func Rotate90ClockWise(a [3][3]int) [3][3]int {

		start = 012
				345
				678
		end = 	630
				741
				258


}*/

/*func transpose(a [3][3]int) [3][3]int {

} */
