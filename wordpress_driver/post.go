package wordpressdriver

import (
	"fmt"
	"go-wordpress-driver/wordpress_driver/queries"

	"gorm.io/gorm"
)

type Post struct {
	ID      int    `gorm:"column:ID"`
	Title   string `gorm:"column:post_title"`
	Content string `gorm:"column:post_content"`
	Status  string `gorm:"column:post_status"`
}

func (p *Post) GetPosts(prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_status = ?", "publish").Find(&posts)

	return posts, result.Error
}

func (p *Post) GetPostByID(id int, prefixTable string) (*Post, error) {
	var post Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("id = ?", id).First(&post)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &post, result.Error
}

func (p *Post) GetPostByCategoryId(categoryID int) ([]Post, error) {
	var posts []Post

	result := DB.Raw(queries.GetPostsByCategory, categoryID).Scan(&posts)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return posts, result.Error
}

func (p *Post) GetLatestPosts(limit int, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_status = ?", "publish").Order("post_date DESC").Limit(limit).Find(&posts)

	return posts, result.Error
}

func (p *Post) GetPostsByAuthor(authorID int, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_status = ? AND post_author = ?", "publish", authorID).Find(&posts)

	return posts, result.Error
}

func (p *Post) CountPublishedPosts(prefixTable string) (int64, error) {
	var count int64

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Model(&Post{}).Where("post_status = ?", "publish").Count(&count)

	return count, result.Error
}

func (p *Post) SearchPosts(keyword string, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("post_status = ? AND (post_title LIKE ? OR post_content LIKE ?)", "publish", query, query).Find(&posts)

	return posts, result.Error
}

func (p *Post) GetMostCommentedPosts(limit int, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Raw(queries.GetMostCommentedPosts, limit).Scan(&posts)

	return posts, result.Error
}

func (p *Post) GetPostsByTag(tagID int, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Raw(queries.GetPostsByTag, tagID).Scan(&posts)
	return posts, result.Error
}

func (p *Post) CreatePost(post *Post, prefixTable string) (Post, error) {

	var CurrentPost Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Create(post).Scan(&CurrentPost)

	return CurrentPost, result.Error
}

func (p *Post) GetPostsByDateRange(startDate, endDate, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_status = ? AND post_date BETWEEN ? AND ?", "publish", startDate, endDate).Find(&posts)

	return posts, result.Error
}

func (p *Post) GetPostsByAuthorAndDateRange(authorID int, startDate, endDate, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Where("post_status = ? AND post_author = ? AND post_date BETWEEN ? AND ?", "publish", authorID, startDate, endDate).Find(&posts)

	return posts, result.Error
}

func (p *Post) CountPostsByAuthor(authorID int, prefixTable string) (int64, error) {
	var count int64

	pTable := prefixTable + "_posts"

	result := DB.Table(pTable).Model(&Post{}).Where("post_status = ? AND post_author = ?", "publish", authorID).Count(&count)

	return count, result.Error
}

func (p *Post) SearchPostsWithFilters(keyword, status, categoryID, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("post_status = ? AND (post_title LIKE ? OR post_content LIKE ?) AND post_category = ?", status, query, query, categoryID).Find(&posts)

	return posts, result.Error
}

func (p *Post) AdvancedSearch(keyword, status, startDate, endDate, authorID, categoryID, prefixTable string) ([]Post, error) {
	var posts []Post

	pTable := prefixTable + "_posts"
	query := "%" + keyword + "%"

	// ساخت Query با فیلترهای مختلف
	queryBuilder := DB.Table(pTable).Where("post_title LIKE ? OR post_content LIKE ?", query, query)

	if status != "" {
		queryBuilder = queryBuilder.Where("post_status = ?", status)
	}

	if startDate != "" && endDate != "" {
		queryBuilder = queryBuilder.Where("post_date BETWEEN ? AND ?", startDate, endDate)
	}

	if authorID != "" {
		queryBuilder = queryBuilder.Where("post_author = ?", authorID)
	}

	if categoryID != "" {
		queryBuilder = queryBuilder.Where("post_category = ?", categoryID)
	}

	// اجرای query
	result := queryBuilder.Find(&posts)
	return posts, result.Error
}

func (p *Post) GetPostStatistics(prefixTable string) (map[string]int64, error) {
	var statistics = make(map[string]int64)

	pTable := prefixTable + "_posts"

	// شمارش تعداد پست‌ها بر اساس وضعیت
	var publishedCount int64
	result := DB.Table(pTable).Where("post_status = ?", "publish").Count(&publishedCount)
	if result.Error != nil {
		return nil, result.Error
	}
	statistics["published"] = publishedCount

	var draftCount int64
	result = DB.Table(pTable).Where("post_status = ?", "draft").Count(&draftCount)
	if result.Error != nil {
		return nil, result.Error
	}
	statistics["draft"] = draftCount

	// شمارش تعداد پست‌ها بر اساس نویسنده
	var authorStats []struct {
		AuthorID int
		Count    int64
	}
	result = DB.Table(pTable).Select("post_author, COUNT(*) as count").Group("post_author").Scan(&authorStats)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, stat := range authorStats {
		statistics[fmt.Sprintf("author_%d", stat.AuthorID)] = stat.Count
	}

	return statistics, nil
}
