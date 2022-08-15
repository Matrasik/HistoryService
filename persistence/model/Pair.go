package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/datatypes"
)

type Pair struct {
	DoneAt    datatypes.Date `gorm:"type:timestamp"`
	PairsJson postgres.Jsonb `gorm:"type:jsonb;column:pair;"`
	Name      string         `gorm:"type:varchar"`
	PairsReq  PairsReq       `gorm:"-"`
}

func (p *Pair) MarshallPairsReq() error {
	var err error
	p.PairsJson.RawMessage, err = json.Marshal(p.PairsReq)

	return err
}

func (p *Pair) UnmarshallPairsReq() error {
	return json.Unmarshal(p.PairsJson.RawMessage, &p.PairsReq)
}
