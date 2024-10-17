package main

import (
	"fmt"

	"github.com/agnerft/ApiLol/services"
)

func main() {

	tagName := "Agner"
	gameName := "IrmãoDoJorel"
	// url_padrao := "https://americas.api.riotgames.com"
	// apiKey := "RGAPI-bbfc07a8-e6e9-40cb-a5da-1771bba81071"
	// url_me := fmt.Sprintf("riot/account/v1/accounts/by-riot-id/%s/%s", gameName, tagName)

	// cli := &http.Client{}

	// req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s?api_key=%s", url_padrao, url_me, apiKey), nil)
	// if err != nil {
	// 	log.Fatalf("Erro para trazer o usuário %v", err)
	// }

	// resp, err := cli.Do(req)
	// if err != nil {
	// 	log.Fatalf("Erro para fazer o response %v", err)

	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Erro para ler o body %v", err)
	// }

	// fmt.Println("Status:", resp.Status)

	// fmt.Println(string(body))

	response, err := services.RequestUser(tagName, gameName)
	if err != nil {
		fmt.Println("Erro")
	}

	fmt.Println(response.GameName)

}
