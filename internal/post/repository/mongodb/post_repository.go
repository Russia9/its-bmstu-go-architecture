package mongodb

import (
	"context"
	"github.com/Russia9/its-bmstu-go-architecture/pkg/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{db.Collection("posts")}
}

func (p *PostRepository) Create(ctx context.Context, post *domain.Post) error {
	_, err := p.collection.InsertOne(ctx, post)
	return errors.Wrap(err, "mongo")
}

func (p *PostRepository) Get(ctx context.Context, id string) (*domain.Post, error) {
	result := p.collection.FindOne(ctx, bson.M{"_id": id})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrPostNotFound
	} else if result.Err() != nil {
		return nil, errors.Wrap(result.Err(), "mongo")
	}

	var post domain.Post
	err := result.Decode(&post)
	if err != nil {
		return nil, errors.Wrap(result.Err(), "mongo")
	}

	return &post, nil
}

func (p *PostRepository) Update(ctx context.Context, post *domain.Post) error {
	result := p.collection.FindOneAndReplace(ctx, bson.M{"_id": post.ID}, post)
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return domain.ErrPostNotFound
	} else if result.Err() != nil {
		return errors.Wrap(result.Err(), "mongo")
	}

	return nil
}

func (p *PostRepository) Delete(ctx context.Context, id string) error {
	result := p.collection.FindOneAndDelete(ctx, bson.M{"_id": id})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return domain.ErrPostNotFound
	} else if result.Err() != nil {
		return errors.Wrap(result.Err(), "mongo")
	}

	return nil
}
