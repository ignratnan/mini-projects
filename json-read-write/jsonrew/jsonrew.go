package jsonrew

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var employee []Employee
var folderpath string = "files"
var filename string = "employee.json"

type Employee struct {
	Name      string `json: "name"`
	Email     string `json: "email"`
	DateBirth string `json: "date_birth"`
	Phone     string `json: "phone"`
}

func writeJson(folderpath string, filename string, data Employee) error {
	fullpath := filepath.Join(folderpath, filename)
	fmt.Printf("Saving the data to %s file", filename)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to Marshal the data: %w", err)
	}

	err = os.WriteFile(fullpath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write the file: %w", err)
	}

	fmt.Println("The data written to %s successfully.", filename)
	return nil
}

func readJson(folderpath string, filename string) (Employee, error) {
	fullpath := filepath.Join(folderpath, filename)
	var loadEmployee Employee

	fmt.Printf("Reading the data from %s file", filename)

	jsonData, err := os.ReadFile(fullpath)
	if err != nil {
		return loadEmployee, fmt.Errorf("Failed to read the file: %w", err)
	}

	err = json.Unmarshal(jsonData, &loadEmployee)
	if err != nil {
		return loadEmployee, fmt.Errorf("Failed to Unmarshal the data: %w", err)
	}

	fmt.Println("The data loaded from %s successfully.", filename)
	return loadEmployee, nil
}

func inputData() {
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

	writeJson(folderpath, filename, employee_ref)

	employee = append(employee, employee_ref)
}

func mainMenu() string {
	reader := bufio.NewReader(os.Stdin)
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
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	return action
}

func inputMenu() {
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Input Menu")
	fmt.Println("---")
	inputData()
}

func showMenu() {
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Show Menu")
	fmt.Println("---")

}

func updateMenu() {
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Update Menu")
	fmt.Println("---")

}

func deleteMenu() {
	fmt.Println("Employee Management System")
	fmt.Println("---")
	fmt.Println("*Delete Menu")
	fmt.Println("---")

}

func Project() {
	action := mainMenu()
	switch action {
	case "1":
		inputMenu()
	case "2":
		showMenu()
	}
}
