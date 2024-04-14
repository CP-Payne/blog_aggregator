# GET /v1/users

This endpoint retrieves the details of the authenticated user.

## Request

`GET /v1/users`

### Headers

```plaintext
Authorization: ApiKey <api-key>
```

Replace `<api-key>` with the API key provided to the user. The API key is used to authenticate the request.

## Response

### Success Response

- **Code**: `200 OK`
- **Content**: 

```json
{  
	"id": "uuid-here",
	"name": "John Doe",
	"created_at": "timestamp-here",
	"updated_at": "timestamp-here",
	"apikey": "apikey-here"
}
```

The response contains the unique identifier (`id`) for the user, their name (`name`), the timestamps for when the user account was created (`created_at`) and last updated (`updated_at`), as well as the user's API key (`apikey`).


### Error Responses

- **Unauthorized (Invalid or No API Key Provided)**:
    - **Code**: `401 UNAUTHORIZED`
    - **Content**: `{"error": "Invalid ApiKey"}` or `{"error": "Couldn't retrieve user"}`

## Examples

### Retrieve Authenticated User Details

To update a user's email and password, you would make the following request:

```json
curl -X GET 'http://localhost:8080/v1/users' \
-H 'Authorization: ApiKey <api-key>'
```
Replace `<api-key>` with the actual API key provided to the user.

## Notes

- The request must include a valid `Authorization` header with an API key. If the key is missing, invalid, or if the user cannot be retrieved, the server will respond with a `401 UNAUTHORIZED`.
- The `apikey` included in the user's details is intended for use in authenticating subsequent API requests that require it.
- This endpoint will only return the details of the user associated with the provided API key.