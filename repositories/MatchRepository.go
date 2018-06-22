package repositories

import "FootballTeamsPower/models"

func CreateMatch(match models.Match) (models.Match, error) {
	err := db.QueryRow(`INSERT INTO matches(diff_score, opponent, player_id) VALUES ($1, $2, $3) RETURNING id`,
		match.DifferenceOfScore, match.Opponent, match.PlayerId).Scan(&match.Id)
	return match, err
}

func GetCountMatches() (int, error) {
	rows, err := db.Query(`SELECT COUNT(*) FROM matches`)
	var count int
	if err != nil {
		return count, err
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return count, err
}

func FindMatchesByPlayerId(playerId int) ([]models.Match, error) {
	var matches []models.Match
	rows, err := db.Query(`SELECT id, diff_score, opponent FROM matches WHERE player_id = $1`, playerId)
	defer rows.Close()
	if err != nil {
		return matches, err
	}
	for rows.Next() {
		var match models.Match
		err := rows.Scan(&match.Id, &match.DifferenceOfScore, &match.Opponent)
		if err != nil {
			return matches, err
		}
		match.PlayerId = playerId
		matches = append(matches, match)
	}
	return matches, nil
}

func GetCountMatchesByPlayerId(playerId int) (int, error) {
	rows, err := db.Query(`SELECT COUNT(*) FROM matches WHERE player_id = $1`, playerId)
	var count int
	if err != nil {
		return count, err
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return count, err
}