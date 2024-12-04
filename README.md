# Vexora API

Vexora is a mood-based music recommendation system that uses facial emotion detection to suggest personalized music
playlists. The system leverages CNN for emotion detection and K-means clustering for music matching.

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

## Cache

Cache Time : 30 Minutes

- To Refresh Cache, Please use Parameter `refresh=true`

## 🚀 Getting Started

### Run Localy

```
docker compose watch
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

5. Send OTP

```http request
POST /send-otp
```

**Request Body:**

```json
{
  "email": "example@gmail.com"
}
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "OTP sent successfully!",
  "data": null
}
```

6. Verify OTP

```http request
POST /verify-email
```

**Request Body:**

```json
{
  "email": "example@gmail.com",
  "otp": "123456"
}
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "email verification success!",
  "data": null
}
```

7. Reset Password

```http request
POST /reset-password
```

**Request Body:**

```json
{
  "email": "example@gmail.com",
  "otp": "123456",
  "password": "newpassword"
}
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "password reset success!",
  "data": null
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
    "uuid": 1,
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
- `username`: "john_doe_updated"

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": true,
  "message": "update profile success!",
  "data": {
    "uuid": 1,
    "profile_picture": "https://example.com/profile.jpg",
    "file_id": "abc123",
    "name": "John Doe",
    "email": "john@example.com",
    "username": "john_doe",
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
  "previous_password": "currentpass123",
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

- `image`: [File Upload]

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
      "uuid": "abc123",
      "user_id": 1,
      "mood": "happy",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### 2. Get Music History By ID

```http
GET /music/{id}
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "history entry found!",
  "data": [
    {
      "uuid": "2c187355-9a61-4b22-b40b-c4636a01ad52",
      "created_at": "2024-11-30T18:36:23.872084Z",
      "history_uuid": "aaea230a-6846-438f-b566-3ff9f1256622",
      "id": "3puYuuZ7lmlTjIgXBOT01k",
      "playlist_name": "Pilihanku",
      "artist": "MALIQ & D'Essentials",
      "path": "https://open.spotify.com/track/3puYuuZ7lmlTjIgXBOT01k",
      "thumbnail": "https://i.scdn.co/image/ab67616d0000b2734b274090757829034de581df"
    }
  ]
}
```

#### 3. Get Most Mood

```http
GET /history/most-mood
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "success",
  "data": {
    "mood": "happy"
  }
}
```

### Spotify API

#### 1. Get Random Recommendations

```http
GET /spotify/random-playlist
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "random playlist retrieved successfully!",
  "data": {
    "music": [
      {
        "id": "0lYBSQXN6rCTvUZvg9S0lU",
        "playlist_name": "Let Me Love You",
        "artist": "DJ Snake",
        "path": "https://open.spotify.com/track/0lYBSQXN6rCTvUZvg9S0lU",
        "thumbnail": "https://i.scdn.co/image/ab67616d0000b273212d776c31027c511f0ee3bc"
      }
    ]
  }
}
```

#### 2. Get Track By ID

```http
GET /spotify/{id}
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "success",
  "data": {
    "music": [
      {
        "id": "6TQghwJw3Sh8h9mTcV5BR7",
        "playlist_name": "Sewanee Mountain Catfight",
        "artist": "Old Crow Medicine Show",
        "path": "https://open.spotify.com/track/6TQghwJw3Sh8h9mTcV5BR7",
        "thumbnail": "https://i.scdn.co/image/ab67616d0000b27366990be0a5b69a9a1ca8b882"
      }
    ]
  }
}
```

#### 3. Search Tracks

```http
GET /spotify/search?search={query}
```

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "success",
  "data": {
    "music": [
      {
        "id": "0PtJbtW50jcvvswNPn3QGd",
        "playlist_name": "Serana",
        "artist": "For Revenge",
        "path": "https://open.spotify.com/track/0PtJbtW50jcvvswNPn3QGd",
        "thumbnail": "https://i.scdn.co/image/ab67616d0000b27346f02ffc0922f939ed0fd53f"
      },
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

**Success Response (200):**

```json
{
  "success": true,
  "shouldNotify": false,
  "message": "mood detected successfully!",
  "data": {
    "detected_mood": "happy",
    "music": [
      {
        "id": "daskd2312opdask",
        "name": "Happy Track",
        "artist": "Happy Artist",
        "path": "https://open.spotify.com/track/123abc",
        "thumbnail": "https://example.com/album/cover1"
      }
    ],
    "created_at": "timestamp"
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