package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllTeams(t *testing.T) {
	teams := GetAllTeams()
	assert.Equal(t, 32, len(teams))
	assert.Equal(t, "France", teams[0].Name)
	assert.Equal(t, "/frankreich/startseite/verein/3377", teams[0].TeamLink)
	assert.Equal(t, "Spain", teams[1].Name)
	assert.Equal(t, "/spanien/startseite/verein/3375", teams[1].TeamLink)
	assert.Equal(t, "Brazil", teams[2].Name)
	assert.Equal(t, "/brasilien/startseite/verein/3439", teams[2].TeamLink)
}

func TestGetPlayersByTeam(t *testing.T) {
	players := GetPlayersByTeam(Team{Name:"France", TeamLink:"/frankreich/startseite/verein/3377"})
	assert.Equal(t, "Hugo Lloris", players[0].Name)
	assert.Equal(t, "/hugo-lloris/nationalmannschaft/spieler/17965", players[0].PlayerLink)
	assert.Equal(t, "Alphonse Areola", players[1].Name)
	assert.Equal(t, "/alphonse-areola/nationalmannschaft/spieler/120629", players[1].PlayerLink)
	assert.Equal(t, "Steve Mandanda", players[2].Name)
	assert.Equal(t, "/steve-mandanda/nationalmannschaft/spieler/23951", players[2].PlayerLink)
}

func TestGetAllWinningAndInSquadMatchesByPlayer(t *testing.T) {
	player := Player{Name: "Hugo Lloris", PlayerLink: "/hugo-lloris/nationalmannschaft/spieler/17965"}
	matches := GetAllWinningAndInSquadMatchesByPlayer(player)
	assert.Equal(t, "Turkey", matches[0].Opponent)
	assert.Equal(t, 1, matches[0].DifferenceOfScore)
	assert.Equal(t, "Faroe Islands", matches[1].Opponent)
	assert.Equal(t, 1, matches[1].DifferenceOfScore)
}
