package todolist

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// variable declaration
var task_list []string
var task_status []string
var temp_list []string
var temp_status []string

var main_input_list = []string{"1", "2", "3", "4", "5"}
var show_status_list = []string{"1", "2", "3"}

var opin int
var show string
var task string
var edit_num_raw string
var task_num_raw string

func containsInt(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true // Found it!
		} else {

		}
	}
	return false // Not found
}

func main_view() {
	for opin != 5 {
		//Set the variable that should be have zero value
		show = ""
		task = ""
		edit_num_raw = ""
		task_num_raw = ""

		//shorten the function
		reader := bufio.NewReader(os.Stdin)

		//main view
		fmt.Println("==============================")
		fmt.Println("----------TO-DO LIST----------")
		fmt.Println("==============================")
		fmt.Println("---")
		fmt.Println("*Main menu")
		fmt.Println("---")
		fmt.Println("1. Show Task")
		fmt.Println("2. Add Task")
		fmt.Println("3. Edit Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Println("---")
		fmt.Print("Input the number: ")
		opin_raw, _ := reader.ReadString('\n')
		opin_raw = strings.TrimSpace(opin_raw)

		for !containsInt(main_input_list, opin_raw) {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			fmt.Print("Input the number: ")
			opin_raw, _ = reader.ReadString('\n')
			opin_raw = strings.TrimSpace(opin_raw)
		}
		opin, _ = strconv.Atoi(opin_raw)
		fmt.Println("==============================")
		switch opin {
		case 1:
			show_task()
			main_view()
		case 2:
			add_task()
			main_view()
		case 3:
			edit_task()
			main_view()
		case 4:
			delete_task()
			main_view()
		}
	}
}

// case 1
func show_task() {
	for show != "back" {
		//Set the variable that should be have zero value

		//shorten the function
		reader := bufio.NewReader(os.Stdin)

		//show task view
		fmt.Println("==============================")
		fmt.Println("----------TO-DO LIST----------")
		fmt.Println("==============================")
		fmt.Println("*Note: input 'back' to move to Main Menu")
		fmt.Println("---")
		fmt.Println("*Show task")
		fmt.Println("---")
		if len(task_list) == 0 {
			fmt.Println("No tasks available")
		} else {
			fmt.Println("Your tasks:")
			for i, task := range task_list {
				fmt.Printf("%d. %s (%s)\n", i+1, task, task_status[i])
			}
		}
		fmt.Println("---")
		fmt.Println("*Mark your task as Done/On Progress/Not yet")
		fmt.Print("Enter task number: ")
		show_num, _ := reader.ReadString('\n')
		show = strings.TrimSpace(show_num)
		show_int, _ := strconv.Atoi(show)

		for show != "back" && (show_int < 1 || show_int > len(task_list)) {
			fmt.Println("Invalid task number")
			fmt.Print("Enter task number: ")
			show_num, _ = reader.ReadString('\n')
			show = strings.TrimSpace(show_num)
			show_int, _ = strconv.Atoi(show)
		}

		if show != "back" {
			fmt.Println("---")
			fmt.Println("1. Done")
			fmt.Println("2. On progress")
			fmt.Println("3. Not yet")
			fmt.Print("Enter status number: ")
			status, _ := reader.ReadString('\n')
			status = strings.TrimSpace(status)
			for !containsInt(show_status_list, status) {
				fmt.Println("Invalid status number")
				fmt.Print("Enter status number: ")
				status, _ = reader.ReadString('\n')
				status = strings.TrimSpace(status)
			}
			status_int, _ := strconv.Atoi(status)

			switch status_int {
			case 1:
				task_status[show_int-1] = "Done"
			case 2:
				task_status[show_int-1] = "On progress"
			case 3:
				task_status[show_int-1] = "Not yet"
			}
		}
		fmt.Println("==============================")
	}
}

// case 2
func add_task() {
	for task != "back" {
		//shorten the function
		reader := bufio.NewReader(os.Stdin)

		//input task view
		fmt.Println("==============================")
		fmt.Println("----------TO-DO LIST----------")
		fmt.Println("==============================")
		fmt.Println("*Note: input 'back' to move to Main Menu")
		fmt.Println("---")
		fmt.Println("*Add task")
		fmt.Println("---")
		if len(task_list) == 0 {
			fmt.Println("No tasks available")
		} else {
			fmt.Println("Available tasks:")
			for i, task := range task_list {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}
		fmt.Println("---")
		fmt.Print("Enter task: ")
		task_raw, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("Error:", err)
			fmt.Print("Enter task: ")
			task_raw, err = reader.ReadString('\n')
		}
		task = strings.TrimSpace(task_raw)
		if task != "back" {
			task_list = append(task_list, task)
			task_status = append(task_status, "Not yet")
		}
		fmt.Println("==============================")
	}
}

// case 3
func edit_task() {
	for edit_num_raw != "back" {
		//shorten the function
		reader := bufio.NewReader(os.Stdin)

		//edit task view
		fmt.Println("==============================")
		fmt.Println("----------TO-DO LIST----------")
		fmt.Println("==============================")
		fmt.Println("*Note: input 'back' to move to Main Menu")
		fmt.Println("---")
		fmt.Println("*Edit task")
		fmt.Println("---")
		if len(task_list) == 0 {
			fmt.Println("No tasks available")
		} else {
			fmt.Println("Available tasks:")
			for i, j := range task_list {
				fmt.Printf("%d. %s\n", i+1, j)
			}
		}
		fmt.Println("---")
		fmt.Print("Enter task number to edit: ")
		edit_num_raw, _ = reader.ReadString('\n')
		edit_num_raw = strings.TrimSpace(edit_num_raw)
		edit_num, _ := strconv.Atoi(edit_num_raw)
		for edit_num_raw != "back" && (edit_num < 1 || edit_num > len(task_list)) {
			fmt.Println("Invalid task number")
			fmt.Print("Enter task number to edit: ")
			edit_num_raw = strings.TrimSpace(edit_num_raw)
			edit_num, _ = strconv.Atoi(edit_num_raw)
		}

		if edit_num_raw != "back" {
			fmt.Print("Edit task: ")
			edit_task_raw, _ := reader.ReadString('\n')
			edit_task_raw = strings.TrimSpace(edit_task_raw)
			task_list[edit_num-1] = edit_task_raw
			task_status[edit_num-1] = "Not yet"
		}
		fmt.Println("==============================")
	}
}

// case 4
func delete_task() {
	for task_num_raw != "back" {
		//shorten the function
		reader := bufio.NewReader(os.Stdin)

		//delete task view
		fmt.Println("==============================")
		fmt.Println("----------TO-DO LIST----------")
		fmt.Println("==============================")
		fmt.Println("*Note: input 'back' to move to Main Menu")
		fmt.Println("---")
		fmt.Println("*Delete task")
		fmt.Println("---")
		if len(task_list) == 0 {
			fmt.Println("No tasks available")
		} else {
			fmt.Println("Available tasks:")
			for i, j := range task_list {
				fmt.Printf("%d. %s\n", i+1, j)
			}
		}
		fmt.Println("---")
		fmt.Print("Enter task number to delete: ")
		task_num_raw, _ = reader.ReadString('\n')
		task_num_raw = strings.TrimSpace(task_num_raw)
		task_num, _ := strconv.Atoi(task_num_raw)
		for task_num_raw != "back" && (task_num < 1 || task_num > len(task_list)) {
			fmt.Println("Invalid task number")
			fmt.Print("Enter task number to delete: ")
			task_num_raw, _ = reader.ReadString('\n')
			task_num_raw = strings.TrimSpace(task_num_raw)
			task_num, _ = strconv.Atoi(task_num_raw)
		}

		if task_num_raw != "back" {
			temp_list = []string{}
			for i := range task_list {
				if task_list[i] == task_list[task_num-1] {
					continue
				}
				temp_list = append(temp_list, task_list[i])
				temp_status = append(temp_status, task_status[i])
			}
			task_list = temp_list
		}
		fmt.Println("==============================")
	}
}

func Project() {
	main_view()
}
