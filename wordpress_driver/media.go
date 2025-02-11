package wordpressdriver

import "gorm.io/gorm"

type Media struct {
	ID        int    `gorm:"column:ID"`
	PostTitle string `gorm:"column:post_title"`
	PostName  string `gorm:"column:post_name"`
	Guid      string `gorm:"column:guid"`
	MimeType  string `gorm:"column:post_mime_type"`
	PostType  string `gorm:"column:post_type"`
}

func (w *WpFacade) Media() *Media {
	return &Media{}
}

func (m *Media) GetAllMedia(prefixTable string) ([]Media, error) {
	var media []Media
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_type = ?", "attachment").Find(&media)
	return media, result.Error
}

func (m *Media) GetMediaByID(id int, prefixTable string) (*Media, error) {
	var media Media
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("ID = ? AND post_type = ?", id, "attachment").First(&media)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &media, result.Error
}

func (m *Media) GetMediaByMimeType(mimeType, prefixTable string) ([]Media, error) {
	var media []Media
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_mime_type = ? AND post_type = ?", mimeType, "attachment").Find(&media)
	return media, result.Error
}

func (m *Media) SearchMedia(keyword, prefixTable string) ([]Media, error) {
	var media []Media
	pTable := prefixTable + "_posts"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("(post_title LIKE ? OR post_name LIKE ?) AND post_type = ?", query, query, "attachment").Find(&media)
	return media, result.Error
}
func (m *Media) DeleteMedia(id int, prefixTable string) error {
	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("ID = ? AND post_type = ?", id, "attachment").Delete(&Media{})
	return result.Error
}
