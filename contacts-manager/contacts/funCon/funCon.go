package funCon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
}

var contact []Contact
var valMain bool

func MainMenu() {
	valMain = false
	for !valMain {
		fmt.Println("==============================")
		fmt.Println("-------CONTACTS MANAGER-------")
		fmt.Println("==============================")
		fmt.Println("---")
		fmt.Println("*Main menu")
		fmt.Println("---")
		fmt.Println("1. Show Contacts")
		fmt.Println("2. Add Contacts")
		fmt.Println("3. Edit Contacts")
		fmt.Println("4. Delete Contacts")
		fmt.Println("5. Exit")
		fmt.Println("---")

		val := false
		ret := ""
		for !val {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Please input the number: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading line:", err)
				return
			}
			input = strings.TrimSpace(input)
			switch input {
			case "1", "2", "3", "4", "5":
				ret = input
				val = true
			default:
				fmt.Println("Invalid input!")
			}
		}

		switch ret {
		case "1":
			ShowCon()
		case "2":
			AddCon()
		case "3":
			EditCon()
		case "4":
			DeleteCon()
		case "5":
			valMain = true
		}
	}
}

func ShowCon() {
	fmt.Println("==============================")
	fmt.Println("-------CONTACTS MANAGER-------")
	fmt.Println("==============================")
	fmt.Println("---")
	fmt.Println("*Show Contacts")
	fmt.Println("---")
	num := 0
	fmt.Println("Contacts:")
	for i := 0; i < len(contact); i++ {
		num = i + 1
		fmt.Print(num, ". ", contact[i].Name, ", ", contact[i].Phone, "\n")
	}

	valShow := false
	for !valShow {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input 'back' to back to Main Menu: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)
		switch input {
		case "back":
			valShow = true
		default:
			fmt.Println("Invalid input!")
		}
	}
	MainMenu()
}

func AddCon() {
	valAdd := false
	for !valAdd {
		fmt.Println("==============================")
		fmt.Println("-------CONTACTS MANAGER-------")
		fmt.Println("==============================")
		fmt.Println("---")
		fmt.Println("*Add Contacts")
		fmt.Println("---")
		num := 0
		fmt.Println("Contacts available:")
		for i := 0; i < len(contact); i++ {
			num = i + 1
			fmt.Print(num, ". ", contact[i].Name, ", ", contact[i].Phone, "\n")
		}

		var nameTemp string
		var phoneTemp string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Contact name: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)
		nameTemp = input

		if nameTemp == "back" {
			valAdd = true
			MainMenu()
		}

		fmt.Print("Phone number: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)
		phoneTemp = input
		if phoneTemp == "back" {
			valAdd = true
			MainMenu()
		}

		nameCon := Contact{nameTemp, phoneTemp}
		contact = append(contact, nameCon)
	}
}

func EditCon() {
	valEdit := false
	for !valEdit {
		fmt.Println("==============================")
		fmt.Println("-------CONTACTS MANAGER-------")
		fmt.Println("==============================")
		fmt.Println("---")
		fmt.Println("*Edit Contacts")
		fmt.Println("---")
		num := 0
		fmt.Println("Contacts available:")
		for i := 0; i < len(contact); i++ {
			num = i + 1
			fmt.Print(num, ". ", contact[i].Name, ", ", contact[i].Phone, "\n")
		}

		var nameTemp string
		var phoneTemp string
		var upName *string
		var upPhone *string

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input the number of contact: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)

		if input == "back" {
			valEdit = true
			MainMenu()
		}

		inputNum, errNum := strconv.ParseInt(input, 10, 64)
		if errNum != nil {
			fmt.Println("Error reading line:", errNum)
			return
		}

		fmt.Print("Contact name: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)

		if input == "back" {
			valEdit = true
			MainMenu()
		}

		nameTemp = input

		fmt.Print("Phone number: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)

		if input == "back" {
			valEdit = true
			MainMenu()
		}

		phoneTemp = input

		ind := inputNum - 1
		upName = &contact[ind].Name
		upPhone = &contact[ind].Phone

		*upName = nameTemp
		*upPhone = phoneTemp
	}
}

func DeleteCon() {
	valDelete := false
	for !valDelete {
		fmt.Println("==============================")
		fmt.Println("-------CONTACTS MANAGER-------")
		fmt.Println("==============================")
		fmt.Println("---")
		fmt.Println("*Delete Contacts")
		fmt.Println("---")
		num := 0
		fmt.Println("Contacts available:")
		for i := 0; i < len(contact); i++ {
			num = i + 1
			fmt.Print(num, ". ", contact[i].Name, ", ", contact[i].Phone, "\n")
		}

		var contactTemp []Contact

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please input the number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
			return
		}
		input = strings.TrimSpace(input)

		if input == "back" {
			valDelete = true
			MainMenu()
		}

		inputNum, errNum := strconv.ParseInt(input, 10, 64)
		if errNum != nil {
			fmt.Println("Error reading line:", errNum)
			return
		}
		ind := inputNum - 1
		for i := range contact {
			if contact[i] == contact[ind] {
				continue
			}
			contactTemp = append(contactTemp, contact[i])
		}
		contact = contactTemp
	}
}
