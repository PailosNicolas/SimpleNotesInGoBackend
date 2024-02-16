# SimpleNotesInGoBackend
This is a small project for getting used to work with go webservers, just a simple note taking backend.
I wanted to learn how to user authenticate and make a simple CRUD

## Endpoints
### `POST /users/`:
Creates a new user.

Body:
```json
{
  "username": "user",
  "password": "password"
}
```

Response:
```json
{
  "id": "b03e09ae-8582-42ea-a95c-1f2c4bd35802",
  "created_at": "2024-02-16T21:16:12.695386Z",
  "updated_at": "2024-02-16T21:16:12.695386Z",
  "username": "user"
}
```

### `POST /login/`:
Log in with user and password, returns access and refresh JWT tokens.

Body:
```json
{
  "username": "user",
  "password": "password"
}
```

Response:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyZWZyZXNoIiwic3ViIjoiYjAzZTA5YWUtODU4Mi00MmVhLWE5NWMtMWYyYzRiZDM1ODAyIiwiZXhwIjoxNzEzMzAyMzMzLCJpYXQiOjE3MDgxMTgzMzN9.h-KHsx_a89Qa7AL7_5fc5NLwt3C_7ySDh7WZjo_ZSUY"
}
```

### `POST /notes/`:
Creates a new note for the user that its loged in, auth via `Bearer token` using the `access_token`.

Authentication: Bearer Token
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg
```

Body:
```json
{
  "title": "Note title",
  "body": "This is the note body"
}
```

Response:
```json
{
  "id": "90e496c5-d0d1-4174-ae6b-96a362409397",
  "title": "Note title",
  "body": "This is the note body",
  "created_at": "2024-02-16T21:22:16.994501Z",
  "updated_at": "2024-02-16T21:22:16.994502Z"
}
```

### `PUT /notes/`:
Updates a note via ID, if `title` or `body` is `""` that field will remain the same, auth via `Bearer token` using the `access_token`.

Authentication: Bearer Token
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg
```

Body:
```json
{
  "id": "90e496c5-d0d1-4174-ae6b-96a362409397",
  "title": "",
  "body": "This is the new note body"
}
```

Response:
```json
{
  "id": "90e496c5-d0d1-4174-ae6b-96a362409397",
  "title": "Note title",
  "body": "This is the new note body",
  "created_at": "2024-02-16T21:22:16.994501Z",
  "updated_at": "2024-02-16T21:24:11.729657Z"
}
```