package service

import "context"

type ProductService interface {
	CreateProduct(ctx context.Context, name, brand, description string, userID int32) error
}
