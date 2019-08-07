package usecase

import (
	"github.com/PhantomX7/go-cleanarch/author"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/request"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/response"
	"github.com/PhantomX7/go-cleanarch/models"
	"github.com/jinzhu/copier"
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

func (a *AuthorUsecase) Update(authorID int64, author request.AuthorUpdateRequest) (models.Author, error) {
	authorM, err := a.authorRepo.FindByID(authorID)
	if err != nil {
		return authorM, err
	}

	// copy content of request into author model found by id
	_ = copier.Copy(&authorM, &author)

	err = a.authorRepo.Update(&authorM)
	if err != nil {
		return authorM, err
	}
	return authorM, nil
}

func (a *AuthorUsecase) Index(paginationConfig request.PaginationConfig) ([]models.Author, response.AuthorPaginationMeta, error) {
	return nil, response.AuthorPaginationMeta{}, nil
}

func (a *AuthorUsecase) Show(authorID int64) (models.Author, error) {
	return a.authorRepo.FindByID(authorID)
}
