//Ying Vang
//February 8, 2015
//Branch Testing
//In this push, I trying to push a branch to github.com

package main

import (
	"fmt"
)

type student struct {
	name       string
	testScore1 int
	testScore2 int
	average    int
}

//Functions
func avg(t1, t2 int) int {
	var x int = (t1 + t2) / 2
	return x

}

func modifyAvg(t *student) {
	*t = student{average: 99, name: t.name}
}

//Multiple returns
func SumProductDiff(i, j int) (int, int, int) {
    return i+j, i*j, i-j
}

const (
	pie = 3.14
)

func main() {
	var csClass [3]student
	csClass[0].name = "Mat Jones"
	csClass[0].testScore1, csClass[0].testScore2 = 77, 93
	csClass[1].name = "Mike Lee"
	csClass[1].testScore1, csClass[1].testScore2 = 78, 94
	csClass[2].name = "John Doe"
	csClass[2].testScore1, csClass[2].testScore2 = 86, 89

	for i := 0; i < 3; i++ {
		csClass[i].average = avg(csClass[i].testScore1, csClass[i].testScore2)
		fmt.Printf("%s, average test score is %d percent.\n", csClass[i].name, csClass[i].average)
	}

	modifyAvg(&csClass[0])
	fmt.Printf("now, %s, average test score is %d percent.\n", csClass[0].name, csClass[0].average)
	
	sum, prod, diff := SumProductDiff(3,4)
    fmt.Println("Sum:", sum, "| Product:",prod, "| Diff:", diff)
    
    var r float64
	fmt.Printf("Enter the radius? ")
	fmt.Scanf("%f", &r) //geting input

	var answer float64 = (pie) * (r * r)

	fmt.Printf("The radius of the circle is: %f", answer)

}
