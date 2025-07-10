package main

import (
	"github.com/ignratnan/mini-projects/cli-calculator/calculator"
	"github.com/ignratnan/mini-projects/contacts-manager/contacts"
	"github.com/ignratnan/mini-projects/to-do-list/todolist"
)

func main() {
	opt := 3

	switch opt {
	case 1:
		calculator.Project()
	case 2:
		todolist.Project()
	case 3:
		contacts.Project()
	}
}
