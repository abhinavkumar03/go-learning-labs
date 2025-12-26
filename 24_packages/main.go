package main

import (
	"fmt"

	"github.com/abhinavkumar03/go-learning-labs/auth"
	"github.com/abhinavkumar03/go-learning-labs/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("abhinavkumar03", "password")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "Abhinav Kumar",
	}

	fmt.Println(user.Name)
	color.Red(user.Email)

}
