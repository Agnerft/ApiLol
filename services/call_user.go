package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/agnerft/ApiLol/domain"
)

type ServiceRequest struct {
	httpUser *http.Client
}

const (
	url_padrao = "https://americas.api.riotgames.com"
	apiKey     = "RGAPI-bbfc07a8-e6e9-40cb-a5da-1771bba81071"
	url_me     = "riot/account/v1/accounts/by-riot-id"
)

func NewServiceUser() *ServiceRequest {
	return &ServiceRequest{
		httpUser: &http.Client{},
	}
}

func RequestUser(tagName, gameName string) (domain.User, error) {

	cli := &http.Client{}
	// response, err := s.httpUser.Get(fmt.Sprintf("%s/%s/%s/%s?api_key=%s", url_padrao, url_me, gameName, tagName, apiKey))
	// if err != nil {
	// 	fmt.Printf("Erro ao fazer a requisição HTTP, code %s", response.Status)
	// }
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s?api_key=%s", url_padrao, url_me, gameName, tagName, apiKey), nil)
	if err != nil {
		log.Fatalf("Erro para trazer o usuário %v", err)
	}

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatalf("Erro para fazer o response %v", err)

	}

	defer resp.Body.Close()

	var usr domain.User

	body, err := readBody(resp)
	if err != nil {
		return domain.User{}, err
	}

	err = json.Unmarshal([]byte(body), &usr)
	if err != nil {
		log.Fatal("Erro ao decodificar o JSON:", err)
		return domain.User{}, err
	}

	fmt.Println("Status:", resp.Status)

	fmt.Println(usr)

	return usr, nil

}

func readBody(res *http.Response) (string, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
