# Vexora API Documentation

## Overview

Vexora is a mobile application that provides personalized Spotify music recommendations based on facial emotion
recognition. The application uses:

- Convolutional Neural Network (CNN) for facial emotion detection
- K-means clustering for music mood classification
- Spotify API integration for music recommendations

### Key Features

- **Facial Emotion Recognition**: Detects user's mood through facial expressions using CNN, categorizing emotions into:
    - Happy 😊
    - Sad 😢
    - Angry 😠
    - Neutral/Calm 😐

- **Intelligent Music Recommendation**:
    - Utilizes K-means clustering to categorize Spotify tracks based on audio features
    - Matches detected emotions with appropriate music clusters
    - Provides personalized playlist recommendations based on current mood

- **User Management**:
    - Secure authentication system
    - Personal history tracking
    - Profile customization

## Base URL

```
http://localhost:5555/api/v1
```

## Authentication

All API endpoints except `/login` and `/register` require JWT Bearer token authentication.

```http
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
  "username": "string",
  "password": "string"
}
```

**Response:**

```json
{
  "code": 200,
  "status": "success",
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
  "name": "string",
  "email": "string",
  "username": "string",
  "password": "string"
}
```

### User Management

#### Get User Profile

```http
GET /user
```

**Response:**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "id": 1,
    "profile_picture": "string",
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

- `name`: string
- `email`: string
- `username`: string
- `profile_picture`: file (PNG, JPEG, JPG)

### Music History

#### Get All History

```http
GET /history
```

**Response:**

```json
[
  {
    "code": 0,
    "status": "string",
    "data": [
      {
        "id": 0,
        "user_id": 0,
        "mood": "string",
        "created_at": "2024-11-02T22:55:02.980Z"
      }
    ]
  }
]
```

#### Get Specific History

```http
GET /history/{id}
```

**Response:**

```json
{
  "code": 0,
  "status": "string",
  "data": {
    "id": 0,
    "user_id": 0,
    "mood": "string",
    "created_at": "2024-11-02T22:56:14.785Z",
    "playlist": [
      {
        "id": 0,
        "history_id": 0,
        "name_track": "string",
        "path": "string",
        "thumbnail": "string"
      }
    ]
  }
}
```

### Music Recommendation

```http request
POST /mood-detection
```

**Request Body:**

```json
{
  "user_id": 0,
  "image": "base64",
  "genres": [
    "string",
    "string"
  ],
  "limit": 10
}
```

**Response:**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "detected_mood": "happy",
    "recommended_tracks": [
      {
        "track_id": "spotify:track:123456",
        "name": "Happy Song",
        "artist": "Happy Artist",
        "preview_url": "https://p.scdn.co/mp3-preview/...",
        "spotify_url": "https://open.spotify.com/track/123456",
        "image_url": "https://i.scdn.co/image/..."
      }
    ],
    "playlist_id": "spotify:playlist:789xyz"
  }
}
```

## Complete OpenAPI Specification

For detailed API documentation, you can:

1. View our [OpenAPI Specification](./api/api-spec.yaml)
2. Import our [Postman Collection](./postman-collection.json)

## Error Responses

All endpoints return error responses in the following format:

```json
{
  "code": 400,
  "status": "error message"
}
```

Common HTTP Status Codes:

- 200: Success
- 400: Bad Request
- 401: Unauthorized
- 404: Not Found
- 500: Internal Server Error