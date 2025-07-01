package main

import (
	"github.com/ignratnan/mini-projects/cli-calculator/calculator"
	"github.com/ignratnan/mini-projects/to-do-list/todolist"
)

func main() {
	opt := 2

	switch opt {
	case 1:
		calculator.Project()
	case 2:
		todolist.Project()
	}
}
