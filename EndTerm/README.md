### Use docker command to run services 
`docker-compose up`

### Use example.http to send request to services 

### Services:

* **Authentication:**  Register, log in, manage user tokens.
* **Social Media:** Access, manage, and create posts.

### Authentication API:

##### Sign Up:

```
POST http://0.0.0.0:8081/v1/auth/sign-up

Content-Type: application/json

Body:
{
  "firstName": "String",
  "lastName": "String",
  "login": "String (unique)",
  "password": "String"
}

Response:
{
  "id": "Integer",
  // Other relevant user details
}
```

##### Sign In:

```
POST http://0.0.0.0:8081/v1/auth/sign-in

Content-Type: application/json

Body:
{
  "login": "String",
  "password": "String"
}

Response:
{
  "id": "Integer",
  "token": "String (JWT)"
}
```

##### Sign Out:

```
POST http://0.0.0.0:8081/v1/auth/sign-out

Headers:
Authorization: Bearer <JWT token>

Response:
{}
```

##### Refresh Token:

```
POST http://0.0.0.0:8081/v1/auth/refresh

Headers:
Authorization: Bearer <JWT token>

Response:
{
  "token": "String (new JWT)"
}
```

##### Check Token:

```
POST http://0.0.0.0:8081/v1/auth/check

Headers:
Authorization: Bearer <JWT token>

Response:
{
  "id": "Integer",
  // Other relevant user details
}
```

### Social Media API:

##### Get Posts:

```
GET http://0.0.0.0:8080/api/v1/posts

Headers:
Authorization: Bearer <JWT token>

Response:
[
  {
    "id": "Integer",
    "title": "String",
    "text": "String",
    "image": "String (optional)",
    // Other post details
  }
]
```

##### Get Single Post:

```
GET http://0.0.0.0:8080/api/v1/posts/<id>

Headers:
Authorization: Bearer <JWT token>

Path Parameters:
<id>: Integer (unique post identifier)

Response:
{
  "id": "Integer",
  "title": "String",
  "text": "String",
  "image": "String (optional)",
  // Other post details
}
```

##### Create Post:

```
POST http://0.0.0.0:8080/api/v1/posts

Headers:
Authorization: Bearer <JWT token>

Body:
{
  "title": "String",
  "text": "String",
  "image": "String (optional)"
}

Response:
{
  "id": "Integer",
  // Other post details
}
```

##### Delete Post:

```
DELETE http://0.0.0.0:8080/api/v1/posts/<id>

Headers:
Authorization: Bearer <JWT token>

Path Parameters:
<id>: Integer (unique post identifier)

Response:
{}
```
