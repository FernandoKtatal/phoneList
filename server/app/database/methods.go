package database

import (
	"context"
	"postapi/app/models"
)

func (d *DB) Insert(collection string, p *models.Phones) error {
	_,err := d.GetCollection(collection).InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Select(collection string, country string, valid bool) (out *[]models.JsonPost, err error) {
	filter := models.Phones{
		Country:     country,
		State:       valid,
	}
	cur, err := d.GetCollection(collection).Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}
	cur.All(context.Background(), &out)

	return out, err
}
