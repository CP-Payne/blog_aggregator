
# POST /v1/feeds

This endpoint allows an authenticated user to create a new feed and automatically follows it.

## Request

`POST /v1/feeds`

### Headers
```plaintext
Authorization: ApiKey <api-key>
Content-Type: application/json
```
 

Replace `<api-key>` with the API key received upon account creation.

### Body Parameters

- `name`: A string containing the name of the feed.
- `url`: A string containing the URL of the feed.

### Example Body

```json
{
  "name": "Tech News",
  "url": "http://technews.example.com/rss"
}
```

## Response

### Success Response

- **Code**: `201 CREATED`
- **Content**:

```json
{   
	"feed": {
	    "id": "uuid-here",
	    "created_at": "timestamp-here",
	    "updated_at": "timestamp-here",
	    "name": "Tech News",
	    "url": "http://technews.example.com/rss",
	    "user_id": "uuid-here",
	    "last_fetched_at": null
	},
	"feed_follows": {
		"id": "uuid-here",
	    "feed_id": "uuid-here",
	    "user_id": "uuid-here",
	    "created_at": "timestamp-here",
	    "updated_at": "timestamp-here"
	}
}
```


### Error Responses

- **Bad Request (Parameter Decoding Error)**:
    - **Code**: `400 BAD REQUEST`
    - **Content**: `{"error": "Couldn't decode parameters"}`
- **Internal Server Error (Feed Creation Error)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Couldn't create feed"}`
- **Internal Server Error (Feed Follow Creation Error)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Something went wrong. Couldn't create follow feed"}`

## Examples

### Create and Automatically Follow a New Feed

```bash
curl -X POST 'http://localhost:8080/v1/feeds' \
-H 'Content-Type: application/json' \
-H 'Authorization: ApiKey <api-key>' \
-d '{"name": "Tech News", "url": "http://technews.example.com/rss"}'
```

Replace `<api-key>` with the actual API key.

## Notes

- The `name` and `url` fields in the body are required for creating a new feed.
- Upon successfully creating a feed, the system also automatically creates a follow record for the user, associating them with the new feed.
- Error responses can vary based on parameter decoding, feed creation, or feed follow creation issues, and appropriate error messages are provided for each scenario.