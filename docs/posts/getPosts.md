
# GET /v1/posts

This endpoint retrieves posts from the RSS feeds that the authenticated user is subscribed to, optionally limited by a specified number.

## Request

`GET /v1/posts`

### Query Parameters

- `limit`: An optional integer specifying the maximum number of posts to retrieve. If not provided, defaults to 10.

### Headers

```plaintext
Authorization: ApiKey <api-key>
```

Replace `<api-key>` with the API key received upon authentication.

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
    "title": "Post Title",
    "url": "http://example.com/post",
    "description": "Post description",
    "published_at": "timestamp-here",
    "feed_id": "uuid-here"
  },
  // More posts
]
```
The response includes a list of posts with details such as ID, creation and update timestamps, title, URL, description, publication timestamp, and the ID of the feed from which the post originates.

### Error Responses

If the query includes an invalid `author_id`, or if there is an internal error when retrieving chirps:

- **Bad Request (Invalid Limit Parameter)**:
    - **Code**: `400 BAD REQUEST`
    - **Content**: `{"error": "Invalid limit parameter provided"}`
- **Internal Server Error (Error Retrieving Posts)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Couldn't retrieve latest posts"}`
- **Unauthorized (Invalid or No API Key Provided)**:
    - **Code**: `401 UNAUTHORIZED`
    - **Content**: `{"error": "Invalid ApiKey"}` or `{"error": "Couldn't retrieve user"}`
     
## Examples
### Retrieve Latest Posts

```bash
curl -X GET 'http://localhost:8080/v1/posts?limit=5' \
-H 'Authorization: ApiKey <api-key>'
```
Replace `<api-key>` with the actual API key.

## Notes
- This endpoint is secured and requires a valid `Authorization` header with an API key.
- The `limit` query parameter can be used to control how many posts are returned. If `limit` is not specified, the endpoint defaults to returning 10 posts.
- The API will return error responses if the limit parameter is incorrectly formatted or if there is a server error while fetching the posts.