package entity

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/google/uuid"
	"github.com/zmb3/spotify"
)

type RandomMusic struct {
	ID        string `json:"id"`
	MusicName string `json:"playlist_name"`
	Artist    string `json:"artist"`
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
}

type PlaylistResponse struct {
	Music []RandomMusic `json:"music"`
}

func NewPlaylist(id string, name string, artist string, path string, thumbnail string) *RandomMusic {
	return &RandomMusic{
		ID:        id,
		MusicName: name,
		Artist:    artist,
		Path:      path,
		Thumbnail: thumbnail,
	}
}

func NewPlaylistResponse(playlists []RandomMusic) *PlaylistResponse {
	return &PlaylistResponse{
		Music: playlists,
	}
}

type MoodDetectionResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

var MoodTrackAttributes = map[string]spotify.TrackAttributes{
	"sad": *spotify.NewTrackAttributes().
		MaxDanceability(0.65). // Cluster 3
		MaxEnergy(0.42).       // Cluster 3
		MaxValence(0.30).      // Cluster 3
		MaxAcousticness(0.32), // Cluster 3

	"happy": *spotify.NewTrackAttributes().
		MinDanceability(0.70). // Cluster 1
		MaxEnergy(0.65).       // Cluster 1
		MinValence(0.59).      // Cluster 1
		MaxSpeechiness(0.50),  // Cluster 1

	"angry": *spotify.NewTrackAttributes().
		MaxDanceability(0.44). // Cluster 2
		MinEnergy(0.84).       // Cluster 2
		MaxValence(0.28).      // Cluster 2
		MaxLiveness(0.32),     // Cluster 2

	"neutral": *spotify.NewTrackAttributes().
		MaxDanceability(0.56).     // Cluster 0
		MaxEnergy(0.80).           // Cluster 0
		MaxValence(0.27).          // Cluster 0
		MaxInstrumentalness(0.81), // Cluster 0
}

type History struct {
	types.Entity
	UserUUID uuid.UUID      `json:"user_uuid"`
	Mood     string         `json:"mood"`
	Music    []entity.Music `gorm:"foreignKey:history_uuid" json:"music"`
}

func NewHistory(userUUID uuid.UUID, mood string) *History {
	return &History{
		UserUUID: userUUID,
		Mood:     mood,
	}
}

func (h *History) TableName() string {
	return "history"
}
