package plugin

import "time"

type Query struct {
	Tags string `json:"tags"`
}

type DataQuery struct {
	Query
	Start time.Time
	End   time.Time
}

type RawDataQuery struct {
	DataQuery
	Direction  int `json:"direction"`
	Count      int `json:"count"`
	MinQuality int `json:"minQuality"`
}
