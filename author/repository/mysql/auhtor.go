package mysql

import (
	"github.com/PhantomX7/go-cleanarch/author"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/request"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/response"
	"github.com/PhantomX7/go-cleanarch/models"
	"github.com/PhantomX7/go-cleanarch/util/errors"
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

func (a *AuthorRepository) Update(author *models.Author) error {
	return a.db.Save(author).Error
}

func (a *AuthorRepository) FindAll(config request.PaginationConfig) ([]models.Author, response.AuthorPaginationMeta, error) {
	return nil, response.AuthorPaginationMeta{}, nil
}
func (a *AuthorRepository) FindByID(authorID int64) (models.Author, error) {
	model := models.Author{}

	err := a.db.Where("id = ?", authorID).First(&model).Error

	if gorm.IsRecordNotFoundError(err) {
		return model, errors.ErrNotFound
	}

	return model, err
}
