package main

import (
	"fmt"
)

const EmployeeWage = 500
const ContractorWage = 100
const FreelancerWage = 10


type Worker interface {
	calculateSalary() int
}

type Employee struct {
	days int
}
type Contractor struct {
	days int
}
type Freelancer struct {
	hours int
}

func (emp Employee) calculateSalary() int {
	return EmployeeWage *emp.days
}

func (contractor Contractor) calculateSalary() int {
	return ContractorWage *contractor.days
}
func (freelancer Freelancer) calculateSalary() int {
	return FreelancerWage *freelancer.hours
}
func calculate(worker Worker)  {
	fmt.Println(worker.calculateSalary())
}

func main() {
	emp := Employee{28}
	ctr := Contractor{28}
	frl := Freelancer{290}
	calculate(emp)
	calculate(ctr)
	calculate(frl)
}

