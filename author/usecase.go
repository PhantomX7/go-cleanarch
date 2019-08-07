package author

import (
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/request"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/response"
	"github.com/PhantomX7/go-cleanarch/models"
)

type AuthorUsecase interface {
	Create(author models.Author) (models.Author, error)
	Update(authorID int64, author request.AuthorUpdateRequest) (models.Author, error)
	Index(paginationConfig request.PaginationConfig) ([]models.Author, response.AuthorPaginationMeta, error)
	Show(authorID int64) (models.Author, error)
}
