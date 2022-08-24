package controllers

import (
	"fmt"
)

func ShowUser(username string) error {
	fmt.Println("username : ", username)

	return nil
}
