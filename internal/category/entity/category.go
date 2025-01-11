package categoryentity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type Category struct {
	Name           string      `json:"name"`
	SubCategory    []*Category `json:"subCategory,omitempty"`
	ParentCategory *Category   `json:"parentCategory,omitempty"`
	Icon           string      `json:"icon"`

	entity.BaseEntity
}

func NewCategory(name, icon string, parentCategory *Category) *Category {
	return &Category{
		Name:           name,
		ParentCategory: parentCategory,
		Icon:           icon,
	}
}
