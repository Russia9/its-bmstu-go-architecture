package domain

import (
	"context"
	"errors"
	"time"
)

type Post struct {
	ID          string `bson:"_id"`
	Title       string
	Content     string
	Attachments []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

var ErrPostNotFound = errors.New("post not found")

type PostUsecase interface {
	Create(ctx context.Context, title string, content string, attachments []string) (*Post, error)
	Get(ctx context.Context, id string) (*Post, error)
	Update(ctx context.Context, id string, title string, content string, attachments []string) (*Post, error)
	Delete(ctx context.Context, id string) error
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	Get(ctx context.Context, id string) (*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id string) error
}
