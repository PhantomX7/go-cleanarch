package request

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// QueryHelper represent helper to get query string data
type QueryHelper struct {
	r  *http.Request
	uv url.Values
}

// NewQueryHelper is a function to create query helper struct
func NewQueryHelper(r *http.Request) *QueryHelper {
	return &QueryHelper{r, r.URL.Query()}
}

// GetString to get query string value with string data type, return defValue if query url not found
func (q *QueryHelper) GetString(p string, defValue string) string {
	sv := q.uv.Get(p)
	if sv != "" {
		return sv
	}
	return defValue
}

// GetInt to get query string value with integer data type, return defValue if query url not found or invalid
func (q *QueryHelper) GetInt(p string, defValue int) int {
	sv := q.uv.Get(p)
	if sv != "" {
		if v, err := strconv.Atoi(sv); err == nil {
			return v
		}
	}
	return defValue
}

// GetFloat to get query string value with float data type, return defValue if query url not found or invalid
func (q *QueryHelper) GetFloat(p string, defValue float64) float64 {
	sv := q.uv.Get(p)
	if sv != "" {
		if v, err := strconv.ParseFloat(sv, 64); err == nil {
			return v
		}
	}
	return defValue
}

// GetBool to get query string value with boolean data type, return defValue if query url not found or invalid
func (q *QueryHelper) GetBool(p string, defValue bool) bool {
	sv := q.uv.Get(p)
	if sv != "" {
		v, err := strconv.ParseBool(sv)
		if err != nil {
			return defValue
		}
		return v
	}
	return defValue
}

// GetDate to get query string value with date data type using RFC3339 to parse string into time.Time, return defValue if query url not found or invalid
func (q *QueryHelper) GetDate(p string, defValue *time.Time) *time.Time {
	sv := q.uv.Get(p)
	if sv != "" {
		v, err := time.Parse(time.RFC3339, sv)
		if err != nil {
			return defValue
		}

		return &v
	}

	return defValue
}
