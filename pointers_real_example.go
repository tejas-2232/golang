package main

import "fmt"

type Employee struct { // custom data type
	empName  string
	position string
	salary   float64
}

func update_salary(emp *Employee, newSalary float64) {
	fmt.Println(emp)
	emp.salary = newSalary
}

// data type of emp is *Employee, emp is a poinetr ref variable-

func main() {

	emp := Employee{
		empName:  "Harry Potter",
		position: "software developer",
		salary:   500000,
	}

	fmt.Println("\n initial Salary:", emp)
	fmt.Println("\n value of emp=")
	update_salary(&emp, 600000) // fetch mem add of emp variable. passing ref and also updating  value at source also

	fmt.Println("\nAfter update Salary:", emp)

}
