package services

import (
	"context"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist/entity"
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

func (s *SpotifyService) GetRandomRecommendations(limit int) (*entity.SpotifyResponse, error) {
	allGenres, err := s.client.GetAvailableGenreSeeds()
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(allGenres), func(i, j int) {
		allGenres[i], allGenres[j] = allGenres[j], allGenres[i]
	})

	seedGenres := allGenres[:5]

	recommendations, err := s.client.GetRecommendations(spotify.Seeds{
		Genres: seedGenres,
	}, &spotify.TrackAttributes{}, &spotify.Options{
		Limit: &limit,
	})

	if err != nil {
		return nil, err
	}

	playlists := make([]entity.RandomPlaylist, 0, len(recommendations.Tracks))

	for _, track := range recommendations.Tracks {
		fullTrack, err := s.client.GetTrack(track.ID)
		if err != nil {
			continue
		}

		playlist := *entity.NewRandomPlaylist(
			fullTrack.Name,
			fullTrack.Artists[0].Name,
			fullTrack.ExternalURLs["spotify"],
			fullTrack.Album.Images[0].URL,
		)
		playlists = append(playlists, playlist)
	}

	return entity.NewSpotifyResponse(playlists), nil

}

func (s *SpotifyService) GetGenreSeeds() ([]string, error) {
	genres, err := s.client.GetAvailableGenreSeeds()
	if err != nil {
		return nil, err
	}

	return genres, nil
}
