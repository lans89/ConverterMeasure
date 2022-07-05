package distances

import "fmt"

func ConvertDistance(unitIn string, unitOut string, number float64) float64 {
	_, err1 := unitMap[unitIn]
	_, err2 := unitMap[unitOut]
	if !err1 || !err2 {
		fmt.Println("Error: UnitIn [", unitIn, "] or UnitOut [", unitOut, "] not found!")
		return -1
	}

	key := unitIn + flow + unitOut
	res, err := convertionMap[key]
	if !err {
		fmt.Println("Error: Key [", key, "] not found!")
		return -1
	}

	switch res.Operation {
	case "*":
		return number * res.Value
	default:
		return number / res.Value
	}
}
