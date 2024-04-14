
# GET /v1/feeds

This endpoint retrieves a list of all feeds available in the system.

## Request

`GET /v1/feeds`

## Response

### Success Response

- **Code**: `200 OK`
- **Content**:

```json
[
  {
    "id": "uuid-here",
    "created_at": "timestamp-here",
    "updated_at": "timestamp-here",
    "name": "Feed Name",
    "url": "http://example.com/rss",
    "user_id": "uuid-here",
    "last_fetched_at": null
  },
  {
    "id": "uuid-here",
    "created_at": "timestamp-here",
    "updated_at": "timestamp-here",
    "name": "Another Feed",
    "url": "http://example.com/rss2",
    "user_id": "uuid-here",
    "last_fetched_at": null
  }
]
```

The response includes an array of all feeds, each represented as a JSON object. Each object includes the feed's ID, creation and update timestamps, name, URL, the ID of the user who created the feed, and the last time the feed was fetched.

### Error Responses

- **Internal Server Error (Feeds Retrieval Error)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Couldn't retrieve feeds"}`

## Examples
### Retrieve All Feeds

```bash
curl -X GET 'http://localhost:8080/v1/feeds'
```

## Notes

- The `last_fetched_at` field will be `null` if the feed has never been fetched.