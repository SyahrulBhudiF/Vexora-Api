package services

import (
	"context"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"math/rand/v2"
)

type SpotifyService struct {
	client *spotify.Client
}

func NewSpotifyService(clientID string, clientSecret string) *SpotifyService {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	httpClient := config.Client(context.Background())
	client := spotify.NewClient(httpClient)

	return &SpotifyService{
		client: &client,
	}
}

func (s *SpotifyService) GetRecommendations(limit int, trackAttrs *spotify.TrackAttributes) (*entity.PlaylistResponse, error) {
	allGenres, err := s.client.GetAvailableGenreSeeds()
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(allGenres), func(i, j int) {
		allGenres[i], allGenres[j] = allGenres[j], allGenres[i]
	})

	seeds := spotify.Seeds{
		Genres: allGenres[:5],
	}

	options := &spotify.Options{
		Limit: &limit,
	}

	if trackAttrs == nil {
		trackAttrs = spotify.NewTrackAttributes()
	}

	recommendations, err := s.client.GetRecommendations(seeds, trackAttrs, options)
	if err != nil {
		return nil, err
	}

	return helpers.ProcessSimpleTracksAsync(recommendations.Tracks, s.client.GetTrack)
}

func (s *SpotifyService) SearchTracks(query string, limit int) (*entity.PlaylistResponse, error) {
	searchOptions := spotify.Options{
		Limit: &limit,
	}

	result, err := s.client.SearchOpt(query, spotify.SearchTypeTrack, &searchOptions)
	if err != nil {
		return nil, err
	}

	tracks := result.Tracks.Tracks
	if len(tracks) > 10 {
		tracks = tracks[:10]
	}

	return helpers.ProcessFullTracks(tracks)
}

func (s *SpotifyService) GetTrackByID(trackID string) (*entity.PlaylistResponse, error) {
	track, err := s.client.GetTrack(spotify.ID(trackID))
	if err != nil {
		return nil, err
	}

	if len(track.Album.Images) == 0 {
		return entity.NewPlaylistResponse([]entity.RandomMusic{}), nil
	}

	playlist := entity.NewPlaylist(
		track.ID.String(),
		track.Name,
		track.Artists[0].Name,
		track.ExternalURLs["spotify"],
		track.Album.Images[0].URL,
	)

	return entity.NewPlaylistResponse([]entity.RandomMusic{*playlist}), nil
}
