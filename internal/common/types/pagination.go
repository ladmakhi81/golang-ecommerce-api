package types

import "math"

type Pagination struct {
	Page  uint
	Limit uint
}

type PaginationResponse struct {
	TotalPage   uint `json:"totalPage"`
	TotalCount  uint `json:"totalCount"`
	CurrentPage uint `json:"currentPage"`
	Rows        any  `json:"rows"`
}

func NewPagination(page, limit uint) Pagination {
	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

func NewPaginationResponse(
	totalCount uint,
	pagination Pagination,
	rows any,
) PaginationResponse {
	return PaginationResponse{
		TotalPage:   uint(math.Ceil(float64(totalCount) / float64(pagination.Limit))),
		TotalCount:  totalCount,
		Rows:        rows,
		CurrentPage: pagination.Page + 1,
	}
}
