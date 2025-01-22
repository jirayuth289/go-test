package basic

import (
	"fmt"
)

type LoginError struct {
	Username string
	Message  string
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s", e.Username, e.Message)
}

func login(username, password string) error {
	if username != "admin" || password != "password123" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}

	return nil
}

func testHandleErrorLogin() {
	err := login("admin", "password")
	if err != nil {
		switch e := err.(type) {
		case *LoginError:
			fmt.Println("Custom error occurred:", e)
		default:
			// Other types of errors
			fmt.Println("Generic error occurred:", e)
		}
		return
	}
}
