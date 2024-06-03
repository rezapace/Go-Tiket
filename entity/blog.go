package entity

import (
	"time"
)

type Blog struct {
    ID               int64
    Image            string
    Date             string // Format: YYYY-MM-DD
    Title            string
    Description      string
    CreatedAt        time.Time
    UpdatedAt        time.Time
    DeletedAt        time.Time
}

func NewBlog(image, date, title, description string) *Blog {
    return &Blog{
        Image:       image,
        Date:        date,
        Title:       title,
        Description: description,
        CreatedAt:   time.Now(),
    }
}

func UpdateBlog(id int64, image, date, title, description string) *Blog {
    return &Blog{
        ID:          id,
        Image:       image,
        Date:        date,
        Title:       title,
        Description: description,
        UpdatedAt:   time.Now(),
    }
}
