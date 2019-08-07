package author

import "github.com/PhantomX7/go-cleanarch/models"

type AuthorUsecase interface {
	Create(author models.Author) (models.Author, error)
	Index(paginationConfig PaginationConfig) ([]models.Author, AuthorPaginationMeta, error)
	Show(authorID uint64) (models.Author, error)
}
