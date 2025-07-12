package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

var toDoList = make(map[int]string)
var idCounter = 1

func showToDoList() {
	if len(toDoList) == 0 {
		fmt.Println("Your to-do list is empty.")
		return
	}

	fmt.Println("To-Do List:")

	ids := make([]int, 0, len(toDoList))
	for id := range toDoList {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		fmt.Printf("%d: %s\n", id, toDoList[id])
	}
}

func addToDoItem(scanner *bufio.Scanner) {
	item := scanner.Text()
	toDoList[idCounter] = item
	idCounter++
	fmt.Println("Item added.")
}

func deleteTodoItem(id int) {
	if _, exists := toDoList[id]; exists {
		delete(toDoList, id)
		for i := id + 1; i < idCounter; i++ {
			toDoList[i-1] = toDoList[i]
		}
		delete(toDoList, idCounter-1)
		idCounter--
		fmt.Println("Item deleted and IDs adjusted.")
	} else {
		fmt.Println("Item not found.")
	}
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage:", os.Args[0])
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("?> What do you want to do? (SHOW | ADD | DELETE | EXIT)")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		switch strings.ToUpper(input) {
		case "SHOW":
			showToDoList()
		case "ADD":
			fmt.Println("Enter a new to-do item:")
			if !scanner.Scan() {
				break
			}
			addToDoItem(scanner)
		case "DELETE":
			showToDoList()
			fmt.Println("Enter the item number to delete:")
			if !scanner.Scan() {
				break
			}
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}
			deleteTodoItem(id)
		case "EXIT":
			fmt.Println("Exiting the To-Do List application.")
			return
		default:
			fmt.Println("Invalid action. Please use SHOW, ADD, DELETE, or EXIT.")
		}
	}
}
