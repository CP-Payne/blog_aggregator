
# POST /v1/feed_follows

This endpoint allows an authenticated user to follow a specific feed.

## Request

`POST /v1/feed_follows`

### Headers
```plaintext
Authorization: ApiKey <api-key>
Content-Type: application/json
```
Replace `<api-key>` with the API key received upon authentication.

### Body Parameters

- `feed_id`: A UUID representing the ID of the feed to be followed.

### Example Body

```json
{
  "feed_id": "123e4567-e89b-12d3-a456-426614174000"
}
```

## Response

### Success Response

- **Code**: `201 CREATED`
- **Content**:

```json
{   
  "id": "uuid-here",
  "feed_id": "123e4567-e89b-12d3-a456-426614174000",
  "user_id": "uuid-of-the-user",
  "created_at": "timestamp-here",
  "updated_at": "timestamp-here"
}
```
The response includes the `id` of the new feed follow record, along with the `feed_id`, `user_id`, and timestamps for creation and last updated.

### Error Responses

- **Bad Request (Parameter Decoding Error)**:
    - **Code**: `400 BAD REQUEST`
    - **Content**: `{"error": "Couldn't decode parameters"}`
- **Internal Server Error (Failed to Follow Feed)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Couldn't follow feed"}`
- **Unauthorized (Invalid or No API Key Provided)**:
    - **Code**: `401 UNAUTHORIZED`
    - **Content**: `{"error": "Invalid ApiKey"}` or `{"error": "Couldn't retrieve user"}`

## Examples

### Follow a Feed

```bash
curl -X POST 'http://localhost:8080/v1/feed_follows' \
-H 'Authorization: ApiKey <api-key>' \
-H 'Content-Type: application/json' \
-d '{"feed_id": "<feed-id>"}'
```

Replace `<api-key>` with the actual API key and `<feed-id>` with an existing feed id.

## Notes

- To follow a feed, the user must provide a valid `feed_id` in the JSON body of the request.
- This operation requires authentication using a valid API key, and the user ID is derived from the authenticated session associated with the API key.