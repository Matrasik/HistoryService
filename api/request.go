package api

import (
	"HistoryService/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const PORT = ":8000"

func (p *Pairs) GetPair(ctx context.Context, curr string) {
	url := "http://localhost" + PORT + "/" + curr
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error response", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatalln("cant unmarshall http response", err)
	}

}

type Pairs struct {
	Name string      `json:"name" gorm:"type:varchar(50)"`
	Curr []int       `json:"curr" gorm:"type:int"`
	Date []time.Time `json:"date" gorm:"type:timestamp"`
}

func NewPairs() *Pairs {
	return &Pairs{}
}

func (p *Pairs) RequestFlow(curr string, db *gorm.DB, wg *sync.WaitGroup, conf *config.Config) {
	defer wg.Done()
	for {
		ctx := context.Background()
		p.GetPair(ctx, curr)
		for i := range p.Curr {
			db.Exec("insert into pairs (done_at,price, name) values ($1 , $2, $3)", p.Date[i], p.Curr[i], p.Name)
		}
		time.Sleep(time.Duration(conf.UpdateTime) * time.Second)
	}
}
