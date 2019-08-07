package author

import (
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/request"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/response"
	"github.com/PhantomX7/go-cleanarch/models"
)

type AuthorRepository interface {
	Insert(author *models.Author)  error
	Update(author *models.Author)  error
	FindAll(config request.PaginationConfig) ([]models.Author, response.AuthorPaginationMeta, error)
	FindByID(authorId int64) (models.Author, error)
}
