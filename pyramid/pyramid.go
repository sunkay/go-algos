// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a pyramid shape
// with N levels using the # character.  Make sure the
// pyramid has spaces on both the left *and* right hand sides
// --- Examples
//   pyramid(1)
//       '#'
//   pyramid(2)
//       ' # '
//       '###'
//   pyramid(3)
//       '  #  '
//       ' ### '
//       '#####'

package pyramid

import (
	"math"
)

func pyramid(count int) []string {
	var rows []string
	midf := math.Floor(float64((count*2 - 1) / 2))
	mid := int(midf)
	for row := 0; row < count; row++ {
		cols := ""
		for col := 0; col < (count*2 - 1); col++ {
			if mid-row <= col && mid+row >= col {
				cols += "#"
			} else {
				cols += " "
			}
		}
		rows = append(rows, cols)
	}

	return rows
}
