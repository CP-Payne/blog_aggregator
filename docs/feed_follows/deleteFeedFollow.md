# DELETE /v1/feed_follows/{feedFollowID}

This endpoint allows an authenticated user to unfollow a specific feed that they are following.

## Request

`DELETE /v1/feed_follows/{feedFollowID}`

### URL Parameters

- `feedFollowID`: The UUID of the feed follow record to delete.

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
{   
	"success": "Feed unfollowed" 
}
```

### Error Responses

- **Bad Request (Invalid Feed ID)**:
    - **Code**: `400 BAD REQUEST`
    - **Content**: `{"error": "Invalid feed id provided"}`
- **Internal Server Error (Error Unfollowing Feed)**:
    - **Code**: `500 INTERNAL SERVER ERROR`
    - **Content**: `{"error": "Something went wrong"}`
- **Unauthorized (Invalid or No API Key Provided)**:
    - **Code**: `401 UNAUTHORIZED`
    - **Content**: `{"error": "Invalid ApiKey"}` or `{"error": "Couldn't retrieve user"}`

## Examples

### Unfollow a Feed

```bash
curl -X DELETE 'http://localhost:8080/v1/feed_follows/{feedFollowID}' \
-H 'Authorization: ApiKey <api-key>'
```

Replace `{feedFollowID}` with the actual UUID of the feed follow record and `<api-key>` with the actual API key.

## Notes

- The endpoint requires a valid `Authorization` header with an API key to ensure the user is authenticated.
- The user can only unfollow feeds that they are currently following, and they must provide a valid UUID of the feed follow record.
- If the provided `feedFollowID` is not a valid UUID, the server will respond with a `400 BAD REQUEST`.
- If there is an issue with the database operation, such as the feed follow not existing or the user not being authorized to delete this follow, the server will respond with a `500 INTERNAL SERVER ERROR`.