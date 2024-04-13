-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY
    CASE
        WHEN last_fetched_at IS NULL THEN 0
        ELSE 1
    END,
    last_fetched_at ASC
LIMIT $1;



-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = Now(),
    updated_at = Now()
WHERE id = $1;
