package fuzzy

import ( 
	"fmt" 
	"strings"
	"strconv"
)

type Coordinate struct {
	X,Y int
}

type Trapezoid struct {
	LeftDegree0, LeftDegree1, RightDegree0, RightDegree1 Coordinate
}

type FuzzySet struct {
	Name string
    Trapezoid Trapezoid
}

func CreateFuzzySet(line string) FuzzySet {
    
	name := strings.Split(line, "||")[0]
	set := strings.Split(line, "||")[1]
	//result := strings.Split(line, "||")[2]
	
	coordinates := strings.Split(set, ";")

	trapezoid := CreateTrapezoid(coordinates)
	
	fuzzyset := FuzzySet{name,  trapezoid}
	return fuzzyset
}

func CreateTrapezoid(coordinates []string) Trapezoid {
	
	base1 := [2]string{(strings.Split(coordinates[0], ",")[0]), (strings.Split(coordinates[0], ",")[1])}
	top1 := [2]string{(strings.Split(coordinates[1], ",")[0]), (strings.Split(coordinates[1], ",")[1])}
	base2 := [2]string{(strings.Split(coordinates[2], ",")[0]), (strings.Split(coordinates[2], ",")[1])}
	top2 := [2]string{(strings.Split(coordinates[3], ",")[0]), (strings.Split(coordinates[3], ",")[1])}
	
	leftDegree0 := Coordinate{ConvertCoordinateToInt(base1[0]), ConvertCoordinateToInt(base1[1])}
    leftDegree1 := Coordinate{ConvertCoordinateToInt(top1[0]), ConvertCoordinateToInt(top1[1])}
    rightDegree0 := Coordinate{ConvertCoordinateToInt(base2[0]), ConvertCoordinateToInt(base2[1])}
    rightDegree1 := Coordinate{ConvertCoordinateToInt(top2[0]), ConvertCoordinateToInt(top2[1])}
	
	trapezoid := Trapezoid { leftDegree0, leftDegree1, rightDegree0, rightDegree1 }
	
	return trapezoid
}

func ConvertCoordinateToInt(coordinate string) int {
    
	result, err := strconv.Atoi(coordinate)
	fmt.Println("Converting coordinate to Int: %s - Error: %s",result,err) 
	return result
}