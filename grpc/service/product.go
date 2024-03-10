package service

import "context"

type productService struct {
}

var ProductService = &productService{}

// GetProductStock 获取产品库存
func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	// 实现具体的业务逻辑
	stock := p.GetStockById(request.ProdId)
	return &ProductResponse{ProdStock: stock}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return id
}
