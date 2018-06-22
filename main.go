package main

import (
	"log"
	"github.com/gorilla/mux"
	"FootballTeamsPower/controllers"
	"net/http"
	"FootballTeamsPower/repositories"
	"FootballTeamsPower/services"
	"github.com/robfig/cron"
)

func init()  {
	log.Println("Check data in DB")
	countOfTeams, err := repositories.GetCountTeams()
	if err != nil {
		log.Println(err)
		return
	}
	if countOfTeams == 0 {
		log.Println("db is empty")
		services.ParsingData()
	}
	cron.New().AddFunc("@midnight", services.CheckNewMatches)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Headers("Content-Type", "application/json")
	router.HandleFunc("/power", controllers.GetPowers).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
