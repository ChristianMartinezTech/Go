// Small program exploring structs

package main

import "fmt"

type domainAdmin struct {
	Username   string
	Password   string
	Domain     string
	NoAccounts int
}

func main() {
	user1 := &domainAdmin{
		Username:   "Chris",
		Password:   "12345567890",
		Domain:     "tothepoint.app",
		NoAccounts: 1,
	}

	// You can also initializa a struct like this
	user2 := domainAdmin{"Carlos", "0987654321", "tothepoint.app", 2}

	fmt.Printf("User: %s\n%s\n", user1.Username, user1)
	fmt.Printf("User: %s\n%s\n", user2.Username, user2)

	// Changing the values in a Struct instance
	user2.Username = "Alejandro"
	fmt.Println("New user2 Username: " + user2.Username)

	// OR INITIALIZE A STRUCT with new()
	user3 := new(domainAdmin)
	user3.Username = "Rolando"
	user3.Password = "Holi123"
	user3.Domain = "tothepoint.app"
	user3.NoAccounts = 1
	fmt.Println(user3.Username)
	fmt.Println(user3)
}
