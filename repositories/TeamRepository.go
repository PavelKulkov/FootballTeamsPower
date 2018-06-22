package repositories

import (
	"FootballTeamsPower/models"
)

func CreateTeam(team models.Team) (models.Team, error) {
	err := db.QueryRow(`INSERT INTO teams(name, link) VALUES ($1, $2) RETURNING id`,
		team.Name, team.Link).Scan(&team.Id)
	return team, err
}

func GetCountTeams() (int, error){
	rows, err := db.Query(`SELECT COUNT(*) FROM teams`)
	var count int
	if err != nil {
		return count, err
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return count, err
}

func FindTeamByName(name string) (models.Team, error) {
	var dbTeam models.Team
	err := db.QueryRow(`SELECT id, name FROM teams WHERE name = $1`, name).Scan(&dbTeam.Id, &dbTeam.Name)
	return dbTeam, err
}
