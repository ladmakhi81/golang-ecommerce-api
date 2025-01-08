package category_entity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type Category struct {
	Name           string
	SubCategory    []*Category
	ParentCategory *Category
	Icon           string

	entity.BaseEntity
}
