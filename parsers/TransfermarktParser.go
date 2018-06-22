package parsers

import (
	"net/http"
	"strings"
	"math"
	"strconv"
	"log"
	"github.com/PuerkitoBio/goquery"
)

const (
	transferMarktUrl                     = "https://www.transfermarkt.com"
	urlTeams                             = "https://www.transfermarkt.com/weltmeisterschaft-2018/teilnehmer/pokalwettbewerb/WM18"
	selectorTeams                        = "div #yw1 table.items td.links.no-border-links.hauptlink a.vereinprofil_tooltip"
	selectorPlayers                      = "div #yw1 span.hide-for-small a"
	tableWithAllMatchesOnePlayerSelector = "div.large-8.columns div.box div.responsive-table table"
	minutesPlayedInMatchSelector         = "td.rechts"
	winningMatchScoreSelector            = "td a span.greentext"
	opponentSelector                     = "td.no-border-links.hauptlink a"
)

type Player struct {
	Name       string
	PlayerLink string
}
type Team struct {
	Name     string
	TeamLink string
}

type Match struct {
	DifferenceOfScore int
	Opponent          string
}

func GetAllTeams() []Team {
	url, err := http.Get(urlTeams)
	if err != nil {
		log.Println(err)
	}
	document, err := goquery.NewDocumentFromReader(url.Body)
	if err != nil {
		log.Println(err)
	}
	var teams []Team
	document.Find(selectorTeams).Each(
		func(index int, item *goquery.Selection) {
			link, _ := item.Attr("href")
			teams = append(teams, Team{Name: item.Text(), TeamLink: link})
		})
	return teams
}

func GetPlayersByTeam(team Team) [] Player {
	response, err := http.Get(transferMarktUrl + team.TeamLink)
	if err != nil {
		log.Println(err)
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
	}
	var players []Player
	document.Find(selectorPlayers).Each(
		func(index int, item *goquery.Selection) {
			link, _ := item.Attr("href")
			players = append(players, Player{item.Text(), link})
		})
	return players
}

func GetAllWinningAndInSquadMatchesByPlayer(player Player) []Match {
	response, err := http.Get(transferMarktUrl + player.PlayerLink)
	if err != nil {
		log.Println(err)
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
	}
	var matches []Match
	document.Find(tableWithAllMatchesOnePlayerSelector).Last().Find("tr").Each(
		func(index int, item *goquery.Selection) {
			if inSquad(item) && isWinningMatch(item) {
				score := item.Find(winningMatchScoreSelector).First().Text()
				opponent := item.Find(opponentSelector).Text()
				matches = append(matches, Match{getDifferenceOfScore(score), opponent})
			}
		})
	return matches
}

func inSquad(trFromTableWithAllMatches *goquery.Selection) bool {
	return trFromTableWithAllMatches.Find(minutesPlayedInMatchSelector).Length() > 0
}

func isWinningMatch(trFromTableWithAllMatches *goquery.Selection) bool {
	return trFromTableWithAllMatches.Find(winningMatchScoreSelector).Length() > 0
}

func getDifferenceOfScore(score string) int {
	scores := strings.Split(strings.Replace(score, " ", ":", -1), ":")
	firstNumber, err := strconv.Atoi(scores[0])
	if err != nil {
		log.Println(score)
	}
	secondNumber, err := strconv.Atoi(scores[1])
	if err != nil {
		log.Println(score)
	}
	return int(math.Abs(float64(firstNumber - secondNumber)))
}
