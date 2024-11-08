# Vexora API

Vexora is a mood-based music recommendation system that uses facial emotion detection to suggest personalized music playlists. The system leverages CNN for emotion detection and K-means clustering for music matching.

## 🎯 Features

- 👤 User authentication and profile management
- 😊 Facial emotion detection using CNN
- 🎵 Mood-based music recommendations
- 📝 Music history tracking
- 🎼 Playlist management

## 🎭 Supported Mood Categories

- Happy 😊
- Sad 😢
- Angry 😠
- Neutral/Calm 😐

## 🚀 Getting Started

### Base URL
```
http://localhost:5555/api/v1
```

### Authentication
The API uses JWT Bearer token authentication. Include your token in the Authorization header:
```
Authorization: Bearer <your_token>
```

## 📋 API Endpoints

### Authentication

#### 1. Register New User
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

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": true,
  "message": "register success!",
  "data": {
    "id": 1,
    "profile_picture": "https://example.com/default.jpg",
    "file_id": "abc123",
    "name": "John Doe",
    "email": "john@example.com",
    "username": "john_doe",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

**Error Response (400):**
```json
{
  "success": false,
  "shouldNotify": true,
  "message": "username or email already exists",
  "data": null
}
```

#### 2. Login
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

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": true,
  "message": "login success!",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**Error Response (401):**
```json
{
  "success": false,
  "shouldNotify": true,
  "message": "invalid username or password",
  "data": null
}
```

#### 3. Logout
```http
POST /logout
```

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": false,
  "message": "logout success!",
  "data": null
}
```

#### 4. Refresh Token
```http
POST /refresh
```

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": false,
  "message": "refresh token success!",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**Error Response (401):**
```json
{
  "success": false,
  "shouldNotify": true,
  "message": "invalid or expired refresh token",
  "data": null
}
```

### User Management

#### 1. Get User Profile
```http
GET /user
```

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": false,
  "message": "get profile success!",
  "data": {
    "id": 1,
    "profile_picture": "https://example.com/profile.jpg",
    "file_id": "abc123",
    "name": "John Doe",
    "email": "john@example.com",
    "username": "john_doe",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 2. Update Profile
```http
PUT /user
```

**Request Body (multipart/form-data):**
- `name`: "John Doe Updated"
- `email`: "john.updated@example.com"
- `username`: "john_doe_updated"
- `profile_picture`: [File Upload]

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": true,
  "message": "update profile success!",
  "data": {
    "id": 1,
    "profile_picture": "https://example.com/new-profile.jpg",
    "file_id": "xyz789",
    "name": "John Doe Updated",
    "email": "john.updated@example.com",
    "username": "john_doe_updated",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 3. Change Password
```http
PUT /user/change-password
```

**Request Body:**
```json
{
  "current_password": "currentpass123",
  "new_password": "newpass123"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": true,
  "message": "change password success!",
  "data": null
}
```

**Error Response (400):**
```json
{
  "success": false,
  "shouldNotify": true,
  "message": "current password is incorrect",
  "data": null
}
```

#### 4. Update Profile Picture
```http
PUT /user/profile-picture
```

**Request Body (multipart/form-data):**
- `profile_picture`: [File Upload]

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": true,
  "message": "upload image success!",
  "data": null
}
```

### Music History

#### 1. Get All History
```http
GET /history
```

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": false,
  "message": "history retrieved successfully!",
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "mood": "happy",
      "playlist_name": "Happy Vibes",
      "path": "https://example.com/playlists/happy-vibes",
      "created_at": "2024-01-01T00:00:00Z",
      "music": [
        {
          "id": 1,
          "history_id": 1,
          "music_name": "Happy Song",
          "path": "https://example.com/songs/happy-song",
          "thumbnail": "https://example.com/thumbnails/happy-song",
          "artist": "Happy Artist"
        }
      ]
    }
  ]
}
```

#### 2. Get Specific History
```http
GET /history/{id}
```

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": false,
  "message": "history entry found!",
  "data": {
    "id": 1,
    "user_id": 1,
    "mood": "happy",
    "playlist_name": "Happy Vibes",
    "path": "https://example.com/playlists/happy-vibes",
    "created_at": "2024-01-01T00:00:00Z",
    "music": [
      {
        "id": 1,
        "history_id": 1,
        "music_name": "Happy Song",
        "path": "https://example.com/songs/happy-song",
        "thumbnail": "https://example.com/thumbnails/happy-song",
        "artist": "Happy Artist"
      }
    ]
  }
}
```

### Mood Detection & Recommendations

#### Detect Mood and Get Recommendations
```http
POST /mood-detection
```

**Request Body (multipart/form-data):**
- `user_id`: 1
- `image`: [Selfie Image File]
- `genres`: ["pop", "rock", "jazz"]
- `limit`: 20

**Success Response (200):**
```json
{
  "success": true,
  "shouldNotify": false,
  "message": "mood detected successfully!",
  "data": {
    "detected_mood": "happy",
    "confidence_score": 0.95,
    "recommended_tracks": [
      {
        "track_id": "spotify:track:123abc",
        "name": "Happy Track",
        "artist": "Happy Artist",
        "preview_url": "https://example.com/preview/track1",
        "spotify_url": "https://open.spotify.com/track/123abc",
        "image_url": "https://example.com/album/cover1"
      }
    ],
    "playlist_id": "playlist_123abc"
  }
}
```

## 📊 Status Codes

- `200`: Success
- `400`: Bad Request
- `401`: Unauthorized
- `404`: Not Found
- `500`: Internal Server Error

## 🔒 Security Notes

1. All sensitive routes require JWT authentication
2. Refresh tokens are valid for 30 days
3. File uploads are restricted to PNG, JPEG, and JPG formats
4. Profile pictures are processed and stored securely

## 📱 Technical Requirements

- Supports image upload for facial detection
- Handles multipart/form-data for file uploads
- Processes JSON for standard requests
- Returns standardized JSON responses

## 📞 Contact & Support

For support or inquiries, please contact:
- Email: ryu4w@gmail.com

## ⚠️ Rate Limiting

Please note that API endpoints may have rate limiting applied. Contact support for specific limitations.

## 🔄 Version Information

- Current Version: 1.0.0
- Last Updated: 2024