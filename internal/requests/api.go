package requests

type API struct {
	LolSummoner LolSummonerPlugin
}

func NewAPI(lolsummoner LolSummonerPlugin) *API {
	return &API{
		LolSummoner: lolsummoner,
	}
}
