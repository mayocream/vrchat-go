package main

import (
	"github.com/mayocream/vrchat-go"
)

func main() {
	client := vrchat.NewClient("https://api.vrchat.cloud/api/1")

	err := client.Authenticate("username", "password", "totp")
	if err != nil {
		panic(err)
	}

	user, err := client.GetCurrentUser()
	if err != nil {
		panic(err)
	}

	println("logged in as ", user.DisplayName)
}
