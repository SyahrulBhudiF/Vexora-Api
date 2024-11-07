# Vexora API

RESTful API specification for Vexora - A Mood-Based Music Recommendation System

## Features

- User authentication (login, register, refresh token)
- User profile management
- Facial emotion detection
- Mood-based music recommendations
- Music history tracking
- Playlist management

## Mood Categories

- Happy 😊
- Sad 😢
- Angry 😠
- Neutral/Calm 😐

## Base URL

```
http://localhost:5555/api/v1
```

## Authentication

The API uses JWT Bearer token authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_token>
```

## API Endpoints

### Authentication

#### Login

```http
POST /login
```

**Request Body:**

```json
{
  "username": "john_doe",
  "password": "********"
}
```

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "Login successful",
  "data": {
    "accessToken": "string",
    "refreshToken": "string"
  }
}
```

#### Register

```http
POST /register
```

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "username": "john_doe",
  "password": "********"
}
```

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "Registration successful",
  "data": {
    "accessToken": "string",
    "refreshToken": "string"
  }
}
```

### User Management

#### Get User Profile

```http
GET /user
```

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "User profile retrieved successfully",
  "data": {
    "id": 1,
    "profile_picture": "string",
    "file_id": "string",
    "name": "string",
    "email": "string",
    "username": "string",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Update User Profile

```http
PUT /user
```

**Request Body (multipart/form-data):**

- `name`: Updated full name
- `email`: Updated email address
- `username`: Updated username
- `profile_picture`: New profile picture (PNG, JPEG, JPG only)

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "Profile updated successfully",
  "data": {
    "id": 1,
    "profile_picture": "string",
    "file_id": "string",
    "name": "string",
    "email": "string",
    "username": "string",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Change Password

```http
PUT /user/change-password
```

**Request Body:**

```json
{
  "current_password": "string",
  "new_password": "string"
}
```

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "Password changed successfully",
  "data": null
}
```

### Music History

#### Get All History

```http
GET /history
```

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "History retrieved successfully",
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "mood": "happy",
      "playlist_name": "string",
      "path": "string",
      "created_at": "2024-01-01T00:00:00Z",
      "music": [
        {
          "id": 1,
          "history_id": 1,
          "music_name": "string",
          "path": "string",
          "thumbnail": "string",
          "artist": "string"
        }
      ]
    }
  ]
}
```

#### Get Specific History Entry

```http
GET /history/{id}
```

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "History entry found",
  "data": {
    "id": 1,
    "user_id": 1,
    "mood": "happy",
    "playlist_name": "string",
    "path": "string",
    "created_at": "2024-01-01T00:00:00Z",
    "music": [
      {
        "id": 1,
        "history_id": 1,
        "music_name": "string",
        "path": "string",
        "thumbnail": "string",
        "artist": "string"
      }
    ]
  }
}
```

### Mood Detection

#### Detect Mood and Get Recommendations

```http
POST /mood-detection
```

**Request Body (multipart/form-data):**

- `user_id`: ID of the authenticated user
- `image`: Selfie image file (JPEG, PNG)
- `genres` (optional): Array of preferred music genres
    - Available genres: pop, rock, hip-hop, r-n-b, classical, jazz, electronic, indie, metal, country
- `limit` (optional): Maximum number of tracks to recommend (1-50, default: 20)

**Response (200: Success):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "Mood detected successfully",
  "data": {
    "detected_mood": "happy",
    "confidence_score": 0.95,
    "recommended_tracks": [
      {
        "track_id": "string",
        "name": "string",
        "artist": "string",
        "preview_url": "string",
        "spotify_url": "string",
        "image_url": "string"
      }
    ],
    "playlist_id": "string"
  }
}
```

## Error Responses

All endpoints return error responses in the following format:

```json
{
  "success": false,
  "shouldNotify": true,
  "message": "Error description",
  "data": null
}
```

Common HTTP Status Codes:

- 200: Success
- 400: Bad Request
- 401: Unauthorized/Invalid credentials
- 404: Not Found
- 500: Internal Server Error

## Contact

For support or inquiries, please contact:

- Email: ryu4w@gmail.com