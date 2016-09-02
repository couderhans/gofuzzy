package membershipfunctions

import (
	"gofuzzy/membershipfunctions/coordinatesystem"
	"strings"
)

type Trapezoid struct {
	LeftDegree0, LeftDegree1, RightDegree0, RightDegree1 coordinatesystem.Coordinate
}

func CreateTrapezoid(coordinates []string) Trapezoid {

	base1 := [2]string{(strings.Split(coordinates[0], ",")[0]), (strings.Split(coordinates[0], ",")[1])}
	top1 := [2]string{(strings.Split(coordinates[1], ",")[0]), (strings.Split(coordinates[1], ",")[1])}
	base2 := [2]string{(strings.Split(coordinates[2], ",")[0]), (strings.Split(coordinates[2], ",")[1])}
	top2 := [2]string{(strings.Split(coordinates[3], ",")[0]), (strings.Split(coordinates[3], ",")[1])}

	leftDegree0 := coordinatesystem.Coordinate{coordinatesystem.ConvertCoordinateToInt(base1[0]), coordinatesystem.ConvertCoordinateToInt(base1[1])}
	leftDegree1 := coordinatesystem.Coordinate{coordinatesystem.ConvertCoordinateToInt(top1[0]), coordinatesystem.ConvertCoordinateToInt(top1[1])}
	rightDegree0 := coordinatesystem.Coordinate{coordinatesystem.ConvertCoordinateToInt(base2[0]), coordinatesystem.ConvertCoordinateToInt(base2[1])}
	rightDegree1 := coordinatesystem.Coordinate{coordinatesystem.ConvertCoordinateToInt(top2[0]), coordinatesystem.ConvertCoordinateToInt(top2[1])}

	trapezoid := Trapezoid{leftDegree0, leftDegree1, rightDegree0, rightDegree1}

	return trapezoid
}
