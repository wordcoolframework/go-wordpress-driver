package queries

const GetPostsByCategory = `
	SELECT wp_posts.* 
	FROM wp_posts 
	INNER JOIN wp_term_relationships 
	ON wp_posts.ID = wp_term_relationships.object_id 
	WHERE wp_term_relationships.term_taxonomy_id = ? 
	AND wp_posts.post_status = 'publish'
`

const GetMostCommentedPosts = `
	SELECT wp_posts.*, COUNT(wp_comments.comment_ID) AS comment_count
	FROM wp_posts
	LEFT JOIN wp_comments ON wp_posts.ID = wp_comments.comment_post_ID
	WHERE wp_posts.post_status = 'publish'
	GROUP BY wp_posts.ID
	ORDER BY comment_count DESC
	LIMIT ?
`
const GetPostsByTag = `
	SELECT wp_posts.* 
	FROM wp_posts
	INNER JOIN wp_term_relationships 
	ON wp_posts.ID = wp_term_relationships.object_id
	WHERE wp_term_relationships.term_taxonomy_id = ? 
	AND wp_posts.post_status = 'publish'
`
