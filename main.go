package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EmployeeData struct {
	fname  string
	lname  string
	salary float64
}

func main() {
	file, err := os.Open("salary.txt")
	if err != nil {
		fmt.Println("Could not load file!", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var employees []EmployeeData
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 3 {
			fmt.Println("Invalid syntax", line)
			continue
		}
		fname := parts[0]
		lname := parts[1]
		salary := parts[2]
		var salaryValue float64
		fmt.Sscanf(salary, "%f", &salaryValue)
		employeeData := EmployeeData{fname: fname, lname: lname, salary: salaryValue}
		employees = append(employees, employeeData)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	file1, err := os.Create("sameSalary.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	line := fmt.Sprintf("%-10s %-10s %-10s %-10s %-10s\n", "First Name", "Last Name", "First Name", "Last Name", "Salary")
	_, err = file1.WriteString(line)
	if err != nil {
		fmt.Println("error while writing", err)
		return
	}

	for i, firstSalary := range employees {
		for j := i + 1; j < len(employees); j++ {
			secondSalary := employees[j]
			if firstSalary.salary == secondSalary.salary {
				fmt.Printf(" %s\t %s\t %s\t %s\t %.2f\n", firstSalary.fname, firstSalary.lname, secondSalary.fname, secondSalary.lname, secondSalary.salary)
				line := fmt.Sprintf("%-10s %-10s %-10s %-10s %-10.2f\n", firstSalary.fname, firstSalary.lname, secondSalary.fname, secondSalary.lname, secondSalary.salary)
				_, err := file1.WriteString(line)
				if err != nil {
					fmt.Println("error while writing", err)
					return
				}
			}
		}
	}
	/*for _, employee := range employees {
		fmt.Printf("Name: %s, Salary: %.2f\n", employee.name, employee.salary)
	}*/
}
