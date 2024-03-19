package conversion

import (
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringValue := range strings {
		convertedValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, convertedValue)
	}
	return floats, nil
}
