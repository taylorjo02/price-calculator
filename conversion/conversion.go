package conversion

import (
	"errors"
	"strconv"
)

var ErrStringConversionFailed = errors.New("failed to convert string to float")

func StringsToFloats(strings []string) ([]float64, error) {

	var floats []float64
	for _, string := range strings {
		floatVal, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, ErrStringConversionFailed
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}