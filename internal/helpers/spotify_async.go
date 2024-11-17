package helpers

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/playlist/entity"
	"github.com/zmb3/spotify"
	"sync"
)

// ProcessSimpleTracksAsync handles concurrent processing of SimpleTrack types
func ProcessSimpleTracksAsync(tracks []spotify.SimpleTrack, processor func(spotify.ID) (*spotify.FullTrack, error)) (*entity.PlaylistResponse, error) {
	results := make(chan entity.RandomPlaylist, len(tracks))
	var wg sync.WaitGroup

	for _, track := range tracks {
		wg.Add(1)
		go func(track spotify.SimpleTrack) {
			defer wg.Done()

			fullTrack, err := processor(track.ID)
			if err != nil {
				return
			}

			if len(fullTrack.Album.Images) > 0 {
				playlist := entity.NewPlaylist(
					fullTrack.ID.String(),
					fullTrack.Name,
					fullTrack.Artists[0].Name,
					fullTrack.ExternalURLs["spotify"],
					fullTrack.Album.Images[0].URL,
				)
				results <- *playlist
			}
		}(track)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return collectResults(results)
}

// ProcessFullTracksAsync handles concurrent processing of FullTrack types
func ProcessFullTracksAsync(tracks []spotify.FullTrack) (*entity.PlaylistResponse, error) {
	results := make(chan entity.RandomPlaylist, len(tracks))
	var wg sync.WaitGroup

	for _, track := range tracks {
		if len(track.Album.Images) == 0 {
			continue
		}

		wg.Add(1)
		go func(track spotify.FullTrack) {
			defer wg.Done()
			playlist := entity.NewPlaylist(
				track.ID.String(),
				track.Name,
				track.Artists[0].Name,
				track.ExternalURLs["spotify"],
				track.Album.Images[0].URL,
			)
			results <- *playlist
		}(track)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return collectResults(results)
}

func collectResults(results <-chan entity.RandomPlaylist) (*entity.PlaylistResponse, error) {
	var validPlaylists []entity.RandomPlaylist
	for playlist := range results {
		validPlaylists = append(validPlaylists, playlist)
	}
	return entity.NewPlaylistResponse(validPlaylists), nil
}
