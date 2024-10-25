package services

import (
	"context"
	"github.com/bohexists/product-fetcher-svc/api"
	"github.com/bohexists/product-fetcher-svc/internal/adapters/mongo"
	"log"
)

type ProductService struct {
	repo *mongo.ProductRepository
	proto.UnimplementedProductServiceServer
}

func NewProductService(mongoClient *mongo.MongoClient) *ProductService {
	repo := mongo.NewProductRepository(mongoClient)
	return &ProductService{repo: repo}
}

func (s *ProductService) Fetch(ctx context.Context, req *proto.FetchRequest) (*proto.FetchResponse, error) {
	log.Printf("Fetching from URL: %s", req.Url)

	return &proto.FetchResponse{Success: true}, nil
}

func (s *ProductService) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	products, err := s.repo.GetProducts(int(req.Page), int(req.PageSize), req.SortField, req.SortAsc)
	if err != nil {
		return nil, err
	}

	var protoProducts []*proto.Product
	for _, p := range products {
		protoProducts = append(protoProducts, &proto.Product{
			Name:    p.Name,
			Price:   float32(p.Price),
			Updates: int32(p.Updates),
		})
	}

	return &proto.ListResponse{Products: protoProducts}, nil
}
