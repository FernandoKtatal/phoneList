package database

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"postapi/app/models"
)

type PostDB interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	GetCollection(name string) *mongo.Collection
	Insert(collection string, p *models.Phones) error
	Select(collection string, country string, valid bool) (out *[]models.JsonPost, err error)
}

type DB struct {
	client *mongo.Client
	dbname string
}

func (d *DB) Connect(ctx context.Context) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return err
	}
	fmt.Println("Conectando Mongo")
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Ping")
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	d.dbname = Database
	return nil
}

func (d *DB) Disconnect(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}

func (d *DB) GetCollection(name string) *mongo.Collection {
	return d.client.Database(d.dbname).Collection(name)
}
