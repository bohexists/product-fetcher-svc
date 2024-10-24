package mongo

import (
	"context"
	"time"

	"github.com/bohexists/product-fetcher-svc/internal/app"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductRepository struct {
	client *MongoClient
}

func NewProductRepository(client *MongoClient) *ProductRepository {
	return &ProductRepository{client: client}
}

func (r *ProductRepository) InsertProduct(product *app.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.client.db.Collection("products").InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) GetProducts(page, pageSize int, sortField string, sortAsc bool) ([]app.Product, error) {
	var products []app.Product
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.client.db.Collection("products").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product app.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
