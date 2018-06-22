package controllers

import (
	"net/http"
	"encoding/json"
	"FootballTeamsPower/repositories"
)

type TeamsDto struct {
	FirstTeamName  string
	SecondTeamName string
}
type PowersDto struct {
	FirstTeamPower  int
	SecondTeamPower int
}

func GetPowers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var teams TeamsDto
	err := decoder.Decode(&teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	resultDto, err := calcStrengths(teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultDto)
}

func calcStrengths(teams TeamsDto) (PowersDto, error) {
	var result PowersDto
	firstTeam, err := repositories.FindTeamByName(teams.FirstTeamName)
	if err != nil {
		return result, err
	}
	secondTeam, err := repositories.FindTeamByName(teams.SecondTeamName)
	if err != nil {
		return result, err
	}
	result.FirstTeamPower, err = getTeamStrength(firstTeam.Id)
	if err != nil {
		return result, err
	}
	result.SecondTeamPower, err = getTeamStrength(secondTeam.Id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getTeamStrength(teamId int) (int, error) {
	var teamStrength int
	players, err := repositories.FindPlayersByTeamId(teamId)
	if err != nil {
		return teamStrength, err
	}
	for _, player := range players {
		playerStrength, err := getStrengthPlayer(player.Id)
		if err != nil {
			return teamStrength, err
		}
		teamStrength += playerStrength
	}
	return teamStrength, nil
}
func getStrengthPlayer(playerId int) (int, error) {
	var playerStrength int
	matches, err := repositories.FindMatchesByPlayerId(playerId)
	if err != nil {
		return playerStrength, err
	}
	for _, match := range matches {
		playerStrength += match.DifferenceOfScore
	}
	return playerStrength, nil
}
