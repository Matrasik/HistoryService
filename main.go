package main

import (
	"HistoryService/api"
	"HistoryService/config"
	"HistoryService/persistence/migration"
	"HistoryService/service"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"sync"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db, err := gorm.Open("postgres", "user = db password=1234567 port=54320 dbname=currency sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(false)

	migration.Migrate(db)
	defer db.Close()
	wg := &sync.WaitGroup{}
	conf := config.New()
	for _, curr := range conf.CurrPairs {
		wg.Add(1)
		go func(name string) {
			p := api.NewPairs()
			p.RequestFlow(name, db, wg, conf)
		}(curr)
	}
	var timeFrom, hmsFrom, timeTo, hmsTo, curr string
	for {
		fmt.Println("Enter date from(in format dd/mm/yyyy hours:minutes:seconds")
		fmt.Scanln(&timeFrom, &hmsFrom, &timeTo, &hmsTo, curr)
		fmt.Println("Enter date to(in format dd/mm/yyyy hours:minutes:seconds")
		fmt.Scanln(&timeTo, &hmsTo)
		fmt.Println("Enter currency pair")
		fmt.Scanln(&curr)
		service.Searcher(curr, StrToTime(timeFrom+" "+hmsFrom), StrToTime(timeTo+" "+hmsTo), db)
	}

	wg.Wait()
}

func StrToTime(date string) time.Time {
	t, err := time.Parse("2/1/2006 15:04:05", date)
	if err != nil {
		fmt.Println("error in parse time", err)
	}
	fmt.Println(t)
	return t
}
