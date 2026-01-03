package models

type Stock struct {
	Symbol  string  `json:"symbol"`
	Amount  float64 `json:"amount"`
	OrderNo float64 `json:"orderNo"`
}

type SaveStocksRequest struct {
	Stocks []Stock `json:"stocks"`
}
