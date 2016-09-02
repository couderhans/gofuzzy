package logic

import (
	"gofuzzy/membershipfunctions"
	"strings"
)

type FuzzySet struct {
	Name      string
	Trapezoid membershipfunctions.Trapezoid
}

func CreateFuzzySet(line string) FuzzySet {

	name := strings.Split(line, "||")[0]
	set := strings.Split(line, "||")[1]
	//result := strings.Split(line, "||")[2]

	coordinates := strings.Split(set, ";")

	trapezoid := membershipfunctions.CreateTrapezoid(coordinates)

	fuzzyset := FuzzySet{name, trapezoid}
	return fuzzyset
}
