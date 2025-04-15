package service_impl

import (
	"context"
	"fmt"

	v1 "github.com/go-kratos/kratos/v2/api/product/service/v1"
	// v1 "kratos-practice/kratos/api/product/service/v1"
)

type ProductServiceImpl struct {
	v1.UnimplementedProductServiceServer
}

func NewProductServiceImpl() *ProductServiceImpl {
	return &ProductServiceImpl{}
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *v1.CreateProductRequest) (*v1.ProductReply, error) {
	fmt.Printf("Creating product: %+v\n", req)
	return &v1.ProductReply{
		Id:          1,
		Name:        req.Name,
		Brand:       req.Brand,
		Description: req.Description,
		UserId:      req.UserId,
	}, nil
}

// Add GetProduct, UpdateProduct, DeleteProduct, etc. later similarly
