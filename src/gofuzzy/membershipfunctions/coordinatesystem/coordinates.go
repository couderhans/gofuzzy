package coordinatesystem

import (
	"fmt"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

func ConvertCoordinateToInt(coordinate string) int {

	result, err := strconv.Atoi(coordinate)
	fmt.Println("Converting coordinate to Int: %s - Error: %s", result, err)
	return result
}
