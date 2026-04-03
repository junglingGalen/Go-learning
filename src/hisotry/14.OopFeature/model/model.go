package model

import (
	"fmt"
)

type person struct {
	Name string
	age int
	salary float64
}

func NewPerson(name string, age int, salary float64) person {
	per := person{
		Name : name,
	}
	per.SetAge(age)
	per.SetSalary(salary)
	return per
}

func (p *person) SetAge(age int) {
	if age <= 130 && age >= 0 {
		(*p).age = age
	} else {
		fmt.Println("工资不在正常范围内")
	}
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水不在正常范围内")
	}
}

func (p * person) GetAge() int {
	return p.age
}

func (p *person) GetSalary() float64 {
	return p.salary
}
