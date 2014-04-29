package fuzzy

import (
	"fmt"
	"bytes"
	"strings"
)


func GetFuzzySets() []FuzzySet {

    var fuzzysets []FuzzySet
       
	file := ReadFile("tr1.fsd")
	lines := strings.Split(file, "\n")
	for line := range lines {

	}  	
    return fuzzysets
}


func fuzzySetToString(fuzzyset FuzzySet) string {
	
	leftdegree0 := fmt.Sprintf("%v,%v",fuzzyset.Trapezoid.LeftDegree0.X,fuzzyset.Trapezoid.LeftDegree0.Y)
	leftdegree1 := fmt.Sprintf("%v,%v",fuzzyset.Trapezoid.LeftDegree1.X,fuzzyset.Trapezoid.LeftDegree1.Y)
	rightdegree0 := fmt.Sprintf("%v,%v",fuzzyset.Trapezoid.RightDegree0.X,fuzzyset.Trapezoid.RightDegree0.Y)
	rightdegree1 := fmt.Sprintf("%v,%v",fuzzyset.Trapezoid.RightDegree1.X,fuzzyset.Trapezoid.RightDegree1.Y)
				
	trapezoid := fmt.Sprintf("%v;%v;%v;%v",leftdegree0,leftdegree1,rightdegree0,rightdegree1)
	
	output := fmt.Sprintf("%v||%v||",fuzzyset.Name,trapezoid)
	
	return output
	
}