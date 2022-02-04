package finder

import "fmt"

var ErrMalformedInput = fmt.Errorf("malformed input")

func FindElement(els []int) (int, error) {
	if len(els) != 100 {
		return 0, ErrMalformedInput
	}
	var sum, sumIdeal int
	for idx, el := range els {
		sumIdeal += idx + 1		// final value could be const
		sum += el
	}
	return sumIdeal - sum, nil
}
