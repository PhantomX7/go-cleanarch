package author

type AuthorPaginationMeta struct {
	Limit   int
	Offset  int
	Total   int
}

type PaginationConfig interface {
	Limit() int
	Offset() int
	Order() string
	SearchClause() map[string]interface{}
}