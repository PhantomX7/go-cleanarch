package usecase

import (
	"github.com/PhantomX7/go-cleanarch/author"
	"github.com/PhantomX7/go-cleanarch/models"
)

// apply business logic here

type AuthorUsecase struct {
	authorRepo author.AuthorRepository
}

func NewAuthorUsecase(authorRepo author.AuthorRepository) author.AuthorUsecase {
	return &AuthorUsecase{
		authorRepo: authorRepo,
	}
}

func (a *AuthorUsecase) Create(author models.Author) (models.Author, error) {
	err := a.authorRepo.Insert(&author)
	if err != nil {
		return author, err
	}
	return author, nil
}

func (a *AuthorUsecase) Index(paginationConfig author.PaginationConfig) ([]models.Author, author.AuthorPaginationMeta, error) {
	return nil, author.AuthorPaginationMeta{}, nil
}

func (a *AuthorUsecase) Show(authorID uint64) (models.Author, error) {
	return models.Author{}, nil
}
