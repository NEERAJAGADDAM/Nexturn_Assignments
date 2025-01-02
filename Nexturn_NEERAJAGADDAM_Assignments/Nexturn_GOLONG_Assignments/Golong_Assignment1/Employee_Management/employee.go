package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Employee struct to hold employee details
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Constants for departments
const (
	HR      = "HR"
	IT      = "IT"
	Finance = "Finance"
)

var employees []Employee

// AddEmployee adds a new employee after validation
func AddEmployee(id int, name string, age int, department string) error {
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}
	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

// SearchEmployee searches for an employee by ID or name
func SearchEmployee(query string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprint(emp.ID) == query || strings.EqualFold(emp.Name, query) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// ListEmployeesByDepartment lists employees in a given department
func ListEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			result = append(result, emp)
		}
	}
	return result
}

// CountEmployees counts the number of employees in a given department
func CountEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

// ShowAllEmployees displays all employee details
func ShowAllEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees available.")
		return
	}
	fmt.Println("Existing Employees:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Department: %s\n", emp.ID, emp.Name, emp.Age, emp.Department)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add New Employee")
		fmt.Println("2. Search Employee by ID or by Name")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees in a Department")
		fmt.Println("5. Show All Employees")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input, please enter a number between 1 and 6.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter Employee ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID format.")
				continue
			}

			fmt.Print("Enter Employee Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Employee Age: ")
			ageStr, _ := reader.ReadString('\n')
			ageStr = strings.TrimSpace(ageStr)
			age, err := strconv.Atoi(ageStr)
			if err != nil {
				fmt.Println("Invalid Age format.")
				continue
			}

			fmt.Print("Enter Employee Department: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			if err := AddEmployee(id, name, age, department); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully.")
			}

		case 2:
			fmt.Print("Enter Employee ID or Name to search: ")
			query, _ := reader.ReadString('\n')
			query = strings.TrimSpace(query)

			if emp, err := SearchEmployee(query); err == nil {
				fmt.Println("Employee found:", *emp)
			} else {
				fmt.Println("Error:", err)
			}

		case 3:
			fmt.Print("Enter Department to list employees: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			employees := ListEmployeesByDepartment(department)
			if len(employees) == 0 {
				fmt.Println("No employees found in this department.")
			} else {
				fmt.Println("Employees in", department, "Department:")
				for _, emp := range employees {
					fmt.Println(emp)
				}
			}

		case 4:
			fmt.Print("Enter Department to count employees: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			count := CountEmployees(department)
			fmt.Printf("Number of employees in %s: %d\n", department, count)

		case 5:
			ShowAllEmployees()

		case 6:
			fmt.Println("Thankyou for your precious time!")
			return

		default:
			fmt.Println("Invalid choice, please select a valid option.")
		}
	}
}
