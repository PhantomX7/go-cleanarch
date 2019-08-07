package http

import (
	"fmt"
	"github.com/PhantomX7/go-cleanarch/models"
	"github.com/bukalapak/mitra/banner"
	"github.com/bukalapak/mitra/banner/repository/mysql"
	"strconv"
)

// this file is to handle request related and put the request struct here

// request related struct
type AuthorCreateRequest struct {
	Name  string `form:"name" binding:"required,unique=authors.name"`
	Email string `form:"email" binding:"required,email,unique=authors.email"`
}

type AuthorPaginationConfig struct {
	limit        int
	offset       int
	order        string
	searchClause map[string]interface{}
}

func NewAuthorModelFromReq(request AuthorCreateRequest) models.Author {
	return models.Author{
		Name:  request.Name,
		Email: request.Email,
	}
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
	if res > banner.MAX_LIMIT {
		res = banner.MAX_LIMIT
	}
	if res < banner.MIN_LIMIT {
		res = banner.MIN_LIMIT
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
		status := conditions["status"][0]
		if status == "active" {
			res["status"] = mysql.ActiveBannerStatusCondition

		} else if status == "inactive" {
			res["status"] = mysql.InactiveBannerStatusCondition
		}
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
