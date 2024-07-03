package entity

type TradingResponse struct {
	Count int           `json:"count"`
	Total int           `json:"total"`
	Data  []TradingData `json:"data"`
}

type DailyDataResponse struct {
	Count int         `json:"count"`
	Total int         `json:"total"`
	Data  []DailyData `json:"data"`
}

type BasicResponse struct {
	Count int         `json:"count"`
	Total int         `json:"total"`
	Data  []BasicData `json:"data"`
	Meta  Meta        `json:"meta"`
}
