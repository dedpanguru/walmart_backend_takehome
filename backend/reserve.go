package backend

import (
	"errors"
	"strconv"
	"strings"
)

var theater = make(map[rune][]bool) // represents the theater seating
var reserved = 0                    // tracks number of reservations

func init() {
	for key := 'B'; key <= 'J'; key += 2 {
		theater[key] = []bool{false, false, false, false, false}
	}
	/* Theater representation structure
	A buffer (also no one wants to be close to the screen)
	B F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer
	C buffer
	D F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer
	E buffer
	F F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer
	G buffer
	H F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer
	I buffer
	J F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer F buffer buffer buffer
	*/
}

func Reserve(request int) (string, error) {
	// 3 seat buffer between reserved seats
	// 1 row buffer
	// for every non-buffer row, allocate seats 1,5,9,13,17 in rows B,D,F,H,J = 25 seats total -> customer safety
	var bldr strings.Builder // will build a request's output string
	if request+reserved > 25 {
		return "", errors.New("insufficient space in theater to support most recent request")
	}
	reserved += request

	for row := 'J'; row >= 'B'; row -= 2 {
		if theater[row][4] { // skip row if row is reserved entirely already
			continue
		}
		seatNum := 0
		for seatNum < 5 && request > 0 {
			if !theater[row][seatNum] { // if a seat isn't reserved
				theater[row][seatNum] = true //reserve it by marking it as true in theater
				// translate into output
				bldr.WriteRune(row)
				switch seatNum {
				case 0:
					bldr.WriteString(strconv.Itoa(1))
				case 1:
					bldr.WriteString(strconv.Itoa(5))
				case 2:
					bldr.WriteString(strconv.Itoa(9))
				case 3:
					bldr.WriteString(strconv.Itoa(13))
				case 4:
					bldr.WriteString(strconv.Itoa(17))
				}
				bldr.WriteRune(' ')
				request--
			}
			seatNum++
		}
	}
	return bldr.String(), nil
}
