package wordpressdriver

import "gorm.io/gorm"

type Page struct {
	ID          int    `gorm:"column:ID"`
	PostTitle   string `gorm:"column:post_title"`
	PostName    string `gorm:"column:post_name"`
	PostContent string `gorm:"column:post_content"`
	PostStatus  string `gorm:"column:post_status"`
	PostType    string `gorm:"column:post_type"`
}

func (w *WpFacade) Page() *Page {
	return &Page{}
}

func (p *Page) GetAllPages(prefixTable string) ([]Page, error) {
	var pages []Page
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_type = ?", "page").Find(&pages)
	return pages, result.Error
}

func (p *Page) GetPageByID(id int, prefixTable string) (*Page, error) {
	var page Page
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("ID = ? AND post_type = ?", id, "page").First(&page)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &page, result.Error
}

func (p *Page) SearchPages(keyword, prefixTable string) ([]Page, error) {
	var pages []Page
	pTable := prefixTable + "_posts"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("(post_title LIKE ? OR post_content LIKE ?) AND post_type = ?", query, query, "page").Find(&pages)
	return pages, result.Error
}

func (p *Page) AddPage(title, name, content, status, prefixTable string) error {
	pTable := prefixTable + "_posts"
	page := Page{PostTitle: title, PostName: name, PostContent: content, PostStatus: status, PostType: "page"}

	result := DB.Table(pTable).Create(&page)
	return result.Error
}

func (p *Page) UpdatePage(id int, title, name, content, status, prefixTable string) error {
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("ID = ? AND post_type = ?", id, "page").Updates(Page{PostTitle: title, PostName: name, PostContent: content, PostStatus: status})
	return result.Error
}

func (p *Page) DeletePage(id int, prefixTable string) error {
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("ID = ? AND post_type = ?", id, "page").Delete(&Page{})
	return result.Error
}
