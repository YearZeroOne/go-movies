# Go Movies API
## A RESTful API for managing movies, built with Go and Fiber.

### Prerequisites
- Docker
- Docker Compose
- Go (1.22 or later)

# Go Movies API

## API Endpoints

### User Endpoints

- **Register User**
  - **POST** `/register`
  - **Request Body:**
    ```json
    {
    "id": 1,
    "username": "your-username",
    "password": "hashed-password"
    }
    ```
  - **Response:**
    ```json
    {
        "token": "your-jwt-token"
    }
    ```

- **Login**
  - **POST** `/login`
  - **Request Body:**
    ```json
    {
        "username": "your-username",
        "password": "your-password"
    }
    ```
  - **Response:**
    ```json
    {
        "token": "your-jwt-token"
    }
    ```

### Movies Endpoints

- **Get Movies**
  - **GET** `/movies`
  - **Headers:**
    - `Authorization: Bearer your-jwt-token`
  - **Response:**
    ```json
    [
        {
            "id": 1,
            "title": "Movie Title",
            "url": "http://example.com/image.jpg",
            "genre": "Genre",
            "release": "2024-01-01T00:00:00Z"
        }
    ]
    ```

- **Get Movie By ID**
  - **GET** `/movies/:id`
  - **Headers:**
    - `Authorization: Bearer your-jwt-token`
  - **Response:**
    ```json
    {
        "id": 1,
        "title": "Movie Title",
        "url": "http://example.com/image.jpg",
        "genre": "Genre",
        "release": "2024-01-01T00:00:00Z"
    }
    ```

- **Create Movie**
  - **POST** `/movies`
  - **Headers:**
    - `Authorization: Bearer your-jwt-token`
  - **Request Body:**
    ```json
    {
        "title": "Movie Title",
        "url": "http://example.com/image.jpg",
        "genre": "Genre",
        "release": "2024-01-01T00:00:00Z"
    }
    ```

- **Edit Movie**
  - **PATCH** `/movies/:id`
  - **Headers:**
    - `Authorization: Bearer your-jwt-token`
  - **Request Body:**
    ```json
    {
        "title": "Updated Movie Title",
        "url": "http://example.com/updated-image.jpg",
        "genre": "Updated Genre",
        "release": "2024-01-01T00:00:00Z"
    }
    ```

- **Delete Movie By ID**
  - **DELETE** `/movies/:id`
  - **Headers:**
    - `Authorization: Bearer your-jwt-token`
