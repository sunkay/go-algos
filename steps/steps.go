// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a step shape
// with N levels using the # character.  Make sure the
// step has spaces on the right hand side!
// --- Examples
//   steps(2)
//     0  '# '
//		   01
//     1  '##'
//	       01
//   steps(3)
//     0  '#  '
//         012
//     1  '## '
//         012
//     2  '###'
//         012
//   steps(4)
//       '#   '
//       '##  '
//       '### '
//       '####'

package steps

import (
	"fmt"
)

func steps(count int) []string {
	var steps []string

	for i := 0; i < count; i++ {
		step := ""
		for j := 0; j < count; j++ {
			if i >= j {
				step += "#"
			} else {
				step += " "
			}
		}
		fmt.Println(step)
		steps = append(steps, step)
	}

	return steps
}
