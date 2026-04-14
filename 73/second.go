package main

import "fmt"

type userWithPointers struct {
	username string
	password string
}

func (u *userWithPointers) print() {
	fmt.Printf("Username: %s, Password: %s\n", u.username, u.password)
}

func (u *userWithPointers) changePassword(newPassword string) {
	u.password = newPassword
}

func (u *userWithPointers) changeUsername(newUsername string) {
	u.username = newUsername
}

type userWithValues struct {
	username string
	password string
}

func (u userWithValues) print() {
	fmt.Printf("Username: %s, Password: %s\n", u.username, u.password)
}

func (u userWithValues) changePassword(newPassword string) {
	u.password = newPassword
}

func (u userWithValues) changeUsername(newUsername string) {
	u.username = newUsername
}

func main() {
	fmt.Println("Hello, from second.go")

	user1 := userWithPointers{username: "john_doe", password: "12345"}
	user1.print()
	user1.changePassword("new_password")
	user1.changeUsername("john_doe_updated")
	user1.print()

	user2 := userWithValues{username: "jane_doe", password: "54321"}
	user2.print()
	user2.changePassword("new_password")
	user2.changeUsername("jane_doe_updated")
	user2.print()
}
