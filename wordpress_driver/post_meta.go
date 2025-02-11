package wordpressdriver

import "gorm.io/gorm"

type PostMeta struct {
	ID        int    `gorm:"column:meta_id"`
	PostID    int    `gorm:"column:post_id"`
	MetaKey   string `gorm:"column:meta_key"`
	MetaValue string `gorm:"column:meta_value"`
}

func (w *WpFacade) PostMeta() *PostMeta {
	return &PostMeta{}
}

func (p *PostMeta) GetAllPostMeta(prefixTable string) ([]PostMeta, error) {
	var postMeta []PostMeta
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Find(&postMeta)
	return postMeta, result.Error
}

func (p *PostMeta) GetPostMetaByID(id int, prefixTable string) (*PostMeta, error) {
	var postMeta PostMeta
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Where("meta_id = ?", id).First(&postMeta)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &postMeta, result.Error
}

func (p *PostMeta) GetPostMetaByPostID(postID int, prefixTable string) ([]PostMeta, error) {
	var postMeta []PostMeta
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Where("post_id = ?", postID).Find(&postMeta)
	return postMeta, result.Error
}

func (p *PostMeta) GetPostMetaByKey(postID int, key, prefixTable string) (*PostMeta, error) {
	var postMeta PostMeta
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Where("post_id = ? AND meta_key = ?", postID, key).First(&postMeta)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &postMeta, result.Error
}

func (p *PostMeta) SearchPostMeta(keyword, prefixTable string) ([]PostMeta, error) {
	var postMeta []PostMeta
	pTable := prefixTable + "_postmeta"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("meta_value LIKE ?", query).Find(&postMeta)
	return postMeta, result.Error
}

func (p *PostMeta) AddPostMeta(postID int, key, value, prefixTable string) error {
	pTable := prefixTable + "_postmeta"
	postMeta := PostMeta{PostID: postID, MetaKey: key, MetaValue: value}

	result := DB.Table(pTable).Create(&postMeta)
	return result.Error
}

func (p *PostMeta) UpdatePostMeta(postID int, key, value, prefixTable string) error {
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Where("post_id = ? AND meta_key = ?", postID, key).Update("meta_value", value)
	return result.Error
}

func (p *PostMeta) DeletePostMeta(postID int, key, prefixTable string) error {
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Where("post_id = ? AND meta_key = ?", postID, key).Delete(&PostMeta{})
	return result.Error
}

func (p *PostMeta) GetPostsByMetaKeyValue(key, value, prefixTable string) ([]int, error) {
	var postIDs []int
	pTable := prefixTable + "_postmeta"

	result := DB.Table(pTable).Select("post_id").Where("meta_key = ? AND meta_value = ?", key, value).Find(&postIDs)
	return postIDs, result.Error
}

func (p *PostMeta) GetPostMetaWithPostTitle(postID int, prefixTable string) (*struct {
	Title     string
	MetaKey   string
	MetaValue string
}, error) {
	var resultData struct {
		Title     string
		MetaKey   string
		MetaValue string
	}
	pTable := prefixTable + "_postmeta"
	postsTable := prefixTable + "_posts"

	result := DB.Table(pTable).
		Select("wp_posts.post_title as title, wp_postmeta.meta_key, wp_postmeta.meta_value").
		Joins("INNER JOIN "+postsTable+" ON wp_posts.ID = wp_postmeta.post_id").
		Where("wp_postmeta.post_id = ?", postID).
		Scan(&resultData)

	return &resultData, result.Error
}
