package mongo

import (
	"context"
	"github.com/yukitsune/chameleon/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDataSource struct {
	db *mongo.Database
}

func NewMongoDataSource(db *mongo.Database) repository.DataSource {
	return &mongoDataSource{db}
}

func (ds *mongoDataSource) Collection(c string) repository.Repository {
	coll := ds.db.Collection(c)
	repo := &mongoRepository{coll}
	return repo
}

type mongoRepository struct {
	coll *mongo.Collection
}


func (r *mongoRepository) Add(ctx context.Context, doc interface{}) error {
	_, err := r.coll.InsertOne(ctx, doc)
	return err
}

func (r *mongoRepository) Count(ctx context.Context, filter repository.Filter) (int64, error) {
	return r.coll.CountDocuments(ctx, filterToM(filter))
}

func (r *mongoRepository) FindAll(ctx context.Context, filter repository.Filter, result interface{}) error {
	cur, err := r.coll.Find(ctx, filterToM(filter))
	if err != nil {
		return err
	}

	return cur.All(ctx, result)
}

func (r *mongoRepository) FindOne(ctx context.Context, filter repository.Filter, result interface{}) error {
	res := r.coll.FindOne(ctx, filterToM(filter))
	if res.Err() != nil {
		return res.Err()
	}

	return res.Decode(result)
}

func (r *mongoRepository) UpdateById(ctx context.Context, id interface{}, doc interface{}) error {
	_, err := r.coll.UpdateByID(ctx, id, doc)
	return err
}

func (r *mongoRepository) DeleteById(ctx context.Context, id interface{}) error {
	_, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func filterToM(filter repository.Filter) bson.M {
	return bson.M(filter)
}