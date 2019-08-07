package mysql

import (
	"github.com/PhantomX7/go-cleanarch/author"
	"github.com/PhantomX7/go-cleanarch/models"
	"github.com/jinzhu/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) author.AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (a *AuthorRepository) Insert(author *models.Author) error {
	return a.db.Create(author).Error
}

func (a *AuthorRepository) FindAll(config author.PaginationConfig) ([]models.Author, author.AuthorPaginationMeta, error) {
	return nil, author.AuthorPaginationMeta{}, nil
}
func (a *AuthorRepository) FindByID(authorId uint64) (models.Author, error) {
	return models.Author{}, nil
}
