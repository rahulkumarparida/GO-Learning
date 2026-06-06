package main

import (
	"webscarper/asyncronous"
	"webscarper/synchronous"
)

func main()  {
	// This Display how syncronous scraping works
	synchronous.SynchronousBehaviourScraping()
	// This display hwo asyncronous scraping works
	asyncronous.RoutineScrapingBehaviour()
}
