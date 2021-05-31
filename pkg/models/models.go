package models

type Servers struct {
	Interfaces []Interface `json:"Interfaces"`
}

type Interface struct {
	ServerName string  `json:"servername"`
	Name       string  `json:"id"`
	Values     float64 `json:"values"`
}
