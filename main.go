package main

import (
	"One.3/cache"
	"fmt"
)

func main() {
	username := "yxh"
	password, err := cache.GetUserInformation(username)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Password for user %s: %s\n", username, password)
	}
}
