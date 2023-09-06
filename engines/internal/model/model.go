package model

type GetEnginesByIDReq struct {
	CarsID []int `json:"cars_id"`
}

type GetEnginesByIDRes struct {
	Engines []Engine `json:"engines"`
}

type Engine struct {
	Type       string `json:"type"`
	Horsepower string `json:"horsepower"`
}

type HTTPResponse struct {
	Data any    `json:"data"`
	Err  string `json:"err"`
}
