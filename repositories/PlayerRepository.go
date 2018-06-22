package repositories

import "FootballTeamsPower/models"

func CreatePlayer(player models.Player) (models.Player, error) {
	err := db.QueryRow(`INSERT INTO players(name, team_id, link) VALUES ($1, $2, $3) RETURNING id`,
		player.Name, player.TeamId, player.Link).Scan(&player.Id)
	return player, err
}

func GetCountPlayers() (int, error) {
	rows, err := db.Query(`SELECT COUNT(*) FROM players`)
	var count int
	if err != nil {
		return count, err
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return count, err
}

func FindPlayersByTeamId(teamId int) ([]models.Player, error) {
	var players []models.Player
	rows, err := db.Query(`SELECT id, name, link FROM players WHERE team_id = $1`, teamId)
	defer rows.Close()
	if err != nil {
		return players, err
	}
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.Id, &player.Name, &player.Link)
		if err != nil {
			return players, err
		}
		player.TeamId = teamId
		players = append(players, player)
	}
	return players, nil
}

func FindAllPlayers() ([]models.Player, error) {
	var players []models.Player
	rows, err := db.Query(`SELECT id, name, team_id, link FROM players`)
	defer rows.Close()
	if err != nil {
		return players, err
	}
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.Id, &player.Name, &player.TeamId, &player.Link)
		if err != nil {
			return players, err
		}
		players = append(players, player)
	}
	return players, nil
}
