package jsonrew

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var employee []Employee
var employee_input []Employee
var employee_update []Employee
var folderpath string = "json-files"
var filename string = "employee.json"

type Employee struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	DateBirth string `json:"date_birth"`
	Phone     string `json:"phone"`
}

func writeJson(folderpath string, filename string, data interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal the data: %w", err)
	}

	err = os.WriteFile(fullpath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write the file: %w", err)
	}

	return nil
}

func readJson(folderpath string, filename string, target interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := os.ReadFile(fullpath)
	if err != nil {
		return fmt.Errorf("failed to read the file: %w", err)
	}

	err = json.Unmarshal(jsonData, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the data: %w", err)
	}

	return nil
}

func numInput(text string) int {
	num_input := ""
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(text)
	num_input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	num_input = strings.TrimSpace(num_input)
	num_int, err := strconv.Atoi(num_input)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return num_int
}

func inputData() Employee {
	name := ""
	email := ""
	date_birth := ""
	phone := ""

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input the employee name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Please input the employee email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Please input the employee date birth (12 August 1995): ")
	date_birth, _ = reader.ReadString('\n')
	date_birth = strings.TrimSpace(date_birth)

	fmt.Print("Please input the employee phone number: ")
	phone, _ = reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	employee_ref := Employee{
		Name:      name,
		Email:     email,
		DateBirth: date_birth,
		Phone:     phone,
	}

	return employee_ref
}

func writeData(employee_ref Employee) {
	employee_input = append(employee_input, employee_ref)

	employee = employee_input

	err := writeJson(folderpath, filename, employee)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func showData() {
	fmt.Printf("The employee data from '%s' file:\n", filename)
	number := 1
	for _, emp := range employee {
		fmt.Printf("%d.\tEmployee Name: %s, Email: %s, Date of Birth: %s, Phone: %s\n", number, emp.Name, emp.Email, emp.DateBirth, emp.Phone)
		number++
	}
}

func updateData(num_int int, emp_update Employee) []Employee {
	employee_update = nil
	number := 1
	readJson(folderpath, filename, &employee)
	for _, emp := range employee {
		if number == num_int {
			employee_update = append(employee_update, emp_update)
		} else {
			employee_update = append(employee_update, emp)
		}
		number++
	}

	employee = employee_update
	return employee
}

func deleteData(num_int int) []Employee {
	employee_update = nil
	number := 1
	readJson(folderpath, filename, &employee)
	for _, emp := range employee {
		if number != num_int {
			employee_update = append(employee_update, emp)
		}
		number++
	}

	employee = employee_update
	return employee
}

func mainMenu() {
	action := ""

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("######################################################################")
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Main Menu")
	fmt.Println("---")
	fmt.Println("Please select the option below:")
	fmt.Println("1. Input Employee Data")
	fmt.Println("2. Show Employee Data")
	fmt.Println("3. Update Employee Data")
	fmt.Println("4. Delete Employee")
	fmt.Println("5. Exit")
	fmt.Println("---")
	fmt.Print("Please input your option: ")
	action, _ = reader.ReadString('\n')
	action = strings.TrimSpace(action)

	switch action {
	case "1":
		inputMenu()
	case "2":
		showMenu()
	case "3":
		updateMenu()
	case "4":
		deleteMenu()
	case "5":
		fmt.Println("Exiting the program")
	}
}

func inputMenu() {
	fmt.Println("######################################################################")
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Input Menu")
	fmt.Println("---")
	readJson(folderpath, filename, &employee_input)
	emp_input := inputData()
	writeData(emp_input)
	mainMenu()
}

func showMenu() {
	fmt.Println("######################################################################")
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Show Menu")
	fmt.Println("---")
	readJson(folderpath, filename, &employee)
	if employee != nil {
		showData()
	} else {
		fmt.Println("No data available")
	}
	fmt.Println("---")
	fmt.Print("Click enter to back to Main Menu")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	mainMenu()
}

func updateMenu() {
	fmt.Println("######################################################################")
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Update Menu")
	fmt.Println("---")
	readJson(folderpath, filename, &employee)
	if employee != nil {
		showData()
		fmt.Println("---")
		num_input := numInput("Please input the number to update: ")
		emp_up := inputData()
		new_emp := updateData(num_input, emp_up)
		writeJson(folderpath, filename, new_emp)
		fmt.Println("Data successfully updated")
	} else {
		fmt.Println("No data available")
	}
	fmt.Println("---")
	fmt.Print("Click enter to back to Main Menu")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	mainMenu()
}

func deleteMenu() {
	fmt.Println("######################################################################")
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Delete Menu")
	fmt.Println("---")
	readJson(folderpath, filename, &employee)
	if employee != nil {
		showData()
		fmt.Println("---")
		num_input := numInput("Please input the number to update: ")
		new_emp := deleteData(num_input)
		writeJson(folderpath, filename, new_emp)
		fmt.Println("Data successfully deleted")
	} else {
		fmt.Println("No data available")
	}
	fmt.Println("---")
	fmt.Print("Click enter to back to Main Menu")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	mainMenu()
}

func Project() {
	os.MkdirAll(folderpath, 0755)
	mainMenu()
}
