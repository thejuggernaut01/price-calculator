package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error)  {
	var floats []float64

	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			fmt.Println("Converting price to float failed!")
			fmt.Println(err)
			return nil, errors.New("Failed to convert string to float.")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}