package requests

import (
	"fmt"
	"log"

	"github.com/braycarlson/asol"
	"github.com/pipexlul/aramsnatcher/internal/models"
	myutils "github.com/pipexlul/aramsnatcher/utils"
)

const (
	prefix = "/lol-summoner/v1/"
)

type LolSummoner struct{}

// Endpoint builder
func ep(resource string) string {
	return fmt.Sprintf("%s%s", prefix, resource)
}

func (sum *LolSummoner) GetCurrentSummoner(a *asol.Asol) (*models.Summoner, error) {
	resp, err := myutils.CreateAndExecuteRequest(a, models.GET, ep("current-summoner"), nil)

	if err != nil {
		log.Println("GetCurrentSummoner Error")
		fmt.Println(err)
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
