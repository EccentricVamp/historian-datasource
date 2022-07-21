package plugin

import (
	"time"
)

type Sample struct {
	TimeStamp time.Time `json:"TimeStamp"`
	Value     string    `json:"Value"`
	Quality   int       `json:"Quality"`
}

type Datum struct {
	DataType  string   `json:"DataType"`
	ErrorCode int      `json:"ErrorCode"`
	TagName   string   `json:"TagName"`
	Samples   []Sample `json:"Samples"`
}

type Response struct {
	ErrorCode    int     `json:"ErrorCode"`
	ErrorMessage string  `json:"ErrorMessage"`
	Data         []Datum `json:"Data"`
}
