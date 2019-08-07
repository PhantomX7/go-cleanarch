package http

import (
	"github.com/PhantomX7/go-cleanarch/app/api/middleware"
	"github.com/PhantomX7/go-cleanarch/app/api/server"
	"github.com/PhantomX7/go-cleanarch/author"
	"github.com/PhantomX7/go-cleanarch/util/errors"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorUsecase author.AuthorUsecase
}

func NewAuthorHandler(authorUC author.AuthorUsecase) server.Handler {
	return &AuthorHandler{
		authorUsecase: authorUC,
	}
}

func (h *AuthorHandler) Register(r *gin.Engine, m *middleware.Middleware) {



	authorRoute := r.Group("/author")
	{
		authorRoute.GET("/", h.Index)
		authorRoute.POST("/", h.Create)
	}
}

func (h *AuthorHandler) Create(c *gin.Context) {
	var request AuthorCreateRequest

	// validate request
	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	authorModel, err := h.authorUsecase.Create(NewAuthorModelFromReq(request))
	if err != nil {
		_ = c.Error(err).SetType(errors.ErrorTypeUnprocessableEntity)
		return
	}

	c.JSON(200, authorModel)
}

func (h *AuthorHandler) Index(c *gin.Context) {

	c.JSON(200, "test")
}


