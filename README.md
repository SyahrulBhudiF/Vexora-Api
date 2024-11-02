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

#### Get Specific History

```http
GET /history/{id}
```

## Complete OpenAPI Specification

For detailed API documentation, you can:

1. View our [OpenAPI Specification](./api-spec.yaml)
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