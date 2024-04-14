
# GET /v1/feed_follows

This endpoint retrieves a list of all feeds that the authenticated user is following.

## Request

`GET /v1/feed_follows`

### Headers

```plaintext
Authorization: ApiKey <api-key>
```

Replace `<api-key>` with the API key received upon authentication. The API key is used to authenticate the request.

## Response

### Success Response

- **Code**: `200 OK`
- **Content**:

```json
[
  {
    "id": "uuid-here",
    "feed_id": "uuid-here",
    "user_id": "uuid-here",
    "created_at": "timestamp-here",
    "updated_at": "timestamp-here"
  },
  {
    "id": "uuid-here",
    "feed_id": "uuid-here",
    "user_id": "uuid-here",
    "created_at": "timestamp-here",
    "updated_at": "timestamp-here"
  }
]
```
The response includes an array of feed follow records, showing each feed the user follows, with details such as the feed follow ID, feed ID, user ID, and the timestamps of creation and last update.

### Error Responses
If the query includes an invalid `author_id`, or if there is an internal error when retrieving chirps:

- **Internal Server Error (Failed to Retrieve Followed Feeds)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Couldn't retrieve followed feeds"}`
- **Unauthorized (Invalid or No API Key Provided)**:
    - **Code**: `401 UNAUTHORIZED`
    - **Content**: `{"error": "Invalid ApiKey"}` or `{"error": "Couldn't retrieve user"}`

## Examples
### Retrieve All Feeds Followed by User

```bash
curl -X GET 'http://localhost:8080/v1/feed_follows' \
-H 'Authorization: ApiKey <api-key>'
```
Replace `<api-key>` with the actual API key.

## Notes

- This endpoint requires a valid `Authorization` header with an API key. If the key is missing or invalid, or if there is an issue validating the user, the server will respond with a `401 UNAUTHORIZED`.
- If the user is not following any feeds or if there is an error retrieving the feed follows from the database, the server will respond appropriately, either with an empty list or an error message.