package requests

import (
	"github.com/braycarlson/asol"

	"github.com/pipexlul/aramsnatcher/internal/models"
)

type LolSummonerPlugin interface {
	GetCurrentSummoner(a *asol.Asol) (*models.Summoner, error)
	GetCurrentSummonerStatus(a *asol.Asol) (*models.SummonerStatus, error)
}
