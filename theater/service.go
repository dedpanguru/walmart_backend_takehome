package theater

import (
	"errors"
	"strconv"
	"strings"
)

var reserved = 0
var theater = make(map[rune][]bool)

func init() {
	for row := 'B'; row <= 'J'; row++ {
		theater[row] = []bool{false, false, false, false}
	}
}

func Reserve(numSeatsRequested int) (string, error) {
	if reserved+numSeatsRequested > 25 {
		return "", errors.New("insufficient space in theater to support request")
	}
	var bldr strings.Builder
	reserved += numSeatsRequested
	for row := 'J'; row >= 'B'; row-- {
		if theater[row][1] { // first seat is the last to be allocated, so if it is, the row is full and should be skipped
			continue
		}
		seatNum := 0
		for seatNum < 5 && numSeatsRequested > 0 {
			if !theater[row][seatNum] {
				theater[row][seatNum] = true
				bldr.WriteString(string(row) + strconv.Itoa(seatNum*4+1) + " ")
				seatNum++
				numSeatsRequested--
			}
		}
		if numSeatsRequested == 0 {
			break
		}
	}
	return bldr.String(), nil
}
