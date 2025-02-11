package wordpressdriver

import "gorm.io/gorm"

type Category struct {
	ID          int    `gorm:"column:term_id"`
	Name        string `gorm:"column:name"`
	Slug        string `gorm:"column:slug"`
	Description string `gorm:"column:description"`
}

func (w *WpFacade) Category() *Category {
	return &Category{}
}

func (c *Category) GetAllCategories(prefixTable string) ([]Category, error) {
	var categories []Category
	pTable := prefixTable + "_terms"

	result := DB.Table(pTable).Find(&categories)
	return categories, result.Error
}

func (c *Category) GetCategoryByID(id int, prefixTable string) (*Category, error) {
	var category Category
	pTable := prefixTable + "_terms"

	result := DB.Table(pTable).Where("term_id = ?", id).First(&category)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &category, result.Error
}

func (c *Category) GetCategoryBySlug(slug string, prefixTable string) (*Category, error) {
	var category Category
	pTable := prefixTable + "_terms"

	result := DB.Table(pTable).Where("slug = ?", slug).First(&category)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &category, result.Error
}

func (c *Category) SearchCategories(keyword, prefixTable string) ([]Category, error) {
	var categories []Category
	pTable := prefixTable + "_terms"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("name LIKE ? OR description LIKE ?", query, query).Find(&categories)
	return categories, result.Error
}

func (c *Category) GetPostsByCategoryID(categoryID int, prefixTable string) ([]Post, error) {
	var posts []Post
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_status = ? AND post_category = ?", "publish", categoryID).Find(&posts)
	return posts, result.Error
}

func (c *Category) GetCategoriesWithPostCount(prefixTable string) ([]Category, error) {
	var categories []Category
	pTable := prefixTable + "_terms"

	result := DB.Table(pTable).
		Select("wp_terms.term_id, wp_terms.name, COUNT(wp_posts.ID) as post_count").
		Joins("LEFT JOIN wp_posts ON wp_posts.post_category = wp_terms.term_id").
		Where("wp_posts.post_status = ?", "publish").
		Group("wp_terms.term_id").
		Scan(&categories)

	return categories, result.Error
}

func (c *Category) GetCategoriesByPostID(postID int, prefixTable string) ([]Category, error) {
	var categories []Category
	pTable := prefixTable + "_terms"

	result := DB.Table(pTable).
		Joins("INNER JOIN wp_term_taxonomy ON wp_term_taxonomy.term_id = wp_terms.term_id").
		Joins("INNER JOIN wp_term_relationships ON wp_term_relationships.term_taxonomy_id = wp_term_taxonomy.term_taxonomy_id").
		Where("wp_term_relationships.object_id = ?", postID).
		Find(&categories)

	return categories, result.Error
}
