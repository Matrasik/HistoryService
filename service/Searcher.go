package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Pairs struct {
	DoneAt time.Time
	Price  int
	Name   string
}

func Searcher(currency string, timeFrom time.Time, timeTo time.Time, db *gorm.DB) {
	var pair []Pairs
	db.Raw("SELECT * FROM pairs WHERE name = $1 AND done_at BETWEEN $2 AND $3", currency, timeFrom, timeTo).Scan(&pair)
	fmt.Println(pair)
}
