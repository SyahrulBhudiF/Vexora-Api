package entity

type RandomPlaylist struct {
	PlaylistName string `json:"playlist_name"`
	Artist       string `json:"artist"`
	Path         string `json:"path"`
	Thumbnail    string `json:"thumbnail"`
}

type SpotifyResponse struct {
	Playlists []RandomPlaylist `json:"playlists"`
}

func NewRandomPlaylist(name, artist, path, thumbnail string) *RandomPlaylist {
	return &RandomPlaylist{
		PlaylistName: name,
		Artist:       artist,
		Path:         path,
		Thumbnail:    thumbnail,
	}
}

func NewSpotifyResponse(playlists []RandomPlaylist) *SpotifyResponse {
	return &SpotifyResponse{
		Playlists: playlists,
	}
}
