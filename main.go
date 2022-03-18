package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"theater/backend"
)

func main() {
	// grab input file and output file names
	cmdArgs, inFileName := os.Args, ""
	if len(cmdArgs) < 2 {
		panic(errors.New("missing input file"))
	} else if len(cmdArgs) > 3 {
		panic(errors.New("too many command-line arguments"))
	}
	if len(cmdArgs) == 2 {
		inFileName = cmdArgs[1]
	}
	// read request file
	inputFile, err := os.Open("./" + inFileName)
	handle(err)
	defer handle(inputFile.Close())
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// begin reading input
		if len(scanner.Text()) == 0 {
			break
		}
		request := strings.Split(scanner.Text(), " ")
		seatNumsRequested, err := strconv.Atoi(request[1])
		handle(err)
		reservations, err := backend.Reserve(seatNumsRequested) // handles seat allocation
		handle(err)

		// write to output file
		//outFile, err := os.Open(outFileName)
		//handle(err)
		//writer := bufio.NewWriter(outFile)
		//_, err = writer.WriteString(request[0] + " " + reservations + "\n")
		//handle(err)
		//handle(outFile.Close())
		fmt.Println(request[0] + " " + reservations)
	}
}
func handle(err error) {
	if err != nil {
		panic(err)
	}
}
