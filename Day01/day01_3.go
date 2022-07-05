// Third Question

package main

import (
	"fmt"
)

type Emp interface {
	setsalary(int) int
}

type Employee struct {
	jobtype string
}

func (e *Employee) setsalary(t int) int {
	if e.jobtype == "FullTime" {
		return 500 * t
	}
	if e.jobtype == "Contractor" {
		return 100 * t
	}
	if e.jobtype == "Freelancer" {
		return 10 * t
	}
	return 0
}

func main() {
	var emp1, emp2, emp3 Emp
	//var emp1 Employee

	emp1 = &Employee{jobtype: "FullTime"}
	emp2 = &Employee{jobtype: "Contractor"}
	emp3 = &Employee{jobtype: "Freelancer"}

	fmt.Println(emp1.setsalary(28))
	fmt.Println(emp2.setsalary(28))
	fmt.Println(emp3.setsalary(280))

}