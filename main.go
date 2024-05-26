package main

import (
	"fmt"
	"log"

	"github.com/braycarlson/asol"

	"github.com/pipexlul/aramsnatcher/internal/requests"
)

type Client struct {
	*asol.Asol

	api *requests.API
}

func NewClient() *Client {
	return &Client{
		Asol: asol.NewAsol(),
		api:  requests.NewAPI(&requests.LolSummoner{}),
	}
}

func main() {
	client := NewClient()

	client.OnOpen(func(a *asol.Asol) {
		log.Println("OnOpen")
	})

	client.OnReady(func(a *asol.Asol) {
		log.Println("OnReady")
	})

	client.OnClientClose(func(a *asol.Asol) {
		log.Println("OnClientClose")
	})

	client.OnLogin(func(a *asol.Asol) {
		log.Println("OnLogin")

		// summoner, err := client.api.LolSummoner.GetCurrentSummoner(a)
		// if err != nil {
		// 	log.Println("Error getting current summoner")
		// 	fmt.Println(err)
		// }

		// fmt.Println(summoner)

		status, err := client.api.LolSummoner.GetCurrentSummonerStatus(a)
		if err != nil {
			log.Println("Error getting current summoner status")
			fmt.Println(err)

			return
		}

		fmt.Println(status)
	})

	client.OnLogout(func(a *asol.Asol) {
		log.Println("OnLogout")
	})

	client.OnError(func(err error) {
		log.Println("OnError")
		log.Println(err)
	})

	client.OnReconnect(func(a *asol.Asol) {
		log.Println("OnReconnect")
	})

	client.OnWebsocketClose(func(a *asol.Asol) {
		log.Println("OnWebsocketClose")
	})

	client.OnMessage(
		"/lol-settings/v1/account/lol-collection",
		"Update",
		func(a *asol.Asol, m *asol.Message) {
			log.Println("OnMessage")
			fmt.Println(m)
		},
	)

	client.Start()
}
