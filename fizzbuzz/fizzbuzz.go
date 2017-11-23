// --- Directions
// Write a program that console logs the numbers
// from 1 to n. But for multiples of three print
// “fizz” instead of the number and for the multiples
// of five print “buzz”. For numbers which are multiples
// of both three and five print “fizzbuzz”.
// --- Example
//   fizzBuzz(5);
//   1
//   2
//   fizz
//   4
//   buzz

package fizzbuzz

func fizzBuzz(count int) []interface{} {
	var r []interface{}

	for i := 0; i < count; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			r = append(r, "fizzbuzz")
		} else if (i+1)%3 == 0 {
			r = append(r, "fizz")
		} else if (i+1)%5 == 0 {
			r = append(r, "buzz")
		} else {
			r = append(r, i+1)
		}
	}

	return r
}
