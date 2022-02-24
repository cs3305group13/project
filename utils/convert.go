package utils

import "strconv"

func ConvertToFloat(number string) float64 {
	float, err := strconv.ParseFloat(number, 64)
	CheckError(err)
	
	return	float
}
