-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (url) DO NOTHING;


-- name: GetPostsByUser :many
SELECT 
    p.id,
    p.title,
    p.url,
    p.description,
    p.published_at,
    p.created_at,
    p.updated_at,
    p.feed_id
FROM 
    posts p
INNER JOIN 
    feed_follows ff ON p.feed_id = ff.feed_id
WHERE 
    ff.user_id = $1
ORDER BY 
    p.published_at DESC, p.created_at DESC
LIMIT 
    $2;
