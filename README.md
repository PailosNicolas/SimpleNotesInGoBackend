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

### `GET /notes/`:
Gets every note of the user via `access_token` paginated, can filter by category uuid, auth via `Bearer token` using the `access_token`.

Authentication: Bearer Token
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg
```

Body:
```json
{
  "page": 1,
  "page_size": 2,
  "filter_by_category_uuid": ["96363d02-c080-40a2-b454-105e3f755d5b"]
}
```

Response:
```json
{
  "total": 2,
  "page": 1,
  "page_size": 5,
  "items": [
    {
      "note": {
        "id": "e41d9e0d-9b2f-43be-971b-bcfea9180068",
        "title": "Testing note",
        "body": "This is a note for testing",
        "created_at": "2024-02-26T23:00:40.800129Z",
        "updated_at": "2024-02-26T23:00:40.800129Z"
      },
      "categories": [
        {
          "id": "96363d02-c080-40a2-b454-105e3f755d5b",
          "name": "Work"
        }
      ]
    },
    {
      "note": {
        "id": "4d184f67-d7ce-4488-ae5f-873fc59fc429",
        "title": "Another updated note",
        "body": "This is the new note body",
        "created_at": "2024-02-17T20:08:12.332014Z",
        "updated_at": "2024-02-26T23:10:30.437139Z"
      },
      "categories": [
        {
          "id": "96363d02-c080-40a2-b454-105e3f755d5b",
          "name": "Work"
        }
      ]
    }
  ]
}
```

### `POST /categories/`:
Creates a note category, auth via `Bearer token` using the `access_token`.

Authentication: Bearer Token
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg
```

Body:
```json
{
  "name": "Work"
}
```

Response:
```json
{
  "ID": "96363d02-c080-40a2-b454-105e3f755d5b",
  "Name": "Work"
}
```

### `GET /categories/`:
Gets every category of the user via `access_token` paginated, auth via `Bearer token` using the `access_token`.

Authentication: Bearer Token
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg
```

Body:
```json
{
  "page": 1,
  "page_size": 2
}
```

Response:
```json
{
  "total": 4,
  "page": 1,
  "page_size": 2,
  "items": [
    {
      "ID": "96363d02-c080-40a2-b454-105e3f755d5b",
      "Name": "Work"
    },
    {
      "ID": "f8893fc2-a822-44c8-8e71-4e551ca16833",
      "Name": "Personal"
    }
  ]
}
```

### `DELETE /categories/`:
Deletes a category by its id, auth via `Bearer token` using the `access_token`.

Authentication: Bearer Token
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY2Nlc3MiLCJzdWIiOiJiMDNlMDlhZS04NTgyLTQyZWEtYTk1Yy0xZjJjNGJkMzU4MDIiLCJleHAiOjE3MDgxMjE5MzMsImlhdCI6MTcwODExODMzM30.n14uHTZZ2UhflG5mRG1Xxy1eYo0UPNgJ-0MRczQD0Sg
```

Body:
```json
{
  "category_id": "85ee81f8-c25b-4c45-8580-a3745f36ca2f"
}
```