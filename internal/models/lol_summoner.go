package models

import (
	"fmt"
)

type Summoner struct {
	AccountId                   uint64 `json:"accountId"`
	DisplayName                 string `json:"displayName"`
	GameName                    string `json:"gameName"`
	InternalName                string `json:"internalName"`
	NameChangeFlag              bool   `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int    `json:"percentCompleteForNextLevel"`
	Privacy                     string `json:"privacy"`
	ProfileIconId               int    `json:"profileIconId"`
	Puuid                       string `json:"puuid"`
	RerollPoints                struct {
		CurrentPoints    int `json:"currentPoints"`
		MaxRolls         int `json:"maxRolls"`
		NumberOfRolls    int `json:"numberOfRolls"`
		PointsCostToRoll int `json:"pointsCostToRoll"`
		PointsToReroll   int `json:"pointsToReroll"`
	} `json:"rerollPoints"`
	SummonerId       uint64 `json:"summonerId"`
	SummonerLevel    int    `json:"summonerLevel"`
	TagLine          string `json:"tagLine"`
	Unnamed          bool   `json:"unnamed"`
	XpSinceLastLevel int    `json:"xpSinceLastLevel"`
	XpUntilNextLevel int    `json:"xpUntilNextLevel"`
}

func (s *Summoner) String() string {
	return fmt.Sprintf("%+v", *s)
}
