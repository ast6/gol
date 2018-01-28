package main


type Req struct {
	Ccy     string `json:"ccy"`
	BaseCcy string `json:"base_ccy"`
	Buy     string `json:"buy"`
	Sale    string `json:"sale"`
}
type CalculateRequest struct {
	Amount float32 `json:"amount"`
}

type CalculateResponse struct {
	Amount float32 `json:"amount"`
	Usd    float32 `json:"usd"`
	Eur    float32 `json:"eur"`
	Btc    float32 `json:"btc"`
	Rur    float32 `json:"rur"`
}
