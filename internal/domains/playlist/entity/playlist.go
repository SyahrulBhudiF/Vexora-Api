package entity

type RandomPlaylist struct {
	ID           string `json:"id"`
	PlaylistName string `json:"playlist_name"`
	Artist       string `json:"artist"`
	Path         string `json:"path"`
	Thumbnail    string `json:"thumbnail"`
}

type PlaylistResponse struct {
	Playlists []RandomPlaylist `json:"playlists"`
}

func NewPlaylist(id string, name string, artist string, path string, thumbnail string) *RandomPlaylist {
	return &RandomPlaylist{
		ID:           id,
		PlaylistName: name,
		Artist:       artist,
		Path:         path,
		Thumbnail:    thumbnail,
	}
}

func NewPlaylistResponse(playlists []RandomPlaylist) *PlaylistResponse {
	return &PlaylistResponse{
		Playlists: playlists,
	}
}
