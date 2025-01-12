package types

type Pagination struct {
	Page  uint
	Limit uint
}

func NewPagination(page, limit uint) Pagination {
	return Pagination{
		Page:  page,
		Limit: limit,
	}
}
