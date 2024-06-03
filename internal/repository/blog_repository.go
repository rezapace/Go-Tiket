package repository
// NOTE :
// FOLDER INI UNTUK MENANGANI KE BAGIAN DATABASE DAN QUERY
import (
	"context"

	"Ticketing/entity"

	"gorm.io/gorm"
)

// Blog repository
type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

// GetAllBlogs retrieves all blogs from the database.
func (r *BlogRepository) GetAllBlogs(ctx context.Context) ([]*entity.Blog, error) {
	blogs := make([]*entity.Blog, 0)
	result := r.db.WithContext(ctx).Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

// CreateBlog saves a new Blog to the database.
func (r *BlogRepository) CreateBlog(ctx context.Context, Blog *entity.Blog) error {
	result := r.db.WithContext(ctx).Create(&Blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateBlog updates a Blog in the database.
func (r *BlogRepository) UpdateBlog(ctx context.Context, Blog *entity.Blog) error {
	result := r.db.WithContext(ctx).Model(&entity.Blog{}).Where("id = ?", Blog.ID).Updates(&Blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetBlog retrieves a Blog by its ID from the database.
func (r *BlogRepository) GetBlog(ctx context.Context, id int64) (*entity.Blog, error) {
	Blog := new(entity.Blog)
	result := r.db.WithContext(ctx).First(&Blog, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return Blog, nil
}

// DeleteBlog deletes a Blog from the database.
func (r *BlogRepository) DeleteBlog(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&entity.Blog{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SearchBlog search Blog
func (r *BlogRepository) SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error) {
	blogs := make([]*entity.Blog, 0)
	result := r.db.WithContext(ctx).Where("title LIKE ?", "%"+search+"%").Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}
