package models

/*type GetInfoResponse struct {
	Interfaces []Interface `json:"Servers"`
}*/

type Servers struct {
	//	Name       string      `json:"servername"`
	Interfaces []Interface `json:"Interfaces"`
}

type Interface struct {
	ServerName string  `json:"servername"`
	Name       string  `json:"id"`
	Values     float64 `json:"values"`
}
