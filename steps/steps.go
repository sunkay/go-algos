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

func stepsIter(count int) []string {
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
		steps = append(steps, step)
	}

	return steps
}

func stepsRecursive(count int, currentStep int, steps []string) []string {

	if currentStep == count {
		return steps
	}

	str := ""
	for i := 0; i < count; i++ {
		if currentStep >= i {
			str += "#"
		} else {
			str += " "
		}
	}
	//fmt.Fprintf(os.Stderr, "count=%d, currentStep:%d, str:%s\n", count, currentStep, str)

	return stepsRecursive(count, currentStep+1, append(steps, str))

}
