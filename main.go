package main

import (
	"github.com/ignratnan/mini-projects/cli-calculator/calculator"
	"github.com/ignratnan/mini-projects/contacts-manager/contacts"
	"github.com/ignratnan/mini-projects/file-downloader/downloader"
	"github.com/ignratnan/mini-projects/json-read-write/jsonrew"
	"github.com/ignratnan/mini-projects/to-do-list/todolist"
)

func main() {
	opt := 5

	switch opt {
	case 1:
		calculator.Project()
	case 2:
		todolist.Project()
	case 3:
		contacts.Project()
	case 4:
		downloader.Project()
	case 5:
		jsonrew.Project()
	}
}
