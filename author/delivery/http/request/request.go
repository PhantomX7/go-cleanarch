package request

import (
	"fmt"
	"github.com/PhantomX7/go-cleanarch/models"
	"strconv"
)

// request related struct

type PaginationConfig interface {
	Limit() int
	Offset() int
	Order() string
	SearchClause() map[string]interface{}
}

type AuthorCreateRequest struct {
	Name  string `json:"name" form:"name" binding:"required,unique=authors.name"`
	Email string `json:"email" form:"email" binding:"required,email,unique=authors.email"`
	Age   int    `json:"age" form:"age" binding:"required,gte=18"`
}

func (request AuthorCreateRequest) ToAuthorModel() models.Author {
	return models.Author{
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}
}

type AuthorUpdateRequest struct {
	Name  *string `form:"name" binding:"omitempty,unique=authors.name"`
	Email *string `form:"email" binding:"omitempty,email,unique=authors.email"`
	Age   *int    `form:"age" binding:"omitempty,gte=18"`
}

type AuthorPaginationConfig struct {
	limit        int
	offset       int
	order        string
	searchClause map[string]interface{}
}

func NewAuthorPaginationConfig(conditions map[string][]string) AuthorPaginationConfig {
	authorPaginationConfig := AuthorPaginationConfig{
		limit:        buildLimit(conditions),
		offset:       buildOffset(conditions),
		order:        buildOrder(conditions),
		searchClause: buildSearchClause(conditions),
	}

	return authorPaginationConfig
}

func (a AuthorPaginationConfig) Limit() (res int) {
	return a.limit
}

func (a AuthorPaginationConfig) Order() string {
	return a.order
}

func (a AuthorPaginationConfig) Offset() (res int) {
	return a.offset
}

func (a AuthorPaginationConfig) SearchClause() (res map[string]interface{}) {
	return a.searchClause
}

func buildLimit(conditions map[string][]string) (res int) {
	if len(conditions["limit"]) > 0 {
		res, _ = strconv.Atoi(conditions["limit"][0])
	}

	return
}

func buildOffset(conditions map[string][]string) (res int) {
	if len(conditions["offset"]) > 0 {
		res, _ = strconv.Atoi(conditions["offset"][0])
	}
	return
}

func buildOrder(conditions map[string][]string) string {
	var order string
	if len(conditions["sort"]) > 0 {
		order = conditions["sort"][0]
	}

	orderCol, orderDir := "", ""

	if len(order) > 2 {
		if order[0:1] == "-" {
			orderDir = "desc"
			order = order[1:]
		} else {
			orderDir = "asc"
		}

		if order == "created_at" || order == "order" {
			orderCol = fmt.Sprintf("`%s`", order)
		}
		return orderCol + " " + orderDir
	}

	return ""
}

func buildSearchClause(conditions map[string][]string) (res map[string]interface{}) {
	res = make(map[string]interface{})
	if len(conditions["status"]) > 0 {
		//status := conditions["status"][0]
		//if status == "active" {
		//	res["status"] = mysql.ActiveBannerStatusCondition
		//
		//} else if status == "inactive" {
		//	res["status"] = mysql.InactiveBannerStatusCondition
		//}
	}

	if len(conditions["type"]) > 0 {
		//typeCode := models.BannerTypeString(conditions["type"][0]).ToCode()
		//if typeCode < models.BannerType(len(models.BannerTypeArray)) {
		//	res["type"] = models.BannerTypeString(conditions["type"][0]).ToCode()
		//}
	}

	if len(conditions["title"]) > 0 {
		res["title"] = conditions["title"][0]
	}

	return
}
