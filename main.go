package main

import (
	"fmt"

	"github.com/agnerft/ApiLol/services"
)

func main() {

	tagName := "Agner"
	gameName := "IrmãoDoJorel"

	response, err := services.RequestUser(tagName, gameName)
	if err != nil {
		fmt.Println("Erro")
	}

	fmt.Println(response.Puuid)

}
