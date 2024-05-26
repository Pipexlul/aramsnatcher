package requests

import (
	"fmt"
	"log"

	"github.com/braycarlson/asol"
	"github.com/pipexlul/aramsnatcher/internal/models"
	myutils "github.com/pipexlul/aramsnatcher/utils"
)

type LolSummoner struct{}

func (sum *LolSummoner) GetCurrentSummoner(a *asol.Asol) (*models.Summoner, error) {
	req, err := a.NewGetRequest("/lol-summoner/v1/current-summoner")
	if err != nil {
		log.Println("Request Error")
		fmt.Println(err)

		return nil, err
	}

	resp, err := a.RiotRequest(req)
	if err != nil {
		log.Println("Response Error")
		fmt.Println(err)

		return nil, err
	}

	var summoner models.Summoner
	err = myutils.MapToStruct(resp.(map[string]interface{}), &summoner)
	if err != nil {
		log.Println("MapToStruct Error")
		fmt.Println(err)

		return nil, err
	}

	return &summoner, nil
}

func (sum *LolSummoner) GetCurrentSummonerStatus() {
	panic("not implemented") // TODO: Implement
}
