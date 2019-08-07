package author

import (
	"github.com/PhantomX7/go-cleanarch/models"
)

type AuthorRepository interface {
	Insert(author *models.Author)  error
	FindAll(config PaginationConfig) ([]models.Author, AuthorPaginationMeta, error)
	FindByID(authorId uint64) (models.Author, error)
}
