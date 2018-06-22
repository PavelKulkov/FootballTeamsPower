package services

import (
	"FootballTeamsPower/parsers"
	"FootballTeamsPower/repositories"
	"FootballTeamsPower/models"
	"log"
)

func ParsingData() {
	log.Println("Run parsing data")
	teams := parsers.GetAllTeams()
	for _, parsedTeam := range teams {
		var team models.Team
		team.Name = parsedTeam.Name
		team.Link = parsedTeam.TeamLink
		dbTeam, err := repositories.CreateTeam(team)
		if err != nil {
			log.Fatal(err)
			continue
		}
		go parsePlayersByTeam(parsedTeam, dbTeam.Id)
	}
}

func parsePlayersByTeam(team parsers.Team, teamId int) {
	players := parsers.GetPlayersByTeam(team)
	for _, parsedPlayer := range players {
		var player models.Player
		player.Name = parsedPlayer.Name
		player.Link = parsedPlayer.PlayerLink
		player.TeamId = teamId
		dbPlayer, err := repositories.CreatePlayer(player)
		if err != nil {
			log.Println(err)
			continue
		}
		//TODO Так парсинг намного быстрее, но в таком случаи transfermarkt банит
		//go parseMatchesByPlayer(parsedPlayer, dbPlayer.Id)
		parseMatchesByPlayer(parsedPlayer, dbPlayer.Id)
	}
}

func parseMatchesByPlayer(player parsers.Player, playerId int) {
	matches := parsers.GetAllWinningAndInSquadMatchesByPlayer(player)
	for _, parsedMatch := range matches {
		var match models.Match
		match.DifferenceOfScore = parsedMatch.DifferenceOfScore
		match.Opponent = parsedMatch.Opponent
		match.PlayerId = playerId
		if _, err := repositories.CreateMatch(match); err != nil {
			log.Println(err)
		}
	}
}

func CheckNewMatches() {
	players, err := repositories.FindAllPlayers()
	if err != nil {
		log.Println(err)
	}

	for _, player := range players {
		//TODO transfermarkt банит, подумать как обойти
		// go checkNewMatchesForPlayer(player)
		checkNewMatchesForPlayer(player)
	}

}

func checkNewMatchesForPlayer(player models.Player) {
	countOfmatches, err := repositories.GetCountMatchesByPlayerId(player.Id)
	if err != nil {
		log.Println(err)
		return
	}
	playerForParse := parsers.Player{Name: player.Name, PlayerLink: player.Link}
	winningAndInSquadMatchesByPlayer := parsers.GetAllWinningAndInSquadMatchesByPlayer(playerForParse)
	if countOfmatches != len(winningAndInSquadMatchesByPlayer) {
		for i := countOfmatches; i < len(winningAndInSquadMatchesByPlayer); i++ {
			parsedMatch := winningAndInSquadMatchesByPlayer[i]
			var match models.Match
			match.DifferenceOfScore = parsedMatch.DifferenceOfScore
			match.Opponent = parsedMatch.Opponent
			match.PlayerId = player.Id
			if _, err := repositories.CreateMatch(match); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
