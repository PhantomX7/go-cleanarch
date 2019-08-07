package http

import (
	"github.com/PhantomX7/go-cleanarch/app/api/middleware"
	"github.com/PhantomX7/go-cleanarch/app/api/server"
	"github.com/PhantomX7/go-cleanarch/author"
	"github.com/PhantomX7/go-cleanarch/author/delivery/http/request"
	"github.com/PhantomX7/go-cleanarch/util/errors"
	"github.com/gin-gonic/gin"
	"strconv"
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
		authorRoute.GET("/:id", h.Show)
		authorRoute.POST("/", h.Create)
		authorRoute.PUT("/:id", h.Update)
	}
}

func (h *AuthorHandler) Create(c *gin.Context) {
	var req request.AuthorCreateRequest

	// validate request
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	authorModel, err := h.authorUsecase.Create(req.ToAuthorModel())
	if err != nil {
		_ = c.Error(err).SetType(errors.ErrorTypeUnprocessableEntity)
		return
	}

	c.JSON(200, authorModel)
}

func (h *AuthorHandler) Update(c *gin.Context) {
	var req request.AuthorUpdateRequest
	authorID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	// validate request
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	authorModel, err := h.authorUsecase.Update(authorID, req)
	if err != nil {
		_ = c.Error(err).SetType(errors.ErrorTypeUnprocessableEntity)
		return
	}

	c.JSON(200, authorModel)
}

func (h *AuthorHandler) Index(c *gin.Context) {
	c.JSON(200, "test")
}

func (h *AuthorHandler) Show(c *gin.Context) {
	authorID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	authorModel, err := h.authorUsecase.Show(authorID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, authorModel)
}
