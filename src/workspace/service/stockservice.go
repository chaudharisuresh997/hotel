package service

type StockService interface {
	AddStock(Stock) SResponse
	GetStock(stockid string) (interface{}, SResponse)
}
type Stock struct {
}

func AddStock(Stock) SResponse {
	return SResponse{}
}

func GetStock(stockid string) (interface{}, SResponse) {
	return "", SResponse{}
}
