package wordpressdriver

import "gorm.io/gorm"

type Comment struct {
	ID          int    `gorm:"column:comment_ID"`
	PostID      int    `gorm:"column:comment_post_ID"`
	Author      string `gorm:"column:comment_author"`
	AuthorEmail string `gorm:"column:comment_author_email"`
	AuthorURL   string `gorm:"column:comment_author_url"`
	AuthorIP    string `gorm:"column:comment_author_IP"`
	Date        string `gorm:"column:comment_date"`
	DateGmt     string `gorm:"column:comment_date_gmt"`
	Content     string `gorm:"column:comment_content"`
	Karma       int    `gorm:"column:comment_karma"`
	Approved    string `gorm:"column:comment_approved"`
	Agent       string `gorm:"column:comment_agent"`
	Type        string `gorm:"column:comment_type"`
	ParentID    int    `gorm:"column:comment_parent"`
	UserID      int    `gorm:"column:user_id"`
}

func (w *WpFacade) Comment() *Comment {
	return &Comment{}
}

func (c *Comment) GetAllComments(prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Find(&comments)
	return comments, result.Error
}

func (c *Comment) GetCommentByID(id int, prefixTable string) (*Comment, error) {
	var comment Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_ID = ?", id).First(&comment)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &comment, result.Error
}

func (c *Comment) GetCommentsByPostID(postID int, prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_post_ID = ? AND comment_approved = ?", postID, "1").Find(&comments)
	return comments, result.Error
}

func (c *Comment) SearchComments(keyword, prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("comment_content LIKE ?", query).Find(&comments)
	return comments, result.Error
}

func (c *Comment) GetCommentsWithUserEmail(email, prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_author_email = ?", email).Find(&comments)
	return comments, result.Error
}

func (c *Comment) GetApprovedComments(prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_approved = ?", "1").Find(&comments)
	return comments, result.Error
}

func (c *Comment) GetRecentComments(limit int, prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Order("comment_date DESC").Limit(limit).Find(&comments)
	return comments, result.Error
}

func (c *Comment) GetPendingComments(prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_approved = ?", "0").Find(&comments)
	return comments, result.Error
}

func (c *Comment) ApproveComment(id int, prefixTable string) error {
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_ID = ?", id).Update("comment_approved", "1")
	return result.Error
}

func (c *Comment) DeleteComment(id int, prefixTable string) error {
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_ID = ?", id).Delete(&Comment{})
	return result.Error
}

func (c *Comment) GetCommentsByAuthorEmail(email, prefixTable string) ([]Comment, error) {
	var comments []Comment
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_author_email = ?", email).Find(&comments)
	return comments, result.Error
}

func (c *Comment) CountCommentsByPostID(postID int, prefixTable string) (int64, error) {
	var count int64
	pTable := prefixTable + "_comments"

	result := DB.Table(pTable).Where("comment_post_ID = ?", postID).Count(&count)
	return count, result.Error
}
