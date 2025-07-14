package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type contact struct {
	Name  string
	Email string
	Phone string
}

var contactsList = make(map[int]contact)
var idCounter = 1

func showContacts() {
	if len(contactsList) == 0 {
		fmt.Println("Your to-do list is empty.")
		return
	}

	fmt.Println("To-Do List:")

	ids := make([]int, 0, len(contactsList))
	for id := range contactsList {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		fmt.Printf("[%d] %s\nMail: %s\nPhone %s\n", id, contactsList[id].Name, contactsList[id].Email, contactsList[id].Phone)
		fmt.Println("--------------------------------------------------")
	}
}

func addContact(name string, email string, phone string) {
	contact := contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	contactsList[idCounter] = contact
	idCounter++
	fmt.Println("Contact added.")
}

func deleteContact(id int) {
	if _, exists := contactsList[id]; exists {
		delete(contactsList, id)
		for i := id + 1; i < idCounter; i++ {
			contactsList[i-1] = contactsList[i]
		}
		delete(contactsList, idCounter-1)
		idCounter--
		fmt.Println("Contact deleted and IDs adjusted.")
	} else {
		fmt.Println("Contact not found.")
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
			showContacts()
		case "ADD":
			fmt.Print("Name: ")
			if !scanner.Scan() {
				break
			}
			name := scanner.Text()
			fmt.Print("Email: ")
			if !scanner.Scan() {
				break
			}
			email := scanner.Text()
			fmt.Print("Phone: ")
			if !scanner.Scan() {
				break
			}
			phone := scanner.Text()
			addContact(name, email, phone)
		case "DELETE":
			showContacts()
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
			deleteContact(id)
		case "EXIT":
			fmt.Println("Exiting the To-Do List application.")
			return
		default:
			fmt.Println("Invalid action. Please use SHOW, ADD, DELETE, or EXIT.")
		}
	}
}
