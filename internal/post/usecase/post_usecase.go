package usecase

import (
	"context"
	"github.com/Russia9/its-bmstu-go-architecture/pkg/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

type PostUsecase struct {
	repo domain.PostRepository
}

func NewPostUsecase(repo domain.PostRepository) *PostUsecase {
	return &PostUsecase{repo: repo}
}

func (p *PostUsecase) Create(ctx context.Context, content string, attachments []string) (*domain.Post, error) {
	post := &domain.Post{
		ID:          uuid.NewString(),
		Content:     content,
		Attachments: attachments,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := p.repo.Create(ctx, post)
	if err != nil {
		return nil, errors.Wrap(err, "post repo")
	}

	return post, nil
}

func (p *PostUsecase) Get(ctx context.Context, id string) (*domain.Post, error) {
	post, err := p.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "post repo")
	}

	return post, nil
}

func (p *PostUsecase) Update(ctx context.Context, id string, content string, attachments []string) (*domain.Post, error) {
	post, err := p.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "post repo")
	}

	post.Content = content
	post.Attachments = attachments
	post.UpdatedAt = time.Now()

	err = p.repo.Update(ctx, post)
	if err != nil {
		return nil, errors.Wrap(err, "post repo")
	}

	return post, nil
}

func (p *PostUsecase) Delete(ctx context.Context, id string) error {
	return errors.Wrap(p.Delete(ctx, id), "post repo")
}