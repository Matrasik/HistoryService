package model

import "time"

type PairsReq struct {
	name  string
	queue []int
	date  []time.Time
}
