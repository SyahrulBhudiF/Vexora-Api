package _interface

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist/entity"
	"github.com/zmb3/spotify"
)

type SpotifyServiceInterface interface {
	// GetRecommendations returns recommended tracks based on given limit and track attributes
	GetRecommendations(limit int, trackAttrs *spotify.TrackAttributes) (*entity.PlaylistResponse, error)

	// GetTrackByID returns a track based on the given ID
	GetTrackByID(id string) (*entity.PlaylistResponse, error)

	// SearchTracks searches for tracks based on the given query
	SearchTracks(query string) (*entity.PlaylistResponse, error)
}
