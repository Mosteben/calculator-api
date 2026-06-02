package models

type Request struct {
	A int    `json:"a"`
	B int    `json:"b"`
	Op string `json:"operation"`
}