package models

type HistoryItem struct {
	A        int     `json:"a"`
	B        int     `json:"b"`
	Op       string  `json:"operation"`
	Result   float64 `json:"result"`
}