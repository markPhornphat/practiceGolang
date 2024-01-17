package main

import "fmt"

type Employee struct {
	employeeID  string
	name        string
	phoneNumber string
}

func main() {
	Employee1 := Employee{
		employeeID:  "101",
		name:        "Mark",
		phoneNumber: "09123456",
	}
	fmt.Println("employee1 = ", Employee1)

	//Array 
	employeeList := [3]Employee{} //To declare the size of array
	employeeList[0] = Employee{
		employeeID: "102",
		name: "Mo",
		phoneNumber: "11111111111",
	}
	employeeList[1] = Employee{
		employeeID: "103",
		name: "Tan",
		phoneNumber: "22222222222",
	}
	employeeList[2] = Employee{
		employeeID: "104",
		name: "Mathew",
		phoneNumber: "33333333333",
	}

	fmt.Println("EmployeeList = ", employeeList)


	//SLice
	employeeSlice := []Employee{}
	employeeSlice = append(employeeSlice, employeeList[0], employeeList[1], employeeList[2]);
	
	fmt.Println("EmployeeSlice = ",employeeSlice)
}