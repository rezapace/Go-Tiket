package service

import (
	"context"

	"Ticketing/entity"
)

// BlogUseCase is an interface for Blog-related use cases.
type BlogUseCase interface {
	GetAllBlogs(ctx context.Context) ([]*entity.Blog, error)
	CreateBlog(ctx context.Context, Blog *entity.Blog) error
	GetBlog(ctx context.Context, id int64) (*entity.Blog, error)
	UpdateBlog(ctx context.Context, Blog *entity.Blog) error
	SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
}

type BlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]*entity.Blog, error)
	CreateBlog(ctx context.Context, Blog *entity.Blog) error
	GetBlog(ctx context.Context, id int64) (*entity.Blog, error)
	UpdateBlog(ctx context.Context, Blog *entity.Blog) error
	SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
}

// BlogService is responsible for Blog-related business logic.
type BlogService struct {
	Repository BlogRepository
}

// NewBlogService creates a new instance of BlogService.
func NewBlogService(Repository BlogRepository) *BlogService {
	return &BlogService{Repository: Repository}
}

func (s *BlogService) GetAllBlogs(ctx context.Context) ([]*entity.Blog, error) {
	return s.Repository.GetAllBlogs(ctx)
}

func (s *BlogService) CreateBlog(ctx context.Context, Blog *entity.Blog) error {
	return s.Repository.CreateBlog(ctx, Blog)
}

func (s *BlogService) UpdateBlog(ctx context.Context, Blog *entity.Blog) error {
	return s.Repository.UpdateBlog(ctx, Blog)
}

func (s *BlogService) GetBlog(ctx context.Context, id int64) (*entity.Blog, error) {
	return s.Repository.GetBlog(ctx, id)
}

func (s *BlogService) DeleteBlog(ctx context.Context, id int64) error {
	return s.Repository.DeleteBlog(ctx, id)
}

// search Blog
func (s *BlogService) SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error) {
	return s.Repository.SearchBlog(ctx, search)
}