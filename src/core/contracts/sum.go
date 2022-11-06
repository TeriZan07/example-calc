package contracts

type CalcReq struct {
	Numbers []float64 `json:"numbers"`
}

type CalcResp struct {
	Result float64 `json:"result"`
	Count  int     `json:"count"`
}
