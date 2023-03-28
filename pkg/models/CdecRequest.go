package models

type CdecRequest struct {
	FromLocation string `json:"from_location"`
	ToLocation   string `json:"to_location"`
	Weight       int    `json:"weight"`
	Lenght       int    `json:"length"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type CdecRequestDto struct {
	FromLocation Location  `json:"from_location"`
	ToLocation   Location  `json:"to_location"`
	Packages     []Package `json:"packages"`
}

type Package struct {
	Weight int `json:"weight"`
	Lenght int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Location struct {
	Address string `json:"address"`
}
