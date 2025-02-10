package wordpressdriver

type Category struct {
	ID          int    `gorm:"column:term_id"`
	Name        string `gorm:"column:name"`
	Slug        string `gorm:"column:slug"`
	Description string `gorm:"column:description"`
}

func (c *Category) GetAllCategories(prefixTable string) ([]Category, error) {
	var categories []Category
	pTable := prefixTable + "_terms"

	result := DB.Table(pTable).Find(&categories)
	return categories, result.Error
}
